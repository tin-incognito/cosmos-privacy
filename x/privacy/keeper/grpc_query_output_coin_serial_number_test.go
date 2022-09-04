package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "privacy/testutil/keeper"
	"privacy/testutil/nullify"
	"privacy/x/privacy/types"
)

func TestOutputCoinSerialNumberQuery(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestOutputCoinSerialNumber(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetOutputCoinSerialNumberRequest
		response *types.QueryGetOutputCoinSerialNumberResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetOutputCoinSerialNumberRequest{},
			response: &types.QueryGetOutputCoinSerialNumberResponse{OutputCoinSerialNumber: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.OutputCoinSerialNumber(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
