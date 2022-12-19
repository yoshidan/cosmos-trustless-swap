package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"swap/x/swap/types"
)

func (k msgServer) Receive(goCtx context.Context, msg *types.MsgReceive) (*types.MsgReceiveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgReceiveResponse{}, nil
}
