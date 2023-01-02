package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitSell(goCtx context.Context, msg *types.MsgSubmitSell) (*types.MsgSubmitSellResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitSellResponse{}, nil
}
