package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitSell = "submit_sell"

var _ sdk.Msg = &MsgSubmitSell{}

func NewMsgSubmitSell(creator string, tradeAsset string, price string, currency string, amount string, sellerShippingAddr string, sellerNknAddr string, refundAddr string) *MsgSubmitSell {
	return &MsgSubmitSell{
		Creator:            creator,
		TradeAsset:         tradeAsset,
		Price:              price,
		Currency:           currency,
		Amount:             amount,
		SellerShippingAddr: sellerShippingAddr,
		SellerNknAddr:      sellerNknAddr,
		RefundAddr:         refundAddr,
	}
}

func (msg *MsgSubmitSell) Route() string {
	return RouterKey
}

func (msg *MsgSubmitSell) Type() string {
	return TypeMsgSubmitSell
}

func (msg *MsgSubmitSell) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitSell) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitSell) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
