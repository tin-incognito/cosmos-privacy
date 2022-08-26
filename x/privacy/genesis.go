package privacy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"privacy/x/privacy/keeper"
	"privacy/x/privacy/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the serialNumber
	for _, elem := range genState.SerialNumberList {
		k.SetSerialNumber(ctx, elem)
	}
	// Set all the outputCoin
	for _, elem := range genState.OutputCoinList {
		k.SetOutputCoin(ctx, elem)
	}
	// Set all the commitment
	for _, elem := range genState.CommitmentList {
		k.SetCommitment(ctx, elem)
	}
	// Set all the commitmentIndex
	for _, elem := range genState.CommitmentIndexList {
		k.SetCommitmentIndex(ctx, elem)
	}
	// Set all the token
	for _, elem := range genState.TokenList {
		k.SetToken(ctx, elem)
	}
	// Set all the onetimeAddress
	for _, elem := range genState.OnetimeAddressList {
		k.SetOnetimeAddress(ctx, elem)
	}
	// Set all the txPrivacyData
	for _, elem := range genState.TxPrivacyDataList {
		k.SetTxPrivacyData(ctx, elem)
	}
	// Set all the oTACoin
	for _, elem := range genState.OTACoinList {
		k.SetOTACoin(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.SerialNumberList = k.GetAllSerialNumber(ctx)
	genesis.OutputCoinList = k.GetAllOutputCoin(ctx)
	genesis.CommitmentList = k.GetAllCommitment(ctx)
	genesis.CommitmentIndexList = k.GetAllCommitmentIndex(ctx)
	genesis.TokenList = k.GetAllToken(ctx)
	genesis.OnetimeAddressList = k.GetAllOnetimeAddress(ctx)
	genesis.TxPrivacyDataList = k.GetAllTxPrivacyData(ctx)
	genesis.OTACoinList = k.GetAllOTACoin(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
