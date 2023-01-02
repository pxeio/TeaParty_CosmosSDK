package adams

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
)

// Watch is the blocking function that: continuously updates the orders and complete orders
// from the party chain, monitors the accounts in question for new transactions, and then
// updates the party chain with the outcome of the transaction.
func (e *ExchangeServer) Watch(ctx context.Context) {

	// for {
	select {
	case <-ctx.Done():
		e.logger.Info("context done, stopping watch")
		// TODO: tell the party chain that we are stoppig
		return
	default:
		// Fetch the current orders from the party chain.
		orders, err := fetchTradeOrdersFromPartyChain(e.partyNode)
		if err != nil {
			e.logger.Errorw("failed to fetch orders from party chain")
			if !e.dev {
				panic(err)
			}
		}
		e.partyChainOrders = orders

		// Fetch the current complete orders from the party chain.
		completeOrders, err := fetchPendingOrdersFromPartyChain(e.partyNode)
		if err != nil {
			e.logger.Errorw("failed to fetch complete orders from party chain")
			if !e.dev {
				panic(err)
			}
		}
		e.partyChainCompleteOrders = completeOrders

		e.initMonitor(completeOrders)
		// // Monitor the accounts for new transactions.
		// e.monitorAccounts(ctx)
		// // Update the party chain with the outcome of the transactions.
		// e.updatePartyChain(ctx)

	}
	// }

}

// initMonitor
func (e *ExchangeServer) initMonitor(qacor *types.QueryAllPendingOrdersResponse) {
	for _, order := range qacor.PendingOrders {
		e.logger.Infof("Order: %+v", order)
		var ta string
		ta = order.TradeAsset

		const productionTimeLimit = 7200 // 2 hours
		const devTimelimit = 300         // 300 second
		var timeLimit int64
		if e.dev {
			timeLimit = devTimelimit
		} else {
			timeLimit = productionTimeLimit
		}

		biPrice := new(big.Int)
		// biPrice.NewFromString(order.Price, 10)
		d, err := decimal.NewFromString(order.Price)
		if err != nil {
			e.logger.Error("Error converting price to decimal")
			return
		}
		biPrice = d.BigInt()
		// biPrice, ok := biPrice.SetString(order.Price, 10)
		// if !ok {
		// 	e.logger.Error("Error converting price to big.Int")
		// 	return
		// }

		biAmount := new(big.Int)
		dA, err := decimal.NewFromString(order.Amount)
		if err != nil {
			e.logger.Error("Error converting amount to decimal")
			return
		}

		biAmount = dA.BigInt()

		// biAmount, ok = biAmount.SetString(order.Amount, 10)
		// if !ok {
		// 	e.logger.Error("Error converting amount to big.Int")
		// 	return
		// }

		taAmount := biAmount

		if order.TradeAsset == "ANY" {
			// fetch the market price of the trade asset
			marketPrice, err := FetchMarketPriceInUSD(ta)
			if err != nil {
				e.logger.Error("error fetching market price: " + err.Error())
				return
			}

			e.logger.Infof("market price of %s is: %s", ta, marketPrice)

			// calcuate the amount to send to the buyer
			pgto := big.NewFloat(0).SetInt(biPrice)
			bito := big.NewFloat(0).SetInt(marketPrice)

			// convert to big.int

			fl, _ := pgto.Quo(pgto, bito).Float64()

			// taAmount = FloatToBigInt(fl)
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
			OrderID:               order.Index,
			BuyerShippingAddress:  order.BuyerShippingAddress,
			SellerShippingAddress: order.SellerShippingAddress,
			TradeAsset:            ta,
			Price:                 biPrice,
			Currency:              order.Currency,
			Amount:                taAmount,
			Timeout:               timeLimit,
			SellerNKNAddress:      order.SellerNKNAddress,
			BuyerNKNAddress:       order.BuyerNKNAddress,
		}

		buyersAccountWatchRequest := &AccountWatchRequest{}
		sellersAccountWatchRequest := &AccountWatchRequest{}

		e.logger.Infof("currency: %s", co.Currency)
		switch order.Currency {
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
				return
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
				return
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
				return
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
				return
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
				return
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
			return
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
				return
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
				return
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
				return
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
				return
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
				return
			}

			// emit a new event to let Warren know that we need to start watching a new account
			buyersAccountWatchRequest = &AccountWatchRequest{
				Account:       co.BuyerEscrowWallet.PublicAddress,
				TimeOut:       co.Timeout,
				Chain:         POL,
				Amount:        co.Price,
				TransactionID: co.OrderID,
			}
		default:
			e.logger.Error("invalid trade asset chain: " + co.TradeAsset)
			return
		}

		go e.watchAccount(buyersAccountWatchRequest)
		go e.watchAccount(sellersAccountWatchRequest)
		e.logger.Info("now watching order: " + co.OrderID)
	}
}

func (e *ExchangeServer) watchAccount(awr *AccountWatchRequest) {
	e.logger.Info("watching account: " + awr.Account)
	switch awr.Chain {
	case SOL:
		go e.waitAndVerifySOLChain(*awr)
	case CEL:
		go e.waitAndVerifyEVMChain(context.Background(), e.celoNode.rpcClient, e.celoNode.rpcClientTwo, *awr)
	case ETH:
		go e.waitAndVerifyEVMChain(context.Background(), e.ethNode.rpcClient, e.ethNode.rpcClientTwo, *awr)
	case POL:
		go e.waitAndVerifyEVMChain(context.Background(), e.polygonNode.rpcClient, e.polygonNode.rpcClientTwo, *awr)
	}
}
