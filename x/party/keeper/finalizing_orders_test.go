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

func createNFinalizingOrders(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.FinalizingOrders {
	items := make([]types.FinalizingOrders, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetFinalizingOrders(ctx, items[i])
	}
	return items
}

func TestFinalizingOrdersGet(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNFinalizingOrders(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFinalizingOrders(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFinalizingOrdersRemove(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNFinalizingOrders(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFinalizingOrders(ctx,
			item.Index,
		)
		_, found := keeper.GetFinalizingOrders(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestFinalizingOrdersGetAll(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNFinalizingOrders(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFinalizingOrders(ctx)),
	)
}
