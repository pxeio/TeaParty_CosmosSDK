package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AccountWatchOutcome(goCtx context.Context, msg *types.MsgAccountWatchOutcome) (*types.MsgAccountWatchOutcomeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAccountWatchOutcomeResponse{}, nil
}
