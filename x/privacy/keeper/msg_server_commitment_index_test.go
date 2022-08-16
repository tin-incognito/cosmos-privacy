package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "privacy/testutil/keeper"
	"privacy/x/privacy/keeper"
	"privacy/x/privacy/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCommitmentIndexMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.PrivacyKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateCommitmentIndex{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateCommitmentIndex(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetCommitmentIndex(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestCommitmentIndexMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateCommitmentIndex
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateCommitmentIndex{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateCommitmentIndex{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateCommitmentIndex{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PrivacyKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateCommitmentIndex{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateCommitmentIndex(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateCommitmentIndex(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetCommitmentIndex(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestCommitmentIndexMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteCommitmentIndex
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteCommitmentIndex{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteCommitmentIndex{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteCommitmentIndex{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PrivacyKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateCommitmentIndex(wctx, &types.MsgCreateCommitmentIndex{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteCommitmentIndex(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetCommitmentIndex(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
