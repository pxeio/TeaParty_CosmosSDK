package adams

import (
	partyTypes "github.com/TeaPartyCrypto/partychain/x/party/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	OUTCOME_SUCCESS = "success"
	OUTCOME_FAILURE = "failure"
	OUTCOME_TIMEOUT = "timeout"
)

// fetchTradeOrdersFromPartyChain fetches the current list of orders from the TeaParty chain.
func (e *ExchangeServer) fetchTradeOrdersFromPartyChain(ctx sdk.Context) ([]partyTypes.TradeOrders, error) {
	to := e.PartyKeeper.GetAllTradeOrders(ctx)
	return to, nil
}

func (e *ExchangeServer) fetchPendingOrdersFromPartyChain(ctx sdk.Context) ([]partyTypes.PendingOrders, error) {
	po := e.PartyKeeper.GetAllPendingOrders(ctx)
	return po, nil
}

// NotifyPartyChainOfWatchResult calls the update-payment method on the TeaParty chain.
func (e *ExchangeServer) NotifyPartyChainOfWatchResult(result *AccountWatchRequestResult, goctx sdk.Context) error {
	e.logger.Info("Notifying party chain of watch result for " + result.AccountWatchRequest.TransactionID + " on chain " + result.AccountWatchRequest.Chain)

	msg := partyTypes.NewMsgAccountWatchOutcome(result.AccountWatchRequest.TransactionID, result.AccountWatchRequest.Account, result.AccountWatchRequest.Chain, result.AccountWatchRequest.Seller, result.Result)
	msgClient := partyTypes.NewMsgClient(e.partyNode)

	_, err := msgClient.AccountWatchOutcome(goctx, msg)
	if err != nil {
		e.logger.Error("Error notifying party chain of watch result for " + result.AccountWatchRequest.TransactionID + " on chain " + result.AccountWatchRequest.Chain)
		return err
	}

	// update the payment on the party chain
	// ctx := sdk.UnwrapSDKContext(goctx)
	// po := e.PartyKeeper.GetAllPendingOrders(ctx)
	// for _, p := range po {
	// 	if p.Index == result.AccountWatchRequest.TransactionID {
	// 		switch result.Result {
	// 		case OUTCOME_SUCCESS:
	// 			e.logger.Info("Payment successful for " + result.AccountWatchRequest.Account + " on chain " + result.AccountWatchRequest.Chain)
	// 			if !result.AccountWatchRequest.Seller {
	// 				p.BuyerPaymentComplete = true
	// 			} else {
	// 				p.SellerPaymentComplete = true
	// 			}
	// 		case OUTCOME_FAILURE:
	// 			if !result.AccountWatchRequest.Seller {
	// 				p.BuyerPaymentComplete = false
	// 			} else {
	// 				p.SellerPaymentComplete = false
	// 			}
	// 		case OUTCOME_TIMEOUT:
	// 			if !result.AccountWatchRequest.Seller {
	// 				p.BuyerPaymentComplete = false
	// 			} else {
	// 				p.SellerPaymentComplete = false
	// 			}
	// 		}
	// 		e.PartyKeeper.RemovePendingOrders(goctx, p.Index)
	// 		e.logger.Info("Removed pending order " + result.AccountWatchRequest.TransactionID + " from party chain")
	// 		// e.PartyKeeper.SetPendingOrders(goctx, p)
	// 		e.logger.Info("Updated pending order " + result.AccountWatchRequest.TransactionID + " on party chain")
	// 	}
	// }

	return nil
}
