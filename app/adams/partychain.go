package adams

import (
	partyTypes "github.com/TeaPartyCrypto/partychain/x/party/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

// notifyPartyChainOfWatchResult calls the update-payment method on the TeaParty chain.
func (e *ExchangeServer) notifyPartyChainOfWatchResult(result *AccountWatchRequestResult) error {
	return nil
}
