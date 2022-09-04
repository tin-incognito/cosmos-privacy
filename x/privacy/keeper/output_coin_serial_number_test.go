package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "privacy/testutil/keeper"
	"privacy/testutil/nullify"
	"privacy/x/privacy/keeper"
	"privacy/x/privacy/types"
)

func createTestOutputCoinSerialNumber(keeper *keeper.Keeper, ctx sdk.Context) types.OutputCoinSerialNumber {
	item := types.OutputCoinSerialNumber{}
	keeper.SetOutputCoinSerialNumber(ctx, item)
	return item
}

func TestOutputCoinSerialNumberGet(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	item := createTestOutputCoinSerialNumber(keeper, ctx)
	rst, found := keeper.GetOutputCoinSerialNumber(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestOutputCoinSerialNumberRemove(t *testing.T) {
	keeper, ctx := keepertest.PrivacyKeeper(t)
	createTestOutputCoinSerialNumber(keeper, ctx)
	keeper.RemoveOutputCoinSerialNumber(ctx)
	_, found := keeper.GetOutputCoinSerialNumber(ctx)
	require.False(t, found)
}
