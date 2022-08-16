package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"privacy/testutil/sample"
)

func TestMsgCreateSerialNumber_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateSerialNumber
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateSerialNumber{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateSerialNumber{
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

func TestMsgUpdateSerialNumber_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateSerialNumber
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSerialNumber{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSerialNumber{
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

func TestMsgDeleteSerialNumber_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteSerialNumber
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteSerialNumber{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteSerialNumber{
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
