package adams

import (
	"encoding/json"
	"net/http"
)

// FetchTradeOrders is a http route handler that returns all the sell orders
func (e *ExchangeServer) FetchTradeOrders(w http.ResponseWriter, r *http.Request) {
	// update the state of the orders from the party chain
	orders, err := fetchTradeOrdersFromPartyChain(e.partyNode)
	if err != nil {
		e.logger.Error("failed to fetch orders from party chain")
	}
	e.partyChainOrders = orders

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e.partyChainOrders)
}
