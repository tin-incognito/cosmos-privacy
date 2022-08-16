package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "privacy/testutil/keeper"
	"privacy/testutil/nullify"
	"privacy/x/privacy/keeper"
	"privacy/x/privacy/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCommitment(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Commitment {
	items := make([]types.Commitment, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCommitment(ctx, items[i])
	}
	return items
}

func TestCommitmentGet(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNCommitment(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCommitment(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCommitmentRemove(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNCommitment(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCommitment(ctx,
			item.Index,
		)
		_, found := keeper.GetCommitment(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCommitmentGetAll(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNCommitment(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCommitment(ctx)),
	)
}
