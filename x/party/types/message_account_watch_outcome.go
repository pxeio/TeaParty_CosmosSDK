package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAccountWatchOutcome = "account_watch_outcome"

var _ sdk.Msg = &MsgAccountWatchOutcome{}

func NewMsgAccountWatchOutcome(creator string, txID string, buyer bool, paymentOutcome string) *MsgAccountWatchOutcome {
	return &MsgAccountWatchOutcome{
		Creator:        creator,
		TxID:           txID,
		Buyer:          buyer,
		PaymentOutcome: paymentOutcome,
	}
}

func (msg *MsgAccountWatchOutcome) Route() string {
	return RouterKey
}

func (msg *MsgAccountWatchOutcome) Type() string {
	return TypeMsgAccountWatchOutcome
}

func (msg *MsgAccountWatchOutcome) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAccountWatchOutcome) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAccountWatchOutcome) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
