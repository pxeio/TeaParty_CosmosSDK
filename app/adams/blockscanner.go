package adams

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Watch is the blocking function that: continuously updates the orders and complete orders
// from the party chain, monitors the accounts in question for new transactions, and then
// updates the party chain with the outcome of the transaction.
func (e *ExchangeServer) Watch(ctx sdk.Context) {
	// go e.initMonitor(ctx)
	for {
		select {
		case <-ctx.Done():
			e.logger.Info("context done, stopping watch")
			// TODO: tell the party chain that we are stoppig
			return
		default:
			// Fetch current orders from the party chain.
			e.logger.Info("fetching orders from party chain...")
			orders := e.PartyKeeper.GetAllPendingOrders(ctx)

			if len(orders) == 0 {
				e.logger.Info("no orders to monitor")
				time.Sleep(5 * time.Second)
				continue
			}

			e.logger.Info("found orders.. checking if they are being watched or are already completed...")
			if len(e.ordersInProgress) > 0 {
				for _, oip := range e.ordersInProgress {
					// if the order has been paid on both sides, then we can complete the order
					e.logger.Info("checking order for completion: " + oip.Index)
					if oip.BuyerPaymentComplete && oip.SellerPaymentComplete {
						e.logger.Info("order complete, calling complete order")
						// TODO: implement order completion logic
						e.PartyKeeper.RemovePendingOrders(ctx, oip.Index)
						e.logger.Info("removed order from pending orders: " + oip.Index)
						continue
					}
				}
			}

			// look at the orders in progress fetched from the chain state and see if they are
			// already in the exchange server's orders in progress
			for _, o := range orders {
				e.logger.Info("checking order: " + o.Index)
				var found bool
				for _, oip := range e.ordersInProgress {
					if o.Index == oip.Index {
						e.logger.Info("order already being watched: " + o.Index)
						found = true
						continue
					}
				}

				for _, oiw := range e.ordersInWatch {
					if o.Index == oiw {
						e.logger.Info("order already being watched: " + o.Index)
						found = true
						continue
					}
				}

				if !found {
					e.logger.Info("adding order to orders in progress: " + o.Index)
					if len(e.ordersInProgress) > 0 {
						e.ordersInProgress = append(e.ordersInProgress, o)
					} else {
						e.ordersInProgress = make([]types.PendingOrders, 0)
					}
					e.logger.Info("added order to orders in progress: " + o.Index)
					go e.initMonitor(o, ctx)
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}

// initMonitor
func (e *ExchangeServer) initMonitor(watchOrder types.PendingOrders, ctx sdk.Context) error {
	for {
		fmt.Println("initMonitor: checking orders in progress...")
		fmt.Printf("initMonitor: orders in progress: %+v", e.ordersInProgress)
		e.logger.Info("initMonitor: checking order: " + watchOrder.Index)
		for _, oiw := range e.ordersInWatch {
			if oiw == watchOrder.Index {
				e.logger.Info("order already being watched: " + watchOrder.Index)
				continue
			}
		}

		e.logger.Infof("Order: %+v", watchOrder)
		ta := watchOrder.TradeAsset

		const productionTimeLimit = 7200 // 2 hours
		const devTimelimit = 300         // 300 second
		var timeLimit int64
		if e.dev {
			timeLimit = devTimelimit
		} else {
			timeLimit = productionTimeLimit
		}

		d, err := decimal.NewFromString(watchOrder.Price)
		if err != nil {
			e.logger.Error("Error converting price to decimal")
			return err
		}
		biPrice := d.BigInt()

		dA, err := decimal.NewFromString(watchOrder.Amount)
		if err != nil {
			e.logger.Error("Error converting amount to decimal")
			return err
		}
		biAmount := dA.BigInt()

		taAmount := biAmount
		if watchOrder.TradeAsset == "ANY" {
			// fetch the market price of the trade asset
			marketPrice, err := e.fetchMarketPriceInUSD(ta)
			if err != nil {
				e.logger.Error("error fetching market price: " + err.Error())
				return err
			}

			e.logger.Infof("market price of %s is: %s", ta, marketPrice)

			// calcuate the amount to send to the buyer
			pgto := big.NewFloat(0).SetInt(biPrice)
			bito := big.NewFloat(0).SetInt(marketPrice)
			fl, _ := pgto.Quo(pgto, bito).Float64()

			// TODO: TEST THIS!!! this is a big change from how it was before
			taAmount = big.NewInt(int64(fl * 100000000))
			// taAmount = big.NewInt(int64(fl * 100000000))

			e.logger.Infof("calculated amount to send to buyer: %s", fl)
		}

		fmt.Println(biAmount)
		fmt.Println(biPrice)
		fmt.Println(timeLimit)

		// e.logger.Infof("set tradeAsset to: %s", ta)

		co := &CompletedOrder{
			OrderID:               watchOrder.Index,
			BuyerShippingAddress:  watchOrder.BuyerShippingAddress,
			SellerShippingAddress: watchOrder.SellerShippingAddress,
			TradeAsset:            ta,
			Price:                 biPrice,
			Currency:              watchOrder.Currency,
			Amount:                taAmount,
			Timeout:               timeLimit,
			SellerNKNAddress:      watchOrder.SellerNKNAddress,
			BuyerNKNAddress:       watchOrder.BuyerNKNAddress,
		}

		buyersAccountWatchRequest := &AccountWatchRequest{}
		sellersAccountWatchRequest := &AccountWatchRequest{}

		e.logger.Infof("currency: %s", co.Currency)
		switch watchOrder.Currency {
		case SOL:
			// generate a new solana account for the buyer
			acc := e.CreateSolanaAccount()

			co.SellerEscrowWallet = EscrowWallet{
				PublicAddress: acc.PublicKey,
				PrivateKey:    acc.PrivateKey,
				Chain:         SOL,
			}

			if err := e.notifySellerOfBuyer(*co); err != nil {
				e.logger.Error("error notifying seller of buyer: " + err.Error())
				return err
			}

			sellersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.SellerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         SOL,
				Amount:        co.Amount,
				TransactionID: co.OrderID,
				Seller:        true,
			}

		case CEL:
			acc := e.generateEVMAccount(CEL)
			co.SellerEscrowWallet = EscrowWallet{
				ECDSA:         acc,
				PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
				PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
				Chain:         CEL,
			}

			if err := e.notifySellerOfBuyer(*co); err != nil {
				e.logger.Error("error notifying seller of buyer: " + err.Error())
				return err
			}

			sellersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.SellerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         CEL,
				Amount:        co.Amount,
				TransactionID: co.OrderID,
				Seller:        true,
			}

		case ETH:
			acc := e.generateEVMAccount(ETH)
			co.SellerEscrowWallet = EscrowWallet{
				ECDSA:         acc,
				PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
				PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
				Chain:         ETH,
			}

			if err := e.notifySellerOfBuyer(*co); err != nil {
				e.logger.Error("error notifying seller of buyer: " + err.Error())
				return err
			}

			sellersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.SellerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         ETH,
				Amount:        co.Amount,
				TransactionID: co.OrderID,
				Seller:        true,
			}

		case POL:
			acc := e.generateEVMAccount(POL)
			co.SellerEscrowWallet = EscrowWallet{
				ECDSA:         acc,
				PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
				PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
				Chain:         POL,
			}

			if err := e.notifySellerOfBuyer(*co); err != nil {
				e.logger.Error("error notifying seller of buyer: " + err.Error())
				return err
			}

			sellersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.SellerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         POL,
				Amount:        co.Amount,
				TransactionID: co.OrderID,
				Seller:        true,
			}

		case MO:
			acc := e.generateEVMAccount(MO)
			co.SellerEscrowWallet = EscrowWallet{
				ECDSA:         acc,
				PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
				PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
				Chain:         MO,
			}

			if err := e.notifySellerOfBuyer(*co); err != nil {
				e.logger.Error("error notifying seller of buyer: " + err.Error())
				return err
			}

			sellersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.SellerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         MO,
				Amount:        co.Amount,
				TransactionID: co.OrderID,
				Seller:        true,
			}
		default:
			e.logger.Error("error sorting the currency type")
			return fmt.Errorf("error sorting the currency type")
		}

		e.logger.Infof("ta is %s", ta)
		switch ta {
		case SOL:
			acc := e.CreateSolanaAccount()

			co.BuyerEscrowWallet = EscrowWallet{
				PublicAddress: acc.PublicKey,
				PrivateKey:    acc.PrivateKey,
				Chain:         SOL,
			}

			if err := e.sendBuyerPayInfo(*co); err != nil {
				e.logger.Error("error sending buyer pay info: " + err.Error())
				return err
			}

			buyersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.BuyerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         SOL,
				Amount:        co.Price,
				TransactionID: co.OrderID,
			}

		case MO:
			acc := e.generateEVMAccount(MO)
			co.BuyerEscrowWallet = EscrowWallet{
				ECDSA:         acc,
				PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
				PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
				Chain:         MO,
			}

			if err := e.sendBuyerPayInfo(*co); err != nil {
				e.logger.Error("error sending buyer pay info: " + err.Error())
				return err
			}

			buyersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.BuyerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         MO,
				Amount:        co.Price,
				TransactionID: co.OrderID,
			}

		case ETH:
			acc := e.generateEVMAccount(ETH)
			co.BuyerEscrowWallet = EscrowWallet{
				ECDSA:         acc,
				PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
				PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
				Chain:         ETH,
			}

			if err := e.sendBuyerPayInfo(*co); err != nil {
				e.logger.Error("error sending buyer pay info: " + err.Error())
				return err
			}

			buyersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.BuyerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         ETH,
				Amount:        co.Price,
				TransactionID: co.OrderID,
			}
		case CEL:
			acc := e.generateEVMAccount(CEL)
			co.BuyerEscrowWallet = EscrowWallet{
				ECDSA:         acc,
				PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
				PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
				Chain:         CEL,
			}

			if err := e.sendBuyerPayInfo(*co); err != nil {
				e.logger.Error("error sending buyer pay info: " + err.Error())
				return err
			}

			buyersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.BuyerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         CEL,
				Amount:        co.Price,
				TransactionID: co.OrderID,
			}

		case POL:
			acc := e.generateEVMAccount(POL)
			co.BuyerEscrowWallet = EscrowWallet{
				ECDSA:         acc,
				PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
				PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
				Chain:         POL,
			}

			if err := e.sendBuyerPayInfo(*co); err != nil {
				e.logger.Error("error sending buyer pay info: " + err.Error())
				return err
			}

			// emit a new event to let adams know that we need to start watching a new account
			buyersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.BuyerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         POL,
				Amount:        co.Price,
				TransactionID: co.OrderID,
			}
		default:
			e.logger.Error("invalid trade asset chain: " + co.TradeAsset)
			return err
		}

		// add this watch request to the map
		e.ordersInWatch = append(e.ordersInWatch, co.OrderID)
		go e.watchAccount(buyersAccountWatchRequest, ctx)
		go e.watchAccount(sellersAccountWatchRequest, ctx)
		e.logger.Info("Watch function has proccessed for order: " + co.OrderID)
	}
	return nil
}

func (e *ExchangeServer) watchAccount(awr *AccountWatchRequest, ctx sdk.Context) {
	e.logger.Info("watching account: " + awr.Account)
	switch awr.Chain {
	case SOL:
		go e.waitAndVerifySOLChain(*awr, ctx)
	case CEL:
		go e.waitAndVerifyEVMChain(context.Background(), e.celoNode.rpcClient, e.celoNode.rpcClientTwo, *awr, ctx)
	case ETH:
		go e.waitAndVerifyEVMChain(context.Background(), e.ethNode.rpcClient, e.ethNode.rpcClientTwo, *awr, ctx)
	case POL:
		go e.waitAndVerifyEVMChain(context.Background(), e.polygonNode.rpcClient, e.polygonNode.rpcClientTwo, *awr, ctx)
	}
}
