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

func (k Keeper) TradeOrdersAll(c context.Context, req *types.QueryAllTradeOrdersRequest) (*types.QueryAllTradeOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tradeOrderss []types.TradeOrders
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	tradeOrdersStore := prefix.NewStore(store, types.KeyPrefix(types.TradeOrdersKeyPrefix))

	pageRes, err := query.Paginate(tradeOrdersStore, req.Pagination, func(key []byte, value []byte) error {
		var tradeOrders types.TradeOrders
		if err := k.cdc.Unmarshal(value, &tradeOrders); err != nil {
			return err
		}

		tradeOrderss = append(tradeOrderss, tradeOrders)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTradeOrdersResponse{TradeOrders: tradeOrderss, Pagination: pageRes}, nil
}

func (k Keeper) TradeOrders(c context.Context, req *types.QueryGetTradeOrdersRequest) (*types.QueryGetTradeOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTradeOrders(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTradeOrdersResponse{TradeOrders: val}, nil
}
