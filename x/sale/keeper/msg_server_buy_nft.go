package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"swap/x/sale/types"
)

func (k msgServer) BuyNFT(goCtx context.Context, msg *types.MsgBuyNFT) (*types.MsgBuyNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBuyNFTResponse{}, nil
}
