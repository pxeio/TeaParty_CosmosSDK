package types

import (
	"testing"

	"github.com/TeaPartyCrypto/partychain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateOrdersAwaitingFinalizer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateOrdersAwaitingFinalizer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateOrdersAwaitingFinalizer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateOrdersAwaitingFinalizer{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateOrdersAwaitingFinalizer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateOrdersAwaitingFinalizer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateOrdersAwaitingFinalizer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateOrdersAwaitingFinalizer{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteOrdersAwaitingFinalizer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteOrdersAwaitingFinalizer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteOrdersAwaitingFinalizer{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteOrdersAwaitingFinalizer{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
