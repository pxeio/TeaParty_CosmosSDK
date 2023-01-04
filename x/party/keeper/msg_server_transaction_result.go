package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransactionResult(goCtx context.Context, msg *types.MsgTransactionResult) (*types.MsgTransactionResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// remove the order from the list of accounts awaiting finalization
	for _, order := range k.GetAllOrdersAwaitingFinalizer(ctx) {
		if order.Index == msg.TxID {
			k.RemoveOrdersAwaitingFinalizer(ctx, order.Index)
			break
		}
	}
	return &types.MsgTransactionResultResponse{}, nil
}
