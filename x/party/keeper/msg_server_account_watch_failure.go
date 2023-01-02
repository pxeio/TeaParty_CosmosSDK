package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AccountWatchFailure(goCtx context.Context, msg *types.MsgAccountWatchFailure) (*types.MsgAccountWatchFailureResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAccountWatchFailureResponse{}, nil
}
