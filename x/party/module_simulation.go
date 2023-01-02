package party

import (
	"math/rand"

	"github.com/TeaPartyCrypto/partychain/testutil/sample"
	partysimulation "github.com/TeaPartyCrypto/partychain/x/party/simulation"
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = partysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSubmitSell = "op_weight_msg_submit_sell"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitSell int = 100

	opWeightMsgBuy = "op_weight_msg_buy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuy int = 100

	opWeightMsgCancel = "op_weight_msg_cancel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancel int = 100

	opWeightMsgAccountWatchOutcome = "op_weight_msg_account_watch_outcome"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAccountWatchOutcome int = 100

	opWeightMsgAccountWatchFailure = "op_weight_msg_account_watch_failure"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAccountWatchFailure int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	partyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&partyGenesis)
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

	var weightMsgSubmitSell int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitSell, &weightMsgSubmitSell, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitSell = defaultWeightMsgSubmitSell
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitSell,
		partysimulation.SimulateMsgSubmitSell(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuy int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuy, &weightMsgBuy, nil,
		func(_ *rand.Rand) {
			weightMsgBuy = defaultWeightMsgBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuy,
		partysimulation.SimulateMsgBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancel, &weightMsgCancel, nil,
		func(_ *rand.Rand) {
			weightMsgCancel = defaultWeightMsgCancel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancel,
		partysimulation.SimulateMsgCancel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAccountWatchOutcome int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAccountWatchOutcome, &weightMsgAccountWatchOutcome, nil,
		func(_ *rand.Rand) {
			weightMsgAccountWatchOutcome = defaultWeightMsgAccountWatchOutcome
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAccountWatchOutcome,
		partysimulation.SimulateMsgAccountWatchOutcome(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAccountWatchFailure int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAccountWatchFailure, &weightMsgAccountWatchFailure, nil,
		func(_ *rand.Rand) {
			weightMsgAccountWatchFailure = defaultWeightMsgAccountWatchFailure
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAccountWatchFailure,
		partysimulation.SimulateMsgAccountWatchFailure(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
