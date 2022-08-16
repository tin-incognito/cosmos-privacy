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

func createNOnetimeAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.OnetimeAddress {
	items := make([]types.OnetimeAddress, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetOnetimeAddress(ctx, items[i])
	}
	return items
}

func TestOnetimeAddressGet(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNOnetimeAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetOnetimeAddress(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestOnetimeAddressRemove(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNOnetimeAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOnetimeAddress(ctx,
			item.Index,
		)
		_, found := keeper.GetOnetimeAddress(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestOnetimeAddressGetAll(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNOnetimeAddress(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOnetimeAddress(ctx)),
	)
}
