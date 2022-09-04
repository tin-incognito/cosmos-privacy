package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "privacy/testutil/keeper"
	"privacy/x/privacy/keeper"
	"privacy/x/privacy/types"
)

func TestOutputCoinSerialNumberMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.PrivacyKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateOutputCoinSerialNumber{Creator: creator}
	_, err := srv.CreateOutputCoinSerialNumber(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetOutputCoinSerialNumber(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestOutputCoinSerialNumberMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateOutputCoinSerialNumber
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateOutputCoinSerialNumber{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateOutputCoinSerialNumber{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PrivacyKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateOutputCoinSerialNumber{Creator: creator}
			_, err := srv.CreateOutputCoinSerialNumber(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateOutputCoinSerialNumber(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetOutputCoinSerialNumber(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestOutputCoinSerialNumberMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteOutputCoinSerialNumber
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteOutputCoinSerialNumber{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteOutputCoinSerialNumber{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PrivacyKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateOutputCoinSerialNumber(wctx, &types.MsgCreateOutputCoinSerialNumber{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteOutputCoinSerialNumber(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetOutputCoinSerialNumber(ctx)
				require.False(t, found)
			}
		})
	}
}
