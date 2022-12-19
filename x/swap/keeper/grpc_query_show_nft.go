package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"swap/x/swap/types"
)

func (k Keeper) ShowNFT(goCtx context.Context, req *types.QueryShowNFTRequest) (*types.QueryShowNFTResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryShowNFTResponse{}, nil
}
