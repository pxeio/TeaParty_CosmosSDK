package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/TeaPartyCrypto/partychain/testutil/keeper"
	"github.com/TeaPartyCrypto/partychain/testutil/nullify"
	"github.com/TeaPartyCrypto/partychain/x/party/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPendingOrdersQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPendingOrders(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPendingOrdersRequest
		response *types.QueryGetPendingOrdersResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPendingOrdersRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetPendingOrdersResponse{PendingOrders: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPendingOrdersRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetPendingOrdersResponse{PendingOrders: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPendingOrdersRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PendingOrders(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestPendingOrdersQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PartyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPendingOrders(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPendingOrdersRequest {
		return &types.QueryAllPendingOrdersRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PendingOrdersAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PendingOrders), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PendingOrders),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PendingOrdersAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PendingOrders), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PendingOrders),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PendingOrdersAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.PendingOrders),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PendingOrdersAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
