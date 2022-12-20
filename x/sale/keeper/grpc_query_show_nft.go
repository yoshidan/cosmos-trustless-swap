package keeper

import (
	"context"

	"swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowNFT(goCtx context.Context, req *types.QueryShowNFTRequest) (*types.QueryShowNFTResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sale, found := k.GetNFTSale(ctx, req.GetId())
	if !found {
		return nil, types.ErrSaleNotFound
	}
	return &types.QueryShowNFTResponse{
		Sale: &sale,
	}, nil
}
