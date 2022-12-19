package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"swap/x/swap/types"
)

func (k msgServer) ReceiveNFT(goCtx context.Context, msg *types.MsgReceiveNFT) (*types.MsgReceiveNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgReceiveNFTResponse{}, nil
}
