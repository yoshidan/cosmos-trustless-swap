package keeper

import (
	"context"

	"cosmossdk.io/errors"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Cancel(goCtx context.Context, msg *types.MsgCancel) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	swap, found := k.GetSale(ctx, msg.Creator, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrSaleNotFound, "seller = %s, id = %d", msg.Creator, msg.Id)
	}

	if swap.Creator != msg.Creator {
		return nil, errors.Wrapf(types.ErrInvalidSaleData, "seller = %s", swap.Creator)
	}

	seller, err := sdk.AccAddressFromBech32(swap.Creator)
	if err != nil {
		return nil, err
	}

	amount, err := sdk.ParseCoinNormalized(swap.Amount)
	if err != nil {
		return nil, err
	}

	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, seller, sdk.NewCoins(amount)); err != nil {
		return nil, err
	}
	k.DeleteSale(ctx, swap)
	return &types.MsgCancelResponse{}, nil
}
