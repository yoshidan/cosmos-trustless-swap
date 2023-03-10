package keeper

import (
	"context"

	"cosmossdk.io/errors"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowNFT(goCtx context.Context, req *types.QueryShowNFTRequest) (*types.QueryShowNFTResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	sale, found := k.GetNFTSale(ctx, req.Seller, req.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrSaleNotFound, "seller = %s, id = %d", req.Seller, req.Id)
	}
	return &types.QueryShowNFTResponse{
		Sale: &sale,
	}, nil
}
