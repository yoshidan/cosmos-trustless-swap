package keeper

import (
	"context"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelNFT(goCtx context.Context, msg *types.MsgCancelNFT) (*types.MsgCancelNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	swap, found := k.GetNFTSwap(ctx, msg.Creator, msg.Id)
	if !found {
		return nil, types.ErrSwapNotFound
	}

	if swap.Creator != msg.Creator {
		return nil, types.ErrInvalidSwapData
	}

	ownerAddress := k.nftKeeper.GetOwner(ctx, swap.ClassId, swap.NftId)
	if k.accountKeeper.GetModuleAddress(types.ModuleName).String() != ownerAddress.String() {
		return nil, types.ErrInvalidSwapData
	}

	sender, err := sdk.AccAddressFromBech32(swap.Creator)
	if err != nil {
		return nil, err
	}

	if err = k.nftKeeper.Transfer(ctx, swap.ClassId, swap.NftId, sender); err != nil {
		return nil, err
	}
	k.DeleteNFTSwap(ctx, swap)
	return &types.MsgCancelNFTResponse{}, nil
}
