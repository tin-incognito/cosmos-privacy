package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"privacy/testutil/sample"
)

func TestMsgCreateOutputCoinSerialNumber_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateOutputCoinSerialNumber
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateOutputCoinSerialNumber{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateOutputCoinSerialNumber{
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

func TestMsgUpdateOutputCoinSerialNumber_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateOutputCoinSerialNumber
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateOutputCoinSerialNumber{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateOutputCoinSerialNumber{
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

func TestMsgDeleteOutputCoinSerialNumber_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteOutputCoinSerialNumber
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteOutputCoinSerialNumber{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteOutputCoinSerialNumber{
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
