package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/keeper"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"
)

func SimulateMsgBuyNFT(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBuyNFT{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the BuyNFT simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "BuyNFT simulation not implemented"), nil, nil
	}
}
