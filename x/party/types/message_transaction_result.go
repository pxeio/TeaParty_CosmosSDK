package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransactionResult = "transaction_result"

var _ sdk.Msg = &MsgTransactionResult{}

func NewMsgTransactionResult(creator string, txID string, outcome string) *MsgTransactionResult {
	return &MsgTransactionResult{
		Creator: creator,
		TxID:    txID,
		Outcome: outcome,
	}
}

func (msg *MsgTransactionResult) Route() string {
	return RouterKey
}

func (msg *MsgTransactionResult) Type() string {
	return TypeMsgTransactionResult
}

func (msg *MsgTransactionResult) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransactionResult) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransactionResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
