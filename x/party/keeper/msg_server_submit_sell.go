package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SubmitSell(goCtx context.Context, msg *types.MsgSubmitSell) (*types.MsgSubmitSellResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// atempt to find a open sell order from the same seller
	// if found, deny the sell order
	// if not found, create a new sell order
	order, found := k.GetTradeOrders(ctx, msg.Creator)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Account already has an existing open sell order.")
	}

	// create a new sell order
	order = types.TradeOrders{
		Index:              msg.SellerNknAddr,
		TradeAsset:         msg.TradeAsset,
		Price:              msg.Price,
		Currency:           msg.Currency,
		Amount:             msg.Amount,
		SellerShippingAddr: msg.SellerShippingAddr,
		SellerNknAddr:      msg.SellerNknAddr,
		RefundAddr:         msg.RefundAddr,
	}

	// store the sell order
	k.SetTradeOrders(ctx, order)
	return &types.MsgSubmitSellResponse{}, nil
}
