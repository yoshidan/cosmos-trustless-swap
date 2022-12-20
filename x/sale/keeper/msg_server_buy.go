package keeper

import (
	"context"

	"cosmossdk.io/errors"

	"swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Buy(goCtx context.Context, msg *types.MsgBuy) (*types.MsgBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sale, found := k.GetSale(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrSaleNotFound, "id = %d", msg.Id)
	}

	seller, err := sdk.AccAddressFromBech32(sale.Seller)
	if err != nil {
		return nil, err
	}

	buyer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	amount, err := sdk.ParseCoinNormalized(sale.Amount)
	if err != nil {
		return nil, err
	}

	price, err := sdk.ParseCoinNormalized(sale.Price)
	if err != nil {
		return nil, err
	}

	if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, buyer, sdk.NewCoins(amount)); err != nil {
		return nil, err
	}
	if err = k.bankKeeper.SendCoins(ctx, buyer, seller, sdk.NewCoins(price)); err != nil {
		return nil, err
	}
	k.DeleteSale(ctx, sale)
	return &types.MsgBuyResponse{}, nil
}
