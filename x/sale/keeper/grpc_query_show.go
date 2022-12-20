package keeper

import (
	"context"

	"swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Show(goCtx context.Context, req *types.QueryShowRequest) (*types.QueryShowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sale, found := k.GetSale(ctx, req.GetId())
	if !found {
		return nil, types.ErrSaleNotFound
	}
	return &types.QueryShowResponse{
		Sale: &sale,
	}, nil
}
