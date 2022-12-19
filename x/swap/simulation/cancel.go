package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"swap/x/swap/keeper"
	"swap/x/swap/types"
)

func SimulateMsgCancel(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCancel{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Cancel simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Cancel simulation not implemented"), nil, nil
	}
}
