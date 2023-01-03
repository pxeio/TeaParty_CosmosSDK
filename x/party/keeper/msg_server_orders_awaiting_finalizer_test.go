package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/TeaPartyCrypto/partychain/testutil/keeper"
	"github.com/TeaPartyCrypto/partychain/x/party/keeper"
	"github.com/TeaPartyCrypto/partychain/x/party/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestOrdersAwaitingFinalizerMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.PartyKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateOrdersAwaitingFinalizer{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateOrdersAwaitingFinalizer(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetOrdersAwaitingFinalizer(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestOrdersAwaitingFinalizerMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateOrdersAwaitingFinalizer
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateOrdersAwaitingFinalizer{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateOrdersAwaitingFinalizer{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateOrdersAwaitingFinalizer{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PartyKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateOrdersAwaitingFinalizer{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateOrdersAwaitingFinalizer(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateOrdersAwaitingFinalizer(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetOrdersAwaitingFinalizer(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestOrdersAwaitingFinalizerMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteOrdersAwaitingFinalizer
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteOrdersAwaitingFinalizer{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteOrdersAwaitingFinalizer{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteOrdersAwaitingFinalizer{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PartyKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateOrdersAwaitingFinalizer(wctx, &types.MsgCreateOrdersAwaitingFinalizer{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteOrdersAwaitingFinalizer(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetOrdersAwaitingFinalizer(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
