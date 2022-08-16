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

func createNCommitmentIndex(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CommitmentIndex {
	items := make([]types.CommitmentIndex, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetCommitmentIndex(ctx, items[i])
	}
	return items
}

func TestCommitmentIndexGet(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNCommitmentIndex(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCommitmentIndex(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCommitmentIndexRemove(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNCommitmentIndex(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCommitmentIndex(ctx,
			item.Index,
		)
		_, found := keeper.GetCommitmentIndex(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestCommitmentIndexGetAll(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNCommitmentIndex(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCommitmentIndex(ctx)),
	)
}
