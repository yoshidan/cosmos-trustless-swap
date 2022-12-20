package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"swap/x/sale/types"
)

func (k msgServer) Cancel(goCtx context.Context, msg *types.MsgCancel) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCancelResponse{}, nil
}
