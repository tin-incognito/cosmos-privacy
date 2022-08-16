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

func TestTxPrivacyDataQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTxPrivacyData(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTxPrivacyDataRequest
		response *types.QueryGetTxPrivacyDataResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTxPrivacyDataRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetTxPrivacyDataResponse{TxPrivacyData: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTxPrivacyDataRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetTxPrivacyDataResponse{TxPrivacyData: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTxPrivacyDataRequest{
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
			response, err := keeper.TxPrivacyData(wctx, tc.request)
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

func TestTxPrivacyDataQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTxPrivacyData(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTxPrivacyDataRequest {
		return &types.QueryAllTxPrivacyDataRequest{
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
			resp, err := keeper.TxPrivacyDataAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TxPrivacyData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TxPrivacyData),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TxPrivacyDataAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TxPrivacyData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TxPrivacyData),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TxPrivacyDataAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.TxPrivacyData),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TxPrivacyDataAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
