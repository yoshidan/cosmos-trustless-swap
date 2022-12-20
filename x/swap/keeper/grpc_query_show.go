package keeper

import (
	"context"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Show(goCtx context.Context, req *types.QueryShowRequest) (*types.QueryShowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	swap, found := k.GetSwap(ctx, req.Sender, req.Id)
	if !found {
		return nil, types.ErrSwapNotFound
	}
	return &types.QueryShowResponse{
		Swap: &swap,
	}, nil
}
