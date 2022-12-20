package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"swap/x/sale/types"
)

func (k msgServer) CancelNFT(goCtx context.Context, msg *types.MsgCancelNFT) (*types.MsgCancelNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCancelNFTResponse{}, nil
}
