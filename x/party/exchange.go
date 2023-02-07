package party

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	partyTypes "github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solRPC "github.com/gagliardetto/solana-go/rpc"
	"github.com/shopspring/decimal"
)

func (am AppModule) dispatch(ctx sdk.Context, awrr *AccountWatchRequestResult) {
	// ouw, ok := am.keeper.GetOrdersUnderWatch(ctx, awrr.AccountWatchRequest.TransactionID)
	// if !ok {
	// 	fmt.Printf("could not find a matching order under watch for  %+v", awrr)
	// 	fmt.Println("here is a list of all orders under watch")
	// 	aouw := am.keeper.GetAllOrdersUnderWatch(ctx)
	// 	for _, ouw := range aouw {
	// 		fmt.Printf("%+v", ouw)
	// 	}

	// 	return
	// }

	fmt.Printf("looking for finalizing order %+v", awrr)
	pendingFinalizingOrder, ok := am.keeper.GetFinalizingOrders(ctx, awrr.AccountWatchRequest.TransactionID)
	if !ok {
		fmt.Printf("could not find a matching finalizing order for %+v", awrr)
		fmt.Println("here is a list of all finalizing orders")
		afo := am.keeper.GetAllFinalizingOrders(ctx)
		for _, fo := range afo {
			fmt.Printf("%+v", fo)
		}

		return
	}

	switch awrr.Result {
	case OUTCOME_SUCCESS:
		fmt.Println("success")
		// TODO:: remove this

		// compare owu.Chain to the Curency of the pendingFinalizingOrder
		// if they match we know that we have the Buyer's payment completion status
		if awrr.AccountWatchRequest.Chain == pendingFinalizingOrder.Currency {
			fmt.Println("BUYER PAYMENT COMPLETE")
			pendingFinalizingOrder.BuyerPaymentComplete = true
		}
		if awrr.AccountWatchRequest.Chain == pendingFinalizingOrder.TradeAsset {
			fmt.Println("SELLER PAYMENT COMPLETE")
			pendingFinalizingOrder.SellerPaymentComplete = true
		}
	case OUTCOME_FAILURE:
		// TODO:: remove this

		if awrr.AccountWatchRequest.Chain == pendingFinalizingOrder.Currency {
			pendingFinalizingOrder.BuyerPaymentComplete = false
		}
		if awrr.AccountWatchRequest.Chain == pendingFinalizingOrder.TradeAsset {
			pendingFinalizingOrder.SellerPaymentComplete = false
		}
		fmt.Println("creating a refund transaction")
		// create the finalizer order
		sellerFinalizerForRefund := partyTypes.OrdersAwaitingFinalizer{
			Index:            pendingFinalizingOrder.SellerEscrowWalletPublicKey,
			NknAddress:       pendingFinalizingOrder.SellerNKNAddress,
			WalletPrivateKey: pendingFinalizingOrder.SellerEscrowWalletPrivateKey,
			WalletPublicKey:  pendingFinalizingOrder.SellerEscrowWalletPublicKey,
			ShippingAddress:  pendingFinalizingOrder.SellerShippingAddress,
			RefundAddress:    pendingFinalizingOrder.SellerRefundAddress,
			Amount:           pendingFinalizingOrder.Amount,
			Chain:            pendingFinalizingOrder.Currency,
		}

		buyerFinalizerForRefund := partyTypes.OrdersAwaitingFinalizer{
			Index:            pendingFinalizingOrder.BuyerEscrowWalletPublicKey,
			NknAddress:       pendingFinalizingOrder.BuyerNKNAddress,
			WalletPrivateKey: pendingFinalizingOrder.BuyerEscrowWalletPrivateKey,
			WalletPublicKey:  pendingFinalizingOrder.BuyerEscrowWalletPublicKey,
			ShippingAddress:  pendingFinalizingOrder.BuyerShippingAddress,
			RefundAddress:    pendingFinalizingOrder.BuyerRefundAddress,
			Amount:           pendingFinalizingOrder.Price,
			Chain:            pendingFinalizingOrder.TradeAsset,
		}

		fmt.Printf("buyer order awaiting finalizer: %+v ", buyerFinalizerForRefund)
		fmt.Println("------")
		fmt.Printf("seller order awaiting finalizer: %+v ", sellerFinalizerForRefund)

		if err := am.finalizeOrder(ctx, buyerFinalizerForRefund); err != nil {
			// TODO: handle error
			fmt.Println("error: ", err)
		}
		if err := am.finalizeOrder(ctx, sellerFinalizerForRefund); err != nil {
			// TODO: handle error
			fmt.Println("error: ", err)
		}
		am.keeper.RemovePendingOrders(ctx, pendingFinalizingOrder.Index)
		am.keeper.SetFinalizingOrders(ctx, pendingFinalizingOrder)
		// am.keeper.RemovePendingOrders(ctx, pendingFinalizingOrder.Index)
		am.wg.Done()
		return
	case OUTCOME_TIMEOUT:
		if awrr.AccountWatchRequest.Chain == pendingFinalizingOrder.Currency {
			pendingFinalizingOrder.BuyerPaymentComplete = false
		}
		if awrr.AccountWatchRequest.Chain == pendingFinalizingOrder.TradeAsset {
			pendingFinalizingOrder.SellerPaymentComplete = false
		}
		fmt.Println("creating a refund transaction")
		// create the finalizer order
		sellerFinalizerForRefund := partyTypes.OrdersAwaitingFinalizer{
			Index:            pendingFinalizingOrder.SellerEscrowWalletPublicKey,
			NknAddress:       pendingFinalizingOrder.SellerNKNAddress,
			WalletPrivateKey: pendingFinalizingOrder.SellerEscrowWalletPrivateKey,
			WalletPublicKey:  pendingFinalizingOrder.SellerEscrowWalletPublicKey,
			ShippingAddress:  pendingFinalizingOrder.SellerShippingAddress,
			RefundAddress:    pendingFinalizingOrder.SellerRefundAddress,
			Amount:           pendingFinalizingOrder.Amount,
			Chain:            pendingFinalizingOrder.Currency,
		}

		buyerFinalizerForRefund := partyTypes.OrdersAwaitingFinalizer{
			Index:            pendingFinalizingOrder.BuyerEscrowWalletPublicKey,
			NknAddress:       pendingFinalizingOrder.BuyerNKNAddress,
			WalletPrivateKey: pendingFinalizingOrder.BuyerEscrowWalletPrivateKey,
			WalletPublicKey:  pendingFinalizingOrder.BuyerEscrowWalletPublicKey,
			ShippingAddress:  pendingFinalizingOrder.BuyerShippingAddress,
			RefundAddress:    pendingFinalizingOrder.BuyerRefundAddress,
			Amount:           pendingFinalizingOrder.Price,
			Chain:            pendingFinalizingOrder.TradeAsset,
		}

		fmt.Printf("buyer order awaiting finalizer: %+v ", buyerFinalizerForRefund)
		fmt.Println("------")
		fmt.Printf("seller order awaiting finalizer: %+v ", sellerFinalizerForRefund)

		if err := am.finalizeOrder(ctx, buyerFinalizerForRefund); err != nil {
			// TODO: handle error
			fmt.Println("error: ", err)
		}
		if err := am.finalizeOrder(ctx, sellerFinalizerForRefund); err != nil {
			// TODO: handle error
			fmt.Println("error: ", err)
		}
		am.keeper.RemovePendingOrders(ctx, pendingFinalizingOrder.Index)
		am.keeper.SetFinalizingOrders(ctx, pendingFinalizingOrder)
		// am.keeper.RemovePendingOrders(ctx, pendingFinalizingOrder.Index)
		am.wg.Done()
		return
	}

	// check if the buyer and seller have both completed payment
	if pendingFinalizingOrder.BuyerPaymentComplete && pendingFinalizingOrder.SellerPaymentComplete {
		fmt.Println("buyer and seller have both completed payment")
		fmt.Println("creating the OrdersAwaitingFinalizer for both the buyer and seller")
		// Create the OrdersAwaitingFinalizer for both the buyer and seller

		boaf := partyTypes.OrdersAwaitingFinalizer{
			Index:            pendingFinalizingOrder.BuyerEscrowWalletPublicKey,
			NknAddress:       pendingFinalizingOrder.BuyerNKNAddress,
			WalletPrivateKey: pendingFinalizingOrder.SellerEscrowWalletPrivateKey,
			WalletPublicKey:  pendingFinalizingOrder.SellerEscrowWalletPublicKey,
			ShippingAddress:  pendingFinalizingOrder.BuyerShippingAddress,
			RefundAddress:    pendingFinalizingOrder.BuyerRefundAddress,
			Amount:           pendingFinalizingOrder.Price,
			Chain:            pendingFinalizingOrder.Currency,
		}

		soaf := partyTypes.OrdersAwaitingFinalizer{
			Index:            pendingFinalizingOrder.SellerEscrowWalletPublicKey,
			NknAddress:       pendingFinalizingOrder.SellerNKNAddress,
			WalletPrivateKey: pendingFinalizingOrder.BuyerEscrowWalletPrivateKey,
			WalletPublicKey:  pendingFinalizingOrder.BuyerEscrowWalletPublicKey,
			ShippingAddress:  pendingFinalizingOrder.SellerShippingAddress,
			RefundAddress:    pendingFinalizingOrder.SellerRefundAddress,
			Amount:           pendingFinalizingOrder.Price,
			Chain:            pendingFinalizingOrder.TradeAsset,
		}

		fmt.Printf("buyer order awaiting finalizer: %+v ", boaf)
		fmt.Println("------")
		fmt.Printf("seller order awaiting finalizer: %+v ", soaf)

		if err := am.finalizeOrder(ctx, boaf); err != nil {
			// TODO: handle error
			fmt.Println("error: ", err)
		}
		if err := am.finalizeOrder(ctx, soaf); err != nil {
			// TODO: handle error
			fmt.Println("error: ", err)
		}
		am.keeper.RemovePendingOrders(ctx, pendingFinalizingOrder.Index)
		am.keeper.SetFinalizingOrders(ctx, pendingFinalizingOrder)
		am.wg.Done()
		return
	} else {
		// update the pending order
		am.keeper.SetFinalizingOrders(ctx, pendingFinalizingOrder)
	}
	am.wg.Done()
}

