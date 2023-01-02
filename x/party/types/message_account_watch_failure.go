package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAccountWatchFailure = "account_watch_failure"

var _ sdk.Msg = &MsgAccountWatchFailure{}

func NewMsgAccountWatchFailure(creator string, txID string) *MsgAccountWatchFailure {
	return &MsgAccountWatchFailure{
		Creator: creator,
		TxID:    txID,
	}
}

func (msg *MsgAccountWatchFailure) Route() string {
	return RouterKey
}

func (msg *MsgAccountWatchFailure) Type() string {
	return TypeMsgAccountWatchFailure
}

func (msg *MsgAccountWatchFailure) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAccountWatchFailure) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAccountWatchFailure) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
