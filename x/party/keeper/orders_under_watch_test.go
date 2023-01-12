package keeper_test

import (
	"strconv"
	"testing"

	"github.com/TeaPartyCrypto/partychain/x/party/keeper"
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	keepertest "github.com/TeaPartyCrypto/partychain/testutil/keeper"
	"github.com/TeaPartyCrypto/partychain/testutil/nullify"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNOrdersUnderWatch(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.OrdersUnderWatch {
	items := make([]types.OrdersUnderWatch, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
        
		keeper.SetOrdersUnderWatch(ctx, items[i])
	}
	return items
}

func TestOrdersUnderWatchGet(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNOrdersUnderWatch(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetOrdersUnderWatch(ctx,
		    item.Index,
            
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestOrdersUnderWatchRemove(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNOrdersUnderWatch(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOrdersUnderWatch(ctx,
		    item.Index,
            
		)
		_, found := keeper.GetOrdersUnderWatch(ctx,
		    item.Index,
            
		)
		require.False(t, found)
	}
}

func TestOrdersUnderWatchGetAll(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	items := createNOrdersUnderWatch(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOrdersUnderWatch(ctx)),
	)
}
