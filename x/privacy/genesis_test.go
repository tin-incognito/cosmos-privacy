package privacy_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "privacy/testutil/keeper"
	"privacy/testutil/nullify"
	"privacy/x/privacy"
	"privacy/x/privacy/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SerialNumberList: []types.SerialNumber{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		OutputCoinList: []types.OutputCoin{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CommitmentList: []types.Commitment{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CommitmentIndexList: []types.CommitmentIndex{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		TokenList: []types.Token{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		OnetimeAddressList: []types.OnetimeAddress{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		TxPrivacyDataList: []types.TxPrivacyData{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		OTACoinList: []types.OTACoin{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PrivacyKeeper(t)
	privacy.InitGenesis(ctx, *k, genesisState)
	got := privacy.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SerialNumberList, got.SerialNumberList)
	require.ElementsMatch(t, genesisState.OutputCoinList, got.OutputCoinList)
	require.ElementsMatch(t, genesisState.CommitmentList, got.CommitmentList)
	require.ElementsMatch(t, genesisState.CommitmentIndexList, got.CommitmentIndexList)
	require.ElementsMatch(t, genesisState.TokenList, got.TokenList)
	require.ElementsMatch(t, genesisState.OnetimeAddressList, got.OnetimeAddressList)
	require.ElementsMatch(t, genesisState.TxPrivacyDataList, got.TxPrivacyDataList)
	require.ElementsMatch(t, genesisState.OTACoinList, got.OTACoinList)
	// this line is used by starport scaffolding # genesis/test/assert
}