func (am AppModule) sendFunds(order partyTypes.OrdersAwaitingFinalizer) error {
	// convert the order.Amount to a big int
	dA, err := decimal.NewFromString(order.Amount)
	if err != nil {
		return err
	}
	biAmount := dA.BigInt()

	curencyFee := new(big.Int).Div(biAmount, big.NewInt(100))
	biAmount.Sub(biAmount, curencyFee)
	assetFee := new(big.Int).Div(biAmount, big.NewInt(100))
	biAmount.Sub(biAmount, assetFee)

	// send the funds
	switch order.Chain {
	case SOL:
		solClient := solRPC.New("https://api.testnet.solana.com")
		_, err := solClient.GetRecentBlockhash(context.Background(), solRPC.CommitmentRecent)
		if err != nil {
			return err
		}

		if err := am.sendCoreSOLAsset(order.WalletPrivateKey, order.ShippingAddress, order.Index, biAmount, solClient); err != nil {
			return err
		}
	case CEL:
		// create a new client for the second node
		// initialize the ethereum nodes.
		celClient1, err := ethclient.Dial("https://celo-alfajores.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		err = am.sendCoreEVMAsset(order.WalletPrivateKey, order.WalletPublicKey, order.ShippingAddress, biAmount, order.Index, celClient1)
		if err != nil {
			return err
		}
	case POL:
		polyClient1, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		err = am.sendCoreEVMAsset(order.WalletPrivateKey, order.WalletPublicKey, order.ShippingAddress, biAmount, order.Index, polyClient1)
		if err != nil {
			return err
		}
	case ETH:
		ethClient1, err := ethclient.Dial("https://goerli.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}

		err = am.sendCoreEVMAsset(order.WalletPrivateKey, order.WalletPublicKey, order.ShippingAddress, biAmount, order.Index, ethClient1)
		if err != nil {
			return err
		}
	}
	return nil
}

func (am AppModule) finalizeOrder(ctx sdk.Context, order partyTypes.OrdersAwaitingFinalizer) error {
	fmt.Printf("finalizing order: %+v", order)
	if err := am.sendPrivateKey(order); err != nil {
		if err := am.sendFunds(order); err != nil {
			// am.keeper.SetFailedOrders()
			// TODO: we need to notify the party chain that this has happend &|
			// we need to build a reconciler to adjust the parameters in the order
			// and try to force it through again.
			return err
		}

		// find the matching order in the finalizing orders list by seller nkn address
		// and remove it from the list of finalizing orders
		// fo := am.keeper.GetAllFinalizingOrders(ctx)
		// for _, o := range fo {
		// 	if o.SellerNKNAddress == order.NknAddress || o.BuyerNKNAddress == order.NknAddress {
		// 		am.keeper.RemoveFinalizingOrders(ctx, o.Index)
		// 	}
		// }

	}

	// TODO:: notify the party chain that the order has been finished and the funds have been sent
	// and confirmed

	// if err := notifyPartyChainOfTransactionResult(order.Index, "success"); err != nil {
	// 	e.logger.Error("error notifying the party chain that the transaction was successful: " + err.Error())
	// }

	am.keeper.RemoveOrdersAwaitingFinalizer(ctx, order.Index)
	// fetch the orders in finalizer and update it to know that it was paid
	return nil
}

func (am AppModule) watchAccount(ctx sdk.Context, awr *AccountWatchRequest) error {
	switch awr.Chain {
	case SOL:
		solClient := solRPC.New("https://api.testnet.solana.com")
		_, err := solClient.GetRecentBlockhash(context.Background(), solRPC.CommitmentRecent)
		if err != nil {
			return err
		}

		solClient2 := solRPC.New("https://api.testnet.solana.com")
		_, err = solClient2.GetRecentBlockhash(context.Background(), solRPC.CommitmentRecent)
		if err != nil {
			return err
		}
		go am.waitAndVerifySOLChain(ctx, *awr, solClient, solClient2)
	case CEL:
		// create a new client for the second node
		// initialize the ethereum nodes.
		celClient1, err := ethclient.Dial("https://celo-alfajores.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		celClient2, err := ethclient.Dial("https://celo-alfajores.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		go am.waitAndVerifyEVMChain(ctx, celClient1, celClient2, *awr)
	case ETH:
		// initialize the ethereum nodes.
		ethClient1, err := ethclient.Dial("https://goerli.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		ethClient2, err := ethclient.Dial("https://goerli.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		go am.waitAndVerifyEVMChain(ctx, ethClient1, ethClient2, *awr)
	case POL:
		// initialize the ethereum nodes.
		polyClient1, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		polyClient2, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		go am.waitAndVerifyEVMChain(ctx, polyClient1, polyClient2, *awr)
	}
	return nil
}

func (am AppModule) initMonitor(ctx sdk.Context, order partyTypes.PendingOrders) error {
	ta := order.TradeAsset
	const productionTimeLimit = 7200 // 2 hours
	const devTimelimit = 300         // 300 second
	var timeLimit int64
	// timeLimit = devTimelimit
	// if e.dev {
	// 	timeLimit = devTimelimit
	// } else {
	timeLimit = productionTimeLimit
	// }

	biPrice := new(big.Int)
	d, _ := decimal.NewFromString(order.Price)
	// if err != nil {
	// 	e.logger.Error("Error converting price to decimal")
	// 	return
	// }
	biPrice = d.BigInt()

	biAmount := new(big.Int)
	dA, _ := decimal.NewFromString(order.Amount)
	// if err != nil {
	// 	e.logger.Error("Error converting amount to decimal")
	// 	return
	// }

	biAmount = dA.BigInt()
	taAmount := biAmount

	// if order.TradeAsset == "ANY" {
	// 	// fetch the market price of the trade asset
	// 	marketPrice, err := FetchMarketPriceInUSD(ta)
	// 	if err != nil {
	// 		e.logger.Error("error fetching market price: " + err.Error())
	// 		return
	// 	}

	// 	e.logger.Infof("market price of %s is: %s", ta, marketPrice)

	// 	// calcuate the amount to send to the buyer
	// 	pgto := big.NewFloat(0).SetInt(biPrice)
	// 	bito := big.NewFloat(0).SetInt(marketPrice)

	// 	// convert to big.int

	// 	fl, _ := pgto.Quo(pgto, bito).Float64()

	// 	// taAmount = FloatToBigInt(fl)
	// 	// TODO: TEST THIS!!! this is a big change from how it was before
	// 	taAmount = big.NewInt(int64(fl * 100000000))
	// 	// taAmount = big.NewInt(int64(fl * 100000000))
	// 	e.logger.Infof("calculated amount to send to buyer: %s", fl)
	// }

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
		BuyerEscrowWallet: EscrowWallet{
			PublicAddress: order.BuyerEscrowWalletPublicKey,
			PrivateKey:    order.BuyerEscrowWalletPrivateKey,
			Chain:         order.TradeAsset,
		},
	}

	buyersAccountWatchRequest := &AccountWatchRequest{}
	sellersAccountWatchRequest := &AccountWatchRequest{}

	switch order.Currency {
	case SOL:
		// generate a new solana account for the buyer
		acc := createSolanaAccount()

		co.SellerEscrowWallet = EscrowWallet{
			PublicAddress: acc.PublicKey,
			PrivateKey:    acc.PrivateKey,
			Chain:         SOL,
		}

		if err := notifySellerOfBuyer(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		sellersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.SellerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         SOL,
			Amount:        co.Amount,
			TransactionID: co.SellerNKNAddress,
			Seller:        true,
		}

	case CEL:
		acc := generateEVMAccount()
		co.SellerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         CEL,
		}

		if err := notifySellerOfBuyer(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		sellersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.SellerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         CEL,
			Amount:        co.Amount,
			TransactionID: co.SellerNKNAddress,
			Seller:        true,
		}

	case ETH:
		acc := generateEVMAccount()
		co.SellerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         ETH,
		}

		if err := notifySellerOfBuyer(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		sellersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.SellerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         ETH,
			Amount:        co.Amount,
			TransactionID: co.SellerNKNAddress,
			Seller:        true,
		}

	case POL:
		acc := generateEVMAccount()
		co.SellerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         POL,
		}

		if err := notifySellerOfBuyer(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		sellersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.SellerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         POL,
			Amount:        co.Amount,
			TransactionID: co.SellerNKNAddress,
			Seller:        true,
		}

	case MO:
		acc := generateEVMAccount()
		co.SellerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         MO,
		}

		if err := notifySellerOfBuyer(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		sellersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.SellerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         MO,
			Amount:        co.Amount,
			TransactionID: co.SellerNKNAddress,
			Seller:        true,
		}
	default:
		return errors.New("invalid currency")
	}

	switch ta {
	case SOL:
		// acc := createSolanaAccount()
		// co.BuyerEscrowWallet = EscrowWallet{
		// 	PublicAddress: acc.PublicKey,
		// 	PrivateKey:    acc.PrivateKey,
		// 	Chain:         SOL,
		// }

		// if err := sendBuyerPayInfo(*co); err != nil {
		// 	// TODO:: Cancle the order
		// 	return err
		// }

		buyersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.BuyerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         SOL,
			Amount:        co.Price,
			Seller:        false,
			TransactionID: co.SellerNKNAddress,
		}

	case MO:
		// acc := generateEVMAccount()
		// co.BuyerEscrowWallet = EscrowWallet{
		// 	ECDSA:         acc,
		// 	PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
		// 	PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
		// 	Chain:         MO,
		// }

		if err := sendBuyerPayInfo(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}
		buyersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.BuyerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         MO,
			Amount:        co.Price,
			Seller:        false,
			TransactionID: co.SellerNKNAddress,
		}

	case ETH:
		// acc := generateEVMAccount()
		// co.BuyerEscrowWallet = EscrowWallet{
		// 	ECDSA:         acc,
		// 	PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
		// 	PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
		// 	Chain:         ETH,
		// }

		if err := sendBuyerPayInfo(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		buyersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.BuyerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         ETH,
			Amount:        co.Price,
			Seller:        false,
			TransactionID: co.SellerNKNAddress,
		}
	case CEL:
		// acc := generateEVMAccount()
		// co.BuyerEscrowWallet = EscrowWallet{
		// 	ECDSA:         acc,
		// 	PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
		// 	PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
		// 	Chain:         CEL,
		// }

		if err := sendBuyerPayInfo(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		buyersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.BuyerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         CEL,
			Amount:        co.Price,
			Seller:        false,
			TransactionID: co.SellerNKNAddress,
		}

	case POL:
		// acc := generateEVMAccount()
		// co.BuyerEscrowWallet = EscrowWallet{
		// 	ECDSA:         acc,
		// 	PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
		// 	PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
		// 	Chain:         POL,
		// }

		if err := sendBuyerPayInfo(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		// emit a new event to let Warren know that we need to start watching a new account
		buyersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.BuyerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         POL,
			Amount:        co.Price,
			Seller:        false,
			TransactionID: co.SellerNKNAddress,
		}
	default:
		return errors.New("invalid currency")
	}

	am.keeper.SetOrdersUnderWatch(ctx, partyTypes.OrdersUnderWatch{
		Index:            co.BuyerNKNAddress,
		NknAddress:       co.BuyerNKNAddress,
		WalletPrivateKey: co.BuyerEscrowWallet.PrivateKey,
		WalletPublicKey:  co.BuyerEscrowWallet.PublicAddress,
		ShippingAddress:  co.BuyerShippingAddress,
		Amount:           order.Price,
		Chain:            buyersAccountWatchRequest.Chain,
		PaymentComplete:  false,
	})

	am.keeper.SetOrdersUnderWatch(ctx, partyTypes.OrdersUnderWatch{
		Index:            co.SellerNKNAddress,
		NknAddress:       co.SellerNKNAddress,
		WalletPrivateKey: co.SellerEscrowWallet.PrivateKey,
		WalletPublicKey:  co.SellerEscrowWallet.PublicAddress,
		ShippingAddress:  co.SellerShippingAddress,
		Amount:           order.Amount,
		Chain:            sellersAccountWatchRequest.Chain,
		PaymentComplete:  false,
	})

	go am.watchAccount(ctx, buyersAccountWatchRequest)
	go am.watchAccount(ctx, sellersAccountWatchRequest)
	return nil
}
