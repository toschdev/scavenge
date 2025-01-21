package scavenge

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"scavenge/testutil/sample"
	scavengesimulation "scavenge/x/scavenge/simulation"
	"scavenge/x/scavenge/types"
)

// avoid unused import issue
var (
	_ = scavengesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateQuestion = "op_weight_msg_create_question"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateQuestion int = 100

	opWeightMsgCommitAnswer = "op_weight_msg_commit_answer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCommitAnswer int = 100

	opWeightMsgRevealAnswer = "op_weight_msg_reveal_answer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevealAnswer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	scavengeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&scavengeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateQuestion int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateQuestion, &weightMsgCreateQuestion, nil,
		func(_ *rand.Rand) {
			weightMsgCreateQuestion = defaultWeightMsgCreateQuestion
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateQuestion,
		scavengesimulation.SimulateMsgCreateQuestion(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCommitAnswer int
	simState.AppParams.GetOrGenerate(opWeightMsgCommitAnswer, &weightMsgCommitAnswer, nil,
		func(_ *rand.Rand) {
			weightMsgCommitAnswer = defaultWeightMsgCommitAnswer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCommitAnswer,
		scavengesimulation.SimulateMsgCommitAnswer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRevealAnswer int
	simState.AppParams.GetOrGenerate(opWeightMsgRevealAnswer, &weightMsgRevealAnswer, nil,
		func(_ *rand.Rand) {
			weightMsgRevealAnswer = defaultWeightMsgRevealAnswer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevealAnswer,
		scavengesimulation.SimulateMsgRevealAnswer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateQuestion,
			defaultWeightMsgCreateQuestion,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				scavengesimulation.SimulateMsgCreateQuestion(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCommitAnswer,
			defaultWeightMsgCommitAnswer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				scavengesimulation.SimulateMsgCommitAnswer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRevealAnswer,
			defaultWeightMsgRevealAnswer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				scavengesimulation.SimulateMsgRevealAnswer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
