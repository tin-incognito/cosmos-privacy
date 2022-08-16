package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "privacy/testutil/keeper"
	"privacy/testutil/nullify"
	"privacy/x/privacy/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestOutputCoinQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOutputCoin(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetOutputCoinRequest
		response *types.QueryGetOutputCoinResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetOutputCoinRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetOutputCoinResponse{OutputCoin: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetOutputCoinRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetOutputCoinResponse{OutputCoin: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetOutputCoinRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.OutputCoin(wctx, tc.request)
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

func TestOutputCoinQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOutputCoin(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllOutputCoinRequest {
		return &types.QueryAllOutputCoinRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OutputCoinAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.OutputCoin), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.OutputCoin),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OutputCoinAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.OutputCoin), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.OutputCoin),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.OutputCoinAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.OutputCoin),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.OutputCoinAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
