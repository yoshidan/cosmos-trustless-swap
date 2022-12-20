package keeper

import (
	"context"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowNFT(goCtx context.Context, req *types.QueryShowNFTRequest) (*types.QueryShowNFTResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	swap, found := k.GetNFTSwap(ctx, req.Sender, req.Id)
	if !found {
		return nil, types.ErrSwapNotFound
	}
	return &types.QueryShowNFTResponse{
		Swap: &swap,
	}, nil
}
