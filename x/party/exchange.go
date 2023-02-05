package party

import (
	"context"
	"fmt"
	"math/big"

	partyTypes "github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/ethclient"
	solRPC "github.com/gagliardetto/solana-go/rpc"
	"github.com/shopspring/decimal"
)

func (am AppModule) dispatch(ctx sdk.Context, awrr *AccountWatchRequestResult) {
	ouw, ok := am.keeper.GetOrdersUnderWatch(ctx, awrr.AccountWatchRequest.TransactionID)
	if !ok {
		fmt.Println("no orders under watch")
		return
	}

	fmt.Printf("looking for pending order for %+v", ouw)

	pendingFinalizingOrder, ok := am.keeper.GetFinalizingOrders(ctx, ouw.Index)
	if !ok {
		fmt.Println("no order in finalizing orders")
		return
	}

	switch awrr.Result {
	case OUTCOME_SUCCESS:
		fmt.Println("success")
		// TODO:: remove this
		ouw.PaymentComplete = true

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
		ouw.PaymentComplete = false

		if ouw.Chain == pendingFinalizingOrder.Currency {
			pendingFinalizingOrder.BuyerPaymentComplete = false
		}
		if ouw.Chain == pendingFinalizingOrder.TradeAsset {
			pendingFinalizingOrder.SellerPaymentComplete = false
		}
	case OUTCOME_TIMEOUT:
		// TODO:: remove this
		ouw.PaymentComplete = false

		if ouw.Chain == pendingFinalizingOrder.Currency {
			pendingFinalizingOrder.BuyerPaymentComplete = false
		}
		if ouw.Chain == pendingFinalizingOrder.TradeAsset {
			pendingFinalizingOrder.SellerPaymentComplete = false
		}
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
			Amount:           pendingFinalizingOrder.Amount,
			Chain:            pendingFinalizingOrder.Currency,
		}

		soaf := partyTypes.OrdersAwaitingFinalizer{
			Index:            pendingFinalizingOrder.SellerEscrowWalletPublicKey,
			NknAddress:       pendingFinalizingOrder.SellerNKNAddress,
			WalletPrivateKey: pendingFinalizingOrder.BuyerEscrowWalletPrivateKey,
			WalletPublicKey:  pendingFinalizingOrder.BuyerEscrowWalletPublicKey,
			ShippingAddress:  pendingFinalizingOrder.SellerShippingAddress,
			RefundAddress:    pendingFinalizingOrder.SellerRefundAddress,
			Amount:           pendingFinalizingOrder.Amount,
			Chain:            pendingFinalizingOrder.TradeAsset,
		}

		fmt.Printf("buyer order awaiting finalizer: %+v ", boaf)
		fmt.Println("------")
		fmt.Printf("seller order awaiting finalizer: %+v ", soaf)

		am.keeper.SetOrdersAwaitingFinalizer(ctx, boaf)
		am.keeper.SetOrdersAwaitingFinalizer(ctx, soaf)
		// am.keeper.RemovePendingOrders(ctx, pendingFinalizingOrder.Index)
		am.keeper.RemoveOrdersUnderWatch(ctx, ouw.Index)
		am.wg.Done()
		return
	} else {
		// remove the order under watch
		am.keeper.RemoveOrdersUnderWatch(ctx, ouw.Index)
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
		fo := am.keeper.GetAllFinalizingOrders(ctx)
		for _, o := range fo {
			if o.SellerNKNAddress == order.NknAddress || o.BuyerNKNAddress == order.NknAddress {
				am.keeper.RemoveFinalizingOrders(ctx, o.Index)
			}
		}

	}

	// TODO:: notify the party chain that the order has been finished and the funds have been sent
	// and confirmed

	// if err := notifyPartyChainOfTransactionResult(order.Index, "success"); err != nil {
	// 	e.logger.Error("error notifying the party chain that the transaction was successful: " + err.Error())
	// }

	am.keeper.RemoveOrdersAwaitingFinalizer(ctx, order.Index)
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
