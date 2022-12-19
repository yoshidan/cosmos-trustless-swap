package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"swap/x/swap/types"
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

	id := k.AppendSwap(ctx, types.Swap{
		Sender:          msg.Creator,
		Receiver:        msg.Receiver,
		Amount:          msg.Amount,
		AmountToReceive: msg.AmountToReceive,
		Status:          types.SwapStatus_Active,
	})

	return &types.MsgSendResponse{
		Id: id,
	}, nil
}
