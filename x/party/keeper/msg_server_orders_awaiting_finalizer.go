package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateOrdersAwaitingFinalizer(goCtx context.Context, msg *types.MsgCreateOrdersAwaitingFinalizer) (*types.MsgCreateOrdersAwaitingFinalizerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetOrdersAwaitingFinalizer(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var ordersAwaitingFinalizer = types.OrdersAwaitingFinalizer{
		Creator:          msg.Creator,
		Index:            msg.Index,
		NknAddress:       msg.NknAddress,
		WalletPrivateKey: msg.WalletPrivateKey,
		WalletPublicKey:  msg.WalletPublicKey,
		ShippingAddress:  msg.ShippingAddress,
		RefundAddress:    msg.RefundAddress,
		Amount:           msg.Amount,
	}

	k.SetOrdersAwaitingFinalizer(
		ctx,
		ordersAwaitingFinalizer,
	)
	return &types.MsgCreateOrdersAwaitingFinalizerResponse{}, nil
}

func (k msgServer) UpdateOrdersAwaitingFinalizer(goCtx context.Context, msg *types.MsgUpdateOrdersAwaitingFinalizer) (*types.MsgUpdateOrdersAwaitingFinalizerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetOrdersAwaitingFinalizer(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var ordersAwaitingFinalizer = types.OrdersAwaitingFinalizer{
		Creator:          msg.Creator,
		Index:            msg.Index,
		NknAddress:       msg.NknAddress,
		WalletPrivateKey: msg.WalletPrivateKey,
		WalletPublicKey:  msg.WalletPublicKey,
		ShippingAddress:  msg.ShippingAddress,
		RefundAddress:    msg.RefundAddress,
		Amount:           msg.Amount,
	}

	k.SetOrdersAwaitingFinalizer(ctx, ordersAwaitingFinalizer)

	return &types.MsgUpdateOrdersAwaitingFinalizerResponse{}, nil
}

func (k msgServer) DeleteOrdersAwaitingFinalizer(goCtx context.Context, msg *types.MsgDeleteOrdersAwaitingFinalizer) (*types.MsgDeleteOrdersAwaitingFinalizerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetOrdersAwaitingFinalizer(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveOrdersAwaitingFinalizer(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteOrdersAwaitingFinalizerResponse{}, nil
}
