package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/yoshidan/cosmos-trustless-swap/x/swap/keeper"
	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"
)

func SimulateMsgCancelNFT(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCancelNFT{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CancelNFT simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CancelNFT simulation not implemented"), nil, nil
	}
}
