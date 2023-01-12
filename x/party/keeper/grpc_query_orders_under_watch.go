package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrdersUnderWatchAll(c context.Context, req *types.QueryAllOrdersUnderWatchRequest) (*types.QueryAllOrdersUnderWatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ordersUnderWatchs []types.OrdersUnderWatch
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	ordersUnderWatchStore := prefix.NewStore(store, types.KeyPrefix(types.OrdersUnderWatchKeyPrefix))

	pageRes, err := query.Paginate(ordersUnderWatchStore, req.Pagination, func(key []byte, value []byte) error {
		var ordersUnderWatch types.OrdersUnderWatch
		if err := k.cdc.Unmarshal(value, &ordersUnderWatch); err != nil {
			return err
		}

		ordersUnderWatchs = append(ordersUnderWatchs, ordersUnderWatch)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrdersUnderWatchResponse{OrdersUnderWatch: ordersUnderWatchs, Pagination: pageRes}, nil
}

func (k Keeper) OrdersUnderWatch(c context.Context, req *types.QueryGetOrdersUnderWatchRequest) (*types.QueryGetOrdersUnderWatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOrdersUnderWatch(
	    ctx,
	    req.Index,
        )
	if !found {
	    return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOrdersUnderWatchResponse{OrdersUnderWatch: val}, nil
}