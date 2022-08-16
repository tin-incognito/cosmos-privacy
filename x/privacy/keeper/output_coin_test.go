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

func createNOutputCoin(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.OutputCoin {
	items := make([]types.OutputCoin, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetOutputCoin(ctx, items[i])
	}
	return items
}

func TestOutputCoinGet(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNOutputCoin(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetOutputCoin(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestOutputCoinRemove(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNOutputCoin(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOutputCoin(ctx,
			item.Index,
		)
		_, found := keeper.GetOutputCoin(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestOutputCoinGetAll(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNOutputCoin(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOutputCoin(ctx)),
	)
}
