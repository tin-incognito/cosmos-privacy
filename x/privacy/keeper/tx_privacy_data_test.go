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

func createNTxPrivacyData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TxPrivacyData {
	items := make([]types.TxPrivacyData, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTxPrivacyData(ctx, items[i])
	}
	return items
}

func TestTxPrivacyDataGet(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNTxPrivacyData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTxPrivacyData(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTxPrivacyDataRemove(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNTxPrivacyData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTxPrivacyData(ctx,
			item.Index,
		)
		_, found := keeper.GetTxPrivacyData(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTxPrivacyDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	items := createNTxPrivacyData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTxPrivacyData(ctx)),
	)
}
