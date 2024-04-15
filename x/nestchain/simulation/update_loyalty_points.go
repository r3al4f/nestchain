package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/r3al4f/nestchain/x/nestchain/keeper"
	"github.com/r3al4f/nestchain/x/nestchain/types"
)

func SimulateMsgUpdateLoyaltyPoints(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateLoyaltyPoints{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateLoyaltyPoints simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "UpdateLoyaltyPoints simulation not implemented"), nil, nil
	}
}
