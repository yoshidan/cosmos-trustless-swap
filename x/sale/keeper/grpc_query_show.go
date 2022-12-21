package keeper

import (
	"context"

	"cosmossdk.io/errors"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Show(goCtx context.Context, req *types.QueryShowRequest) (*types.QueryShowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sale, found := k.GetSale(ctx, req.Seller, req.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrSaleNotFound, "seller = %s, id = %d", req.Seller, req.Id)
	}
	return &types.QueryShowResponse{
		Sale: &sale,
	}, nil
}
