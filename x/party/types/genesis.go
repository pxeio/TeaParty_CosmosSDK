package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TradeOrdersList:   []TradeOrders{},
		PendingOrdersList: []PendingOrders{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in tradeOrders
	tradeOrdersIndexMap := make(map[string]struct{})

	for _, elem := range gs.TradeOrdersList {
		index := string(TradeOrdersKey(elem.Index))
		if _, ok := tradeOrdersIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tradeOrders")
		}
		tradeOrdersIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pendingOrders
	pendingOrdersIndexMap := make(map[string]struct{})

	for _, elem := range gs.PendingOrdersList {
		index := string(PendingOrdersKey(elem.Index))
		if _, ok := pendingOrdersIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pendingOrders")
		}
		pendingOrdersIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
