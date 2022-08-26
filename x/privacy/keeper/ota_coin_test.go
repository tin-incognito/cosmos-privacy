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

func createNOTACoin(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.OTACoin {
	items := make([]types.OTACoin, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetOTACoin(ctx, items[i])
	}
	return items
}

func TestOTACoinGet(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNOTACoin(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetOTACoin(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestOTACoinRemove(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNOTACoin(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOTACoin(ctx,
			item.Index,
		)
		_, found := keeper.GetOTACoin(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestOTACoinGetAll(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNOTACoin(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOTACoin(ctx)),
	)
}
