package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/TeaPartyCrypto/partychain/testutil/keeper"
	"github.com/TeaPartyCrypto/partychain/testutil/nullify"
	"github.com/TeaPartyCrypto/partychain/x/party/keeper"
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPendingOrders(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PendingOrders {
	items := make([]types.PendingOrders, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPendingOrders(ctx, items[i])
	}
	return items
}

func TestPendingOrdersGet(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNPendingOrders(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPendingOrders(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPendingOrdersRemove(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNPendingOrders(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePendingOrders(ctx,
			item.Index,
		)
		_, found := keeper.GetPendingOrders(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPendingOrdersGetAll(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNPendingOrders(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPendingOrders(ctx)),
	)
}
