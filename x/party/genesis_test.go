package party_test

import (
	"testing"

	keepertest "github.com/TeaPartyCrypto/partychain/testutil/keeper"
	"github.com/TeaPartyCrypto/partychain/testutil/nullify"
	"github.com/TeaPartyCrypto/partychain/x/party"
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TradeOrdersList: []types.TradeOrders{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		PendingOrdersList: []types.PendingOrders{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		OrdersAwaitingFinalizerList: []types.OrdersAwaitingFinalizer{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		OrdersUnderWatchList: []types.OrdersUnderWatch{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PartyKeeper(t)
	party.InitGenesis(ctx, *k, genesisState)
	got := party.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TradeOrdersList, got.TradeOrdersList)
	require.ElementsMatch(t, genesisState.PendingOrdersList, got.PendingOrdersList)
	require.ElementsMatch(t, genesisState.OrdersAwaitingFinalizerList, got.OrdersAwaitingFinalizerList)
	require.ElementsMatch(t, genesisState.OrdersUnderWatchList, got.OrdersUnderWatchList)
	// this line is used by starport scaffolding # genesis/test/assert
}
