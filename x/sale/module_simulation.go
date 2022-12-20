package sale

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"swap/testutil/sample"
	salesimulation "swap/x/sale/simulation"
	"swap/x/sale/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = salesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSell = "op_weight_msg_sell"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSell int = 100

	opWeightMsgCancel = "op_weight_msg_cancel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancel int = 100

	opWeightMsgBuy = "op_weight_msg_buy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuy int = 100

	opWeightMsgSellNFT = "op_weight_msg_sell_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSellNFT int = 100

	opWeightMsgBuyNFT = "op_weight_msg_buy_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuyNFT int = 100

	opWeightMsgCancelNFT = "op_weight_msg_cancel_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelNFT int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	saleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&saleGenesis)
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

	var weightMsgSell int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSell, &weightMsgSell, nil,
		func(_ *rand.Rand) {
			weightMsgSell = defaultWeightMsgSell
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSell,
		salesimulation.SimulateMsgSell(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancel, &weightMsgCancel, nil,
		func(_ *rand.Rand) {
			weightMsgCancel = defaultWeightMsgCancel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancel,
		salesimulation.SimulateMsgCancel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuy int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuy, &weightMsgBuy, nil,
		func(_ *rand.Rand) {
			weightMsgBuy = defaultWeightMsgBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuy,
		salesimulation.SimulateMsgBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSellNFT int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSellNFT, &weightMsgSellNFT, nil,
		func(_ *rand.Rand) {
			weightMsgSellNFT = defaultWeightMsgSellNFT
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSellNFT,
		salesimulation.SimulateMsgSellNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuyNFT int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuyNFT, &weightMsgBuyNFT, nil,
		func(_ *rand.Rand) {
			weightMsgBuyNFT = defaultWeightMsgBuyNFT
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuyNFT,
		salesimulation.SimulateMsgBuyNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelNFT int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelNFT, &weightMsgCancelNFT, nil,
		func(_ *rand.Rand) {
			weightMsgCancelNFT = defaultWeightMsgCancelNFT
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelNFT,
		salesimulation.SimulateMsgCancelNFT(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
