package keeper

import (
	"context"

	"cosmossdk.io/errors"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelNFT(goCtx context.Context, msg *types.MsgCancelNFT) (*types.MsgCancelNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	swap, found := k.GetNFTSale(ctx, msg.Creator, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrSaleNotFound, "seller = %s, id = %d", msg.Creator, msg.Id)
	}

	if swap.Creator != msg.Creator {
		return nil, errors.Wrapf(types.ErrInvalidSaleData, "seller = %s", swap.Creator)
	}

	ownerAddress := k.nftKeeper.GetOwner(ctx, swap.ClassId, swap.NftId)
	if k.accountKeeper.GetModuleAddress(types.ModuleName).String() != ownerAddress.String() {
		return nil, errors.Wrapf(types.ErrInvalidSaleData, "classId = %s, nftId = %s", swap.Creator, swap.NftId)
	}

	seller, err := sdk.AccAddressFromBech32(swap.Creator)
	if err != nil {
		return nil, err
	}

	if err = k.nftKeeper.Transfer(ctx, swap.ClassId, swap.NftId, seller); err != nil {
		return nil, err
	}
	k.DeleteNFTSale(ctx, swap)
	return &types.MsgCancelNFTResponse{}, nil
}
