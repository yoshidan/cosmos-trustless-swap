package keeper

import (
	"context"

	"swap/x/swap/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Receive(goCtx context.Context, msg *types.MsgReceive) (*types.MsgReceiveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	swap, found := k.GetSwap(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrSwapNotFound, "id = %d", msg.Id)
	}

	if swap.Receiver != msg.Creator {
		return nil, types.ErrInsufficientPermission
	}

	if swap.Status != types.SwapStatus_Active {
		return nil, errors.Wrapf(types.ErrInvalidSwapStatus, "actual = %d, expected = %d", swap.Status, types.SwapStatus_Active)
	}

	sender, err := sdk.AccAddressFromBech32(swap.Sender)
	if err != nil {
		return nil, err
	}

	receiver, err := sdk.AccAddressFromBech32(swap.Sender)
	if err != nil {
		return nil, err
	}

	amount, err := sdk.ParseCoinNormalized(swap.Amount)
	if err != nil {
		return nil, err
	}

	amountToReceive, err := sdk.ParseCoinNormalized(swap.AmountToReceive)
	if err != nil {
		return nil, err
	}

	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, sdk.NewCoins(amount)); err != nil {
		return nil, err
	}
	if err = k.bankKeeper.SendCoins(ctx, receiver, sender, sdk.NewCoins(amountToReceive)); err != nil {
		return nil, err
	}
	swap.Status = types.SwapStatus_Closed
	k.SetSwap(ctx, swap)
	return &types.MsgReceiveResponse{}, nil
}
