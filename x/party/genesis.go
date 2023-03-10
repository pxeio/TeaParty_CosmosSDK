package party

import (
	"github.com/TeaPartyCrypto/partychain/x/party/keeper"
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the tradeOrders
	for _, elem := range genState.TradeOrdersList {
		k.SetTradeOrders(ctx, elem)
	}
	// Set all the pendingOrders
	for _, elem := range genState.PendingOrdersList {
		k.SetPendingOrders(ctx, elem)
	}
	// Set all the ordersAwaitingFinalizer
	for _, elem := range genState.OrdersAwaitingFinalizerList {
		k.SetOrdersAwaitingFinalizer(ctx, elem)
	}
	// Set all the ordersUnderWatch
	for _, elem := range genState.OrdersUnderWatchList {
		k.SetOrdersUnderWatch(ctx, elem)
	}
	// Set all the finalizingOrders
	for _, elem := range genState.FinalizingOrdersList {
		k.SetFinalizingOrders(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TradeOrdersList = k.GetAllTradeOrders(ctx)
	genesis.PendingOrdersList = k.GetAllPendingOrders(ctx)
	genesis.OrdersAwaitingFinalizerList = k.GetAllOrdersAwaitingFinalizer(ctx)
	genesis.OrdersUnderWatchList = k.GetAllOrdersUnderWatch(ctx)
	genesis.FinalizingOrdersList = k.GetAllFinalizingOrders(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
