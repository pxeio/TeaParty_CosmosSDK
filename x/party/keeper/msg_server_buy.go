package keeper

import (
	"context"
	"encoding/hex"

	"github.com/TeaPartyCrypto/partychain/x/party/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/gagliardetto/solana-go"
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
	// remove the trade order from the store
	k.RemoveTradeOrders(ctx, tradeOrder.SellerNknAddr)

	var buyerPublicKey string
	var buyerPrivateKey string
	var err error
	// check the chain
	if tradeOrder.TradeAsset == "ethereum" || tradeOrder.TradeAsset == "polygon" || tradeOrder.TradeAsset == "celo" {
		// create a new escrow wallet for the buyer
		err, buyerPrivateKey, buyerPublicKey = generateEVMAccount()
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Failed to generate escrow wallet for buyer.")
		}
	}

	if tradeOrder.TradeAsset == "solana" {
		// create a new escrow wallet for the buyer
		err, buyerPrivateKey, buyerPublicKey = createSolanaAccount()
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Failed to generate escrow wallet for buyer.")
		}
	}

	var sellerPublicKey string
	var sellerPrivateKey string
	// check the chain
	if tradeOrder.TradeAsset == "ethereum" || tradeOrder.TradeAsset == "polygon" || tradeOrder.TradeAsset == "celo" {
		// create a new escrow wallet for the seller
		err, sellerPrivateKey, sellerPublicKey = generateEVMAccount()
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Failed to generate escrow wallet for seller.")
		}
	}
	if tradeOrder.TradeAsset == "solana" {
		// create a new escrow wallet for the seller
		err, sellerPrivateKey, sellerPublicKey = createSolanaAccount()
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Failed to generate escrow wallet for seller.")
		}
	}

	// create a pending-order object
	po := types.PendingOrders{
		Index:                        tradeOrder.SellerNknAddr,
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
	return &types.MsgBuyResponse{}, nil
}

// generateEVMAccount generates a new Ethereum account
// returning error, private key, and public address
func generateEVMAccount() (error, string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return err, "", ""
	}

	// TODO:: fix this again. was lost in a bad git commit
	return nil, hex.EncodeToString(privateKey.D.Bytes()), crypto.PubkeyToAddress(privateKey.PublicKey).String()
}

func createSolanaAccount() (error, string, string) {
	account := solana.NewWallet()

	return nil, account.PrivateKey.String(), account.PublicKey().String()
}
