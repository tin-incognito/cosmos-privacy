package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"privacy/testutil/sample"
)

func TestMsgCreateCommitmentIndex_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateCommitmentIndex
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateCommitmentIndex{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateCommitmentIndex{
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

func TestMsgUpdateCommitmentIndex_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateCommitmentIndex
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateCommitmentIndex{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateCommitmentIndex{
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

func TestMsgDeleteCommitmentIndex_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteCommitmentIndex
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteCommitmentIndex{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteCommitmentIndex{
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
