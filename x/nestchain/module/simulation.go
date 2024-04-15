package nestchain

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/r3al4f/nestchain/testutil/sample"
	nestchainsimulation "github.com/r3al4f/nestchain/x/nestchain/simulation"
	"github.com/r3al4f/nestchain/x/nestchain/types"
)

// avoid unused import issue
var (
	_ = nestchainsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateBooking = "op_weight_msg_create_booking"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBooking int = 100

	opWeightMsgUpdateLoyaltyPoints = "op_weight_msg_update_loyalty_points"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateLoyaltyPoints int = 100

	opWeightMsgProcessPayment = "op_weight_msg_process_payment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProcessPayment int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nestchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&nestchainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateBooking int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateBooking, &weightMsgCreateBooking, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBooking = defaultWeightMsgCreateBooking
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBooking,
		nestchainsimulation.SimulateMsgCreateBooking(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateLoyaltyPoints int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateLoyaltyPoints, &weightMsgUpdateLoyaltyPoints, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLoyaltyPoints = defaultWeightMsgUpdateLoyaltyPoints
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateLoyaltyPoints,
		nestchainsimulation.SimulateMsgUpdateLoyaltyPoints(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgProcessPayment int
	simState.AppParams.GetOrGenerate(opWeightMsgProcessPayment, &weightMsgProcessPayment, nil,
		func(_ *rand.Rand) {
			weightMsgProcessPayment = defaultWeightMsgProcessPayment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProcessPayment,
		nestchainsimulation.SimulateMsgProcessPayment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateBooking,
			defaultWeightMsgCreateBooking,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				nestchainsimulation.SimulateMsgCreateBooking(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateLoyaltyPoints,
			defaultWeightMsgUpdateLoyaltyPoints,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				nestchainsimulation.SimulateMsgUpdateLoyaltyPoints(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgProcessPayment,
			defaultWeightMsgProcessPayment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				nestchainsimulation.SimulateMsgProcessPayment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
