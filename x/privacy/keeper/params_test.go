package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "privacy/testutil/keeper"
	"privacy/x/privacy/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PrivacyKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
