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

func createNSerialNumber(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SerialNumber {
	items := make([]types.SerialNumber, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSerialNumber(ctx, items[i])
	}
	return items
}

func TestSerialNumberGet(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNSerialNumber(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSerialNumber(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSerialNumberRemove(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNSerialNumber(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSerialNumber(ctx,
			item.Index,
		)
		_, found := keeper.GetSerialNumber(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSerialNumberGetAll(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNSerialNumber(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSerialNumber(ctx)),
	)
}
