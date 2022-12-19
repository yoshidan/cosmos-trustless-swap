package keeper

import (
	"context"

	"cosmossdk.io/errors"

	"swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Cancel(goCtx context.Context, msg *types.MsgCancel) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	swap, found := k.GetSwap(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrSwapNotFound, "id = %d", msg.Id)
	}

	if swap.Sender != msg.Creator {
		return nil, types.ErrInsufficientPermission
	}

	sender, err := sdk.AccAddressFromBech32(swap.Sender)
	if err != nil {
		return nil, err
	}

	amount, err := sdk.ParseCoinNormalized(swap.Amount)
	if err != nil {
		return nil, err
	}

	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(amount)); err != nil {
		return nil, err
	}
	k.DeleteSwap(ctx, swap)
	return &types.MsgCancelResponse{}, nil
}
