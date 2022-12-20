package keeper

import (
	"context"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	amount, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}

	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return nil, err
	}

	k.SetSwap(ctx, types.Swap{
		Id:              msg.Id,
		Creator:         msg.Creator,
		Receiver:        msg.Receiver,
		Amount:          msg.Amount,
		AmountToReceive: msg.AmountToReceive,
	})

	return &types.MsgSendResponse{}, nil
}
