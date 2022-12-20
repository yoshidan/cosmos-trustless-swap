package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"swap/x/sale/types"
)

func (k msgServer) SellNFT(goCtx context.Context, msg *types.MsgSellNFT) (*types.MsgSellNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSellNFTResponse{}, nil
}
