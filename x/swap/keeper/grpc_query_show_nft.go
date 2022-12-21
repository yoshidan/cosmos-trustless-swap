package keeper

import (
	"context"
	"cosmossdk.io/errors"

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
		return nil, errors.Wrapf(types.ErrSwapNotFound, "sender = %s, id = %d", req.Sender, req.Id)
	}
	return &types.QueryShowNFTResponse{
		Swap: &swap,
	}, nil
}
