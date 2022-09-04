package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"privacy/x/privacy/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				OutputCoinSerialNumber: &types.OutputCoinSerialNumber{},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated serialNumber",
			genState: &types.GenesisState{
				SerialNumberList: []types.SerialNumber{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated outputCoin",
			genState: &types.GenesisState{
				OutputCoinList: []types.OutputCoin{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated commitment",
			genState: &types.GenesisState{
				CommitmentList: []types.Commitment{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated commitmentIndex",
			genState: &types.GenesisState{
				CommitmentIndexList: []types.CommitmentIndex{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated token",
			genState: &types.GenesisState{
				TokenList: []types.Token{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated onetimeAddress",
			genState: &types.GenesisState{
				OnetimeAddressList: []types.OnetimeAddress{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated txPrivacyData",
			genState: &types.GenesisState{
				TxPrivacyDataList: []types.TxPrivacyData{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated oTACoin",
			genState: &types.GenesisState{
				OTACoinList: []types.OTACoin{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
