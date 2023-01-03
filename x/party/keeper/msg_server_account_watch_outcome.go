package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	OUTCOME_SUCCESS = "success"
	OUTCOME_FAILURE = "failure"
	OUTCOME_TIMEOUT = "timeout"
)

func (k msgServer) AccountWatchOutcome(goCtx context.Context, msg *types.MsgAccountWatchOutcome) (*types.MsgAccountWatchOutcomeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	po := k.GetAllPendingOrders(ctx)
	for _, p := range po {
		if p.Index == msg.TxID {
			switch msg.PaymentOutcome {
			case OUTCOME_SUCCESS:
				if msg.Buyer {
					p.BuyerPaymentComplete = true
				} else {
					p.SellerPaymentComplete = true
				}
			case OUTCOME_FAILURE:
				if msg.Buyer {
					p.BuyerPaymentComplete = false
				} else {
					p.SellerPaymentComplete = false
				}
			case OUTCOME_TIMEOUT:
				if msg.Buyer {
					p.BuyerPaymentComplete = false
				} else {
					p.SellerPaymentComplete = false
				}
			}

			k.RemovePendingOrders(ctx, p.Index)
			k.SetPendingOrders(ctx, p)
		}
	}
	return &types.MsgAccountWatchOutcomeResponse{}, nil
}
