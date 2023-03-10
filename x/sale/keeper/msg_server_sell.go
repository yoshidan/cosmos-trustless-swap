package keeper

import (
	"context"

	"cosmossdk.io/errors"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Sell(goCtx context.Context, msg *types.MsgSell) (*types.MsgSellResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	amount, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}

	if _, found := k.GetSale(ctx, msg.Creator, msg.Id); found {
		return nil, errors.Wrapf(types.ErrSaleExists, "seller = %s, id = %d", msg.Creator, msg.Id)
	}

	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return nil, err
	}

	k.SetSale(ctx, types.Sale{
		Id:      msg.Id,
		Creator: msg.Creator,
		Amount:  msg.Amount,
		Price:   msg.Price,
	})

	return &types.MsgSellResponse{}, nil
}
