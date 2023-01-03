package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrdersAwaitingFinalizerAll(c context.Context, req *types.QueryAllOrdersAwaitingFinalizerRequest) (*types.QueryAllOrdersAwaitingFinalizerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ordersAwaitingFinalizers []types.OrdersAwaitingFinalizer
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ordersAwaitingFinalizerStore := prefix.NewStore(store, types.KeyPrefix(types.OrdersAwaitingFinalizerKeyPrefix))

	pageRes, err := query.Paginate(ordersAwaitingFinalizerStore, req.Pagination, func(key []byte, value []byte) error {
		var ordersAwaitingFinalizer types.OrdersAwaitingFinalizer
		if err := k.cdc.Unmarshal(value, &ordersAwaitingFinalizer); err != nil {
			return err
		}

		ordersAwaitingFinalizers = append(ordersAwaitingFinalizers, ordersAwaitingFinalizer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrdersAwaitingFinalizerResponse{OrdersAwaitingFinalizer: ordersAwaitingFinalizers, Pagination: pageRes}, nil
}

func (k Keeper) OrdersAwaitingFinalizer(c context.Context, req *types.QueryGetOrdersAwaitingFinalizerRequest) (*types.QueryGetOrdersAwaitingFinalizerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOrdersAwaitingFinalizer(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOrdersAwaitingFinalizerResponse{OrdersAwaitingFinalizer: val}, nil
}
