package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"swap/x/sale/types"
)

func (k msgServer) Buy(goCtx context.Context, msg *types.MsgBuy) (*types.MsgBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBuyResponse{}, nil
}
