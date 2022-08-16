package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"privacy/testutil/sample"
)

func TestMsgCreateTxPrivacyData_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTxPrivacyData
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTxPrivacyData{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTxPrivacyData{
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

func TestMsgUpdateTxPrivacyData_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTxPrivacyData
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTxPrivacyData{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTxPrivacyData{
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

func TestMsgDeleteTxPrivacyData_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTxPrivacyData
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTxPrivacyData{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTxPrivacyData{
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
