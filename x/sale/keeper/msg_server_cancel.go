package keeper

import (
	"context"

	"cosmossdk.io/errors"

	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Cancel(goCtx context.Context, msg *types.MsgCancel) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	swap, found := k.GetSale(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrSaleNotFound, "id = %d", msg.Id)
	}

	if swap.Seller != msg.Creator {
		return nil, types.ErrInsufficientPermission
	}

	seller, err := sdk.AccAddressFromBech32(swap.Seller)
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
