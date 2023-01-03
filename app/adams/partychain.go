package adams

import (
	"fmt"

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
	// update the payment on the party chain
	ctx := sdk.UnwrapSDKContext(goctx)
	po := e.PartyKeeper.GetAllPendingOrders(ctx)
	for i, p := range po {
		if p.Index == result.AccountWatchRequest.TransactionID {
			switch result.Result {
			case OUTCOME_SUCCESS:
				if !result.AccountWatchRequest.Seller {
					p.BuyerPaymentComplete = true
				} else {
					p.SellerPaymentComplete = true
				}
			case OUTCOME_FAILURE:
				if !result.AccountWatchRequest.Seller {
					p.BuyerPaymentComplete = false
				} else {
					p.SellerPaymentComplete = false
				}
			case OUTCOME_TIMEOUT:
				if !result.AccountWatchRequest.Seller {
					p.BuyerPaymentComplete = false
				} else {
					p.SellerPaymentComplete = false
				}
			}
			e.PartyKeeper.RemovePendingOrders(goctx, "party13cc8tgtcaa9ag36cn85zhzw90t6hhly8mfuavu")
			fmt.Printf("po %+v", po[i])
			// e.PartyKeeper.SetPendingOrders(goctx, po[i])
			e.logger.Info("Updated pending order " + result.AccountWatchRequest.TransactionID + " on party chain")
		}
	}

	return nil
}
