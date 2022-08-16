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

func TestCommitmentIndexQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCommitmentIndex(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCommitmentIndexRequest
		response *types.QueryGetCommitmentIndexResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetCommitmentIndexRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetCommitmentIndexResponse{CommitmentIndex: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetCommitmentIndexRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetCommitmentIndexResponse{CommitmentIndex: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetCommitmentIndexRequest{
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
			response, err := keeper.CommitmentIndex(wctx, tc.request)
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

func TestCommitmentIndexQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCommitmentIndex(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCommitmentIndexRequest {
		return &types.QueryAllCommitmentIndexRequest{
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
			resp, err := keeper.CommitmentIndexAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CommitmentIndex), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.CommitmentIndex),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CommitmentIndexAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CommitmentIndex), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.CommitmentIndex),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CommitmentIndexAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.CommitmentIndex),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CommitmentIndexAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
