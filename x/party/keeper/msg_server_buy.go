package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k msgServer) Buy(goCtx context.Context, msg *types.MsgBuy) (*types.MsgBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	// atempt to find a matching sell order
	// if found, start the trade
	// if not found, deny the buy order
	tradeOrder, found := k.GetTradeOrders(ctx, msg.TxID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No matching sell order found.")
	}

	// create a new escrow wallet for the buyer
	// buyerEscrowWallet := "0x0000000000000000000000000000000000000001"
	buyerPrivateKey, buyerPublicKey, err := generateEVMAccount()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Failed to generate escrow wallet for buyer.")
	}

	// create a new escrow wallet for the seller
	// sellerEscrowWallet := "0x0000000000000000000000000000000000000002"
	sellerPrivateKey, sellerPublicKey, err := generateEVMAccount()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Failed to generate escrow wallet for seller.")
	}

	// create a pending-order object
	po := types.PendingOrders{
		Index:                        tradeOrder.Index,
		BuyerEscrowWalletPublicKey:   buyerPublicKey,
		BuyerEscrowWalletPrivateKey:  buyerPrivateKey,
		SellerEscrowWalletPublicKey:  sellerPublicKey,
		SellerEscrowWalletPrivateKey: sellerPrivateKey,
		SellerPaymentComplete:        false,
		BuyerPaymentComplete:         false,
		Amount:                       tradeOrder.Amount,
		BuyerShippingAddress:         msg.BuyerShippingAddress,
		BuyerRefundAddress:           msg.RefundAddress,
		BuyerNKNAddress:              msg.BuyerNKNAddress,
		SellerRefundAddress:          tradeOrder.RefundAddr,
		SellerShippingAddress:        tradeOrder.SellerShippingAddr,
		SellerNKNAddress:             tradeOrder.SellerNknAddr,
		TradeAsset:                   tradeOrder.TradeAsset,
		Currency:                     tradeOrder.Currency,
		Price:                        tradeOrder.Price,
	}

	// store the pending order
	k.SetPendingOrders(ctx, po)

	// remove the trade order from the store
	k.RemoveTradeOrders(ctx, po.Index)

	return &types.MsgBuyResponse{}, nil
}

// generateEVMAccount generates a new Ethereum account
// returning error, private key, and public address
func generateEVMAccount() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	// convert the private key to a string
	privateKeyString := hexutil.Encode(privateKeyBytes)[2:]
	return privateKeyString, publicKey, err
}
