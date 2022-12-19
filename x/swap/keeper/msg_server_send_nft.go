package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"swap/x/swap/types"
)

func (k msgServer) SendNFT(goCtx context.Context, msg *types.MsgSendNFT) (*types.MsgSendNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendNFTResponse{}, nil
}
