package keeper

import (
	"context"

	"swap/x/sale/types"

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

	if err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return nil, err
	}

	id := k.AppendSale(ctx, types.Sale{
		Seller: msg.Creator,
		Amount: msg.Amount,
		Price:  msg.Price,
	})

	return &types.MsgSellResponse{
		Id: id,
	}, nil
}
