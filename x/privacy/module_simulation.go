package privacy

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"privacy/testutil/sample"
	privacysimulation "privacy/x/privacy/simulation"
	"privacy/x/privacy/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = privacysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateSerialNumber = "op_weight_msg_serial_number"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSerialNumber int = 100

	opWeightMsgUpdateSerialNumber = "op_weight_msg_serial_number"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSerialNumber int = 100

	opWeightMsgDeleteSerialNumber = "op_weight_msg_serial_number"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSerialNumber int = 100

	opWeightMsgCreateOutputCoin = "op_weight_msg_output_coin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateOutputCoin int = 100

	opWeightMsgUpdateOutputCoin = "op_weight_msg_output_coin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOutputCoin int = 100

	opWeightMsgDeleteOutputCoin = "op_weight_msg_output_coin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteOutputCoin int = 100

	opWeightMsgCreateCommitment = "op_weight_msg_commitment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCommitment int = 100

	opWeightMsgUpdateCommitment = "op_weight_msg_commitment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCommitment int = 100

	opWeightMsgDeleteCommitment = "op_weight_msg_commitment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCommitment int = 100

	opWeightMsgCreateCommitmentIndex = "op_weight_msg_commitment_index"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCommitmentIndex int = 100

	opWeightMsgUpdateCommitmentIndex = "op_weight_msg_commitment_index"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCommitmentIndex int = 100

	opWeightMsgDeleteCommitmentIndex = "op_weight_msg_commitment_index"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCommitmentIndex int = 100

	opWeightMsgCreateToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateToken int = 100

	opWeightMsgUpdateToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateToken int = 100

	opWeightMsgDeleteToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteToken int = 100

	opWeightMsgCreateOnetimeAddress = "op_weight_msg_onetime_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateOnetimeAddress int = 100

	opWeightMsgUpdateOnetimeAddress = "op_weight_msg_onetime_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOnetimeAddress int = 100

	opWeightMsgDeleteOnetimeAddress = "op_weight_msg_onetime_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteOnetimeAddress int = 100

	opWeightMsgCreateTx = "op_weight_msg_create_tx"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTx int = 100

	opWeightMsgCreateTxPrivacyData = "op_weight_msg_tx_privacy_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTxPrivacyData int = 100

	opWeightMsgUpdateTxPrivacyData = "op_weight_msg_tx_privacy_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTxPrivacyData int = 100

	opWeightMsgDeleteTxPrivacyData = "op_weight_msg_tx_privacy_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTxPrivacyData int = 100

	opWeightMsgAirdrop = "op_weight_msg_airdrop"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAirdrop int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	privacyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		SerialNumberList: []types.SerialNumber{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		OutputCoinList: []types.OutputCoin{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		CommitmentList: []types.Commitment{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		CommitmentIndexList: []types.CommitmentIndex{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		TokenList: []types.Token{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		OnetimeAddressList: []types.OnetimeAddress{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		TxPrivacyDataList: []types.TxPrivacyData{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&privacyGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateSerialNumber int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSerialNumber, &weightMsgCreateSerialNumber, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSerialNumber = defaultWeightMsgCreateSerialNumber
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSerialNumber,
		privacysimulation.SimulateMsgCreateSerialNumber(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSerialNumber int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSerialNumber, &weightMsgUpdateSerialNumber, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSerialNumber = defaultWeightMsgUpdateSerialNumber
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSerialNumber,
		privacysimulation.SimulateMsgUpdateSerialNumber(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteSerialNumber int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteSerialNumber, &weightMsgDeleteSerialNumber, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSerialNumber = defaultWeightMsgDeleteSerialNumber
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteSerialNumber,
		privacysimulation.SimulateMsgDeleteSerialNumber(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateOutputCoin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateOutputCoin, &weightMsgCreateOutputCoin, nil,
		func(_ *rand.Rand) {
			weightMsgCreateOutputCoin = defaultWeightMsgCreateOutputCoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateOutputCoin,
		privacysimulation.SimulateMsgCreateOutputCoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateOutputCoin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateOutputCoin, &weightMsgUpdateOutputCoin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOutputCoin = defaultWeightMsgUpdateOutputCoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateOutputCoin,
		privacysimulation.SimulateMsgUpdateOutputCoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteOutputCoin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteOutputCoin, &weightMsgDeleteOutputCoin, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteOutputCoin = defaultWeightMsgDeleteOutputCoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteOutputCoin,
		privacysimulation.SimulateMsgDeleteOutputCoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateCommitment int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCommitment, &weightMsgCreateCommitment, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCommitment = defaultWeightMsgCreateCommitment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCommitment,
		privacysimulation.SimulateMsgCreateCommitment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCommitment int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCommitment, &weightMsgUpdateCommitment, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCommitment = defaultWeightMsgUpdateCommitment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCommitment,
		privacysimulation.SimulateMsgUpdateCommitment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCommitment int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteCommitment, &weightMsgDeleteCommitment, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCommitment = defaultWeightMsgDeleteCommitment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCommitment,
		privacysimulation.SimulateMsgDeleteCommitment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateCommitmentIndex int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCommitmentIndex, &weightMsgCreateCommitmentIndex, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCommitmentIndex = defaultWeightMsgCreateCommitmentIndex
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCommitmentIndex,
		privacysimulation.SimulateMsgCreateCommitmentIndex(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCommitmentIndex int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateCommitmentIndex, &weightMsgUpdateCommitmentIndex, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCommitmentIndex = defaultWeightMsgUpdateCommitmentIndex
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCommitmentIndex,
		privacysimulation.SimulateMsgUpdateCommitmentIndex(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCommitmentIndex int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteCommitmentIndex, &weightMsgDeleteCommitmentIndex, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCommitmentIndex = defaultWeightMsgDeleteCommitmentIndex
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCommitmentIndex,
		privacysimulation.SimulateMsgDeleteCommitmentIndex(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateToken, &weightMsgCreateToken, nil,
		func(_ *rand.Rand) {
			weightMsgCreateToken = defaultWeightMsgCreateToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateToken,
		privacysimulation.SimulateMsgCreateToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateToken, &weightMsgUpdateToken, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateToken = defaultWeightMsgUpdateToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateToken,
		privacysimulation.SimulateMsgUpdateToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteToken, &weightMsgDeleteToken, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteToken = defaultWeightMsgDeleteToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteToken,
		privacysimulation.SimulateMsgDeleteToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateOnetimeAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateOnetimeAddress, &weightMsgCreateOnetimeAddress, nil,
		func(_ *rand.Rand) {
			weightMsgCreateOnetimeAddress = defaultWeightMsgCreateOnetimeAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateOnetimeAddress,
		privacysimulation.SimulateMsgCreateOnetimeAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateOnetimeAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateOnetimeAddress, &weightMsgUpdateOnetimeAddress, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOnetimeAddress = defaultWeightMsgUpdateOnetimeAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateOnetimeAddress,
		privacysimulation.SimulateMsgUpdateOnetimeAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteOnetimeAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteOnetimeAddress, &weightMsgDeleteOnetimeAddress, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteOnetimeAddress = defaultWeightMsgDeleteOnetimeAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteOnetimeAddress,
		privacysimulation.SimulateMsgDeleteOnetimeAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTx int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTx, &weightMsgCreateTx, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTx = defaultWeightMsgCreateTx
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTx,
		privacysimulation.SimulateMsgCreateTx(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTxPrivacyData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTxPrivacyData, &weightMsgCreateTxPrivacyData, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTxPrivacyData = defaultWeightMsgCreateTxPrivacyData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTxPrivacyData,
		privacysimulation.SimulateMsgCreateTxPrivacyData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTxPrivacyData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTxPrivacyData, &weightMsgUpdateTxPrivacyData, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTxPrivacyData = defaultWeightMsgUpdateTxPrivacyData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTxPrivacyData,
		privacysimulation.SimulateMsgUpdateTxPrivacyData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTxPrivacyData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTxPrivacyData, &weightMsgDeleteTxPrivacyData, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTxPrivacyData = defaultWeightMsgDeleteTxPrivacyData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTxPrivacyData,
		privacysimulation.SimulateMsgDeleteTxPrivacyData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAirdrop int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAirdrop, &weightMsgAirdrop, nil,
		func(_ *rand.Rand) {
			weightMsgAirdrop = defaultWeightMsgAirdrop
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAirdrop,
		privacysimulation.SimulateMsgAirdrop(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
