package swap

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/yoshidan/cosmos-trustless-swap/testutil/sample"
	swapsimulation "github.com/yoshidan/cosmos-trustless-swap/x/swap/simulation"
	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = swapsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSend = "op_weight_msg_send"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSend int = 100

	opWeightMsgReceive = "op_weight_msg_receive"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReceive int = 100

	opWeightMsgCancel = "op_weight_msg_cancel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancel int = 100

	opWeightMsgSendNFT = "op_weight_msg_send_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendNFT int = 100

	opWeightMsgCancelNFT = "op_weight_msg_cancel_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelNFT int = 100

	opWeightMsgReceiveNFT = "op_weight_msg_receive_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReceiveNFT int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	swapGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&swapGenesis)
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

	var weightMsgSend int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSend, &weightMsgSend, nil,
		func(_ *rand.Rand) {
			weightMsgSend = defaultWeightMsgSend
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSend,
		swapsimulation.SimulateMsgSend(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReceive int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReceive, &weightMsgReceive, nil,
		func(_ *rand.Rand) {
			weightMsgReceive = defaultWeightMsgReceive
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReceive,
		swapsimulation.SimulateMsgReceive(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancel, &weightMsgCancel, nil,
		func(_ *rand.Rand) {
			weightMsgCancel = defaultWeightMsgCancel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancel,
		swapsimulation.SimulateMsgCancel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendNFT int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendNFT, &weightMsgSendNFT, nil,
		func(_ *rand.Rand) {
			weightMsgSendNFT = defaultWeightMsgSendNFT
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendNFT,
		swapsimulation.SimulateMsgSendNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelNFT int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelNFT, &weightMsgCancelNFT, nil,
		func(_ *rand.Rand) {
			weightMsgCancelNFT = defaultWeightMsgCancelNFT
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelNFT,
		swapsimulation.SimulateMsgCancelNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReceiveNFT int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReceiveNFT, &weightMsgReceiveNFT, nil,
		func(_ *rand.Rand) {
			weightMsgReceiveNFT = defaultWeightMsgReceiveNFT
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReceiveNFT,
		swapsimulation.SimulateMsgReceiveNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
