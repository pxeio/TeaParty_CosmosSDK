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

func createNOrdersAwaitingFinalizer(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.OrdersAwaitingFinalizer {
	items := make([]types.OrdersAwaitingFinalizer, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetOrdersAwaitingFinalizer(ctx, items[i])
	}
	return items
}

func TestOrdersAwaitingFinalizerGet(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNOrdersAwaitingFinalizer(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetOrdersAwaitingFinalizer(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestOrdersAwaitingFinalizerRemove(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNOrdersAwaitingFinalizer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOrdersAwaitingFinalizer(ctx,
			item.Index,
		)
		_, found := keeper.GetOrdersAwaitingFinalizer(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestOrdersAwaitingFinalizerGetAll(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNOrdersAwaitingFinalizer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOrdersAwaitingFinalizer(ctx)),
	)
}
