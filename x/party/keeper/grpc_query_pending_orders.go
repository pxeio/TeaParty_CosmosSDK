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

func (k Keeper) PendingOrdersAll(c context.Context, req *types.QueryAllPendingOrdersRequest) (*types.QueryAllPendingOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pendingOrderss []types.PendingOrders
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pendingOrdersStore := prefix.NewStore(store, types.KeyPrefix(types.PendingOrdersKeyPrefix))

	pageRes, err := query.Paginate(pendingOrdersStore, req.Pagination, func(key []byte, value []byte) error {
		var pendingOrders types.PendingOrders
		if err := k.cdc.Unmarshal(value, &pendingOrders); err != nil {
			return err
		}

		pendingOrderss = append(pendingOrderss, pendingOrders)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPendingOrdersResponse{PendingOrders: pendingOrderss, Pagination: pageRes}, nil
}

func (k Keeper) PendingOrders(c context.Context, req *types.QueryGetPendingOrdersRequest) (*types.QueryGetPendingOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPendingOrders(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPendingOrdersResponse{PendingOrders: val}, nil
}
