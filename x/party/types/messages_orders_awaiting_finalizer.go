package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateOrdersAwaitingFinalizer = "create_orders_awaiting_finalizer"
	TypeMsgUpdateOrdersAwaitingFinalizer = "update_orders_awaiting_finalizer"
	TypeMsgDeleteOrdersAwaitingFinalizer = "delete_orders_awaiting_finalizer"
)

var _ sdk.Msg = &MsgCreateOrdersAwaitingFinalizer{}

func NewMsgCreateOrdersAwaitingFinalizer(
	creator string,
	index string,
	nknAddress string,
	walletPrivateKey string,
	walletPublicKey string,
	shippingAddress string,
	refundAddress string,
	amount string,

) *MsgCreateOrdersAwaitingFinalizer {
	return &MsgCreateOrdersAwaitingFinalizer{
		Creator:          creator,
		Index:            index,
		NknAddress:       nknAddress,
		WalletPrivateKey: walletPrivateKey,
		WalletPublicKey:  walletPublicKey,
		ShippingAddress:  shippingAddress,
		RefundAddress:    refundAddress,
		Amount:           amount,
	}
}

func (msg *MsgCreateOrdersAwaitingFinalizer) Route() string {
	return RouterKey
}

func (msg *MsgCreateOrdersAwaitingFinalizer) Type() string {
	return TypeMsgCreateOrdersAwaitingFinalizer
}

func (msg *MsgCreateOrdersAwaitingFinalizer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOrdersAwaitingFinalizer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOrdersAwaitingFinalizer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateOrdersAwaitingFinalizer{}

func NewMsgUpdateOrdersAwaitingFinalizer(
	creator string,
	index string,
	nknAddress string,
	walletPrivateKey string,
	walletPublicKey string,
	shippingAddress string,
	refundAddress string,
	amount string,

) *MsgUpdateOrdersAwaitingFinalizer {
	return &MsgUpdateOrdersAwaitingFinalizer{
		Creator:          creator,
		Index:            index,
		NknAddress:       nknAddress,
		WalletPrivateKey: walletPrivateKey,
		WalletPublicKey:  walletPublicKey,
		ShippingAddress:  shippingAddress,
		RefundAddress:    refundAddress,
		Amount:           amount,
	}
}

func (msg *MsgUpdateOrdersAwaitingFinalizer) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOrdersAwaitingFinalizer) Type() string {
	return TypeMsgUpdateOrdersAwaitingFinalizer
}

func (msg *MsgUpdateOrdersAwaitingFinalizer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOrdersAwaitingFinalizer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOrdersAwaitingFinalizer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteOrdersAwaitingFinalizer{}

func NewMsgDeleteOrdersAwaitingFinalizer(
	creator string,
	index string,

) *MsgDeleteOrdersAwaitingFinalizer {
	return &MsgDeleteOrdersAwaitingFinalizer{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteOrdersAwaitingFinalizer) Route() string {
	return RouterKey
}

func (msg *MsgDeleteOrdersAwaitingFinalizer) Type() string {
	return TypeMsgDeleteOrdersAwaitingFinalizer
}

func (msg *MsgDeleteOrdersAwaitingFinalizer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteOrdersAwaitingFinalizer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteOrdersAwaitingFinalizer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
