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

func (k Keeper) FinalizingOrdersAll(goCtx context.Context, req *types.QueryAllFinalizingOrdersRequest) (*types.QueryAllFinalizingOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var finalizingOrderss []types.FinalizingOrders
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	finalizingOrdersStore := prefix.NewStore(store, types.KeyPrefix(types.FinalizingOrdersKeyPrefix))

	pageRes, err := query.Paginate(finalizingOrdersStore, req.Pagination, func(key []byte, value []byte) error {
		var finalizingOrders types.FinalizingOrders
		if err := k.cdc.Unmarshal(value, &finalizingOrders); err != nil {
			return err
		}

		finalizingOrderss = append(finalizingOrderss, finalizingOrders)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFinalizingOrdersResponse{FinalizingOrders: finalizingOrderss, Pagination: pageRes}, nil
}

func (k Keeper) FinalizingOrders(goCtx context.Context, req *types.QueryGetFinalizingOrdersRequest) (*types.QueryGetFinalizingOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetFinalizingOrders(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFinalizingOrdersResponse{FinalizingOrders: val}, nil
}
