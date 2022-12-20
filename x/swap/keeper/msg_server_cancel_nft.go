package keeper

import (
	"context"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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

	moduleAddress := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAddress == nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	ownerAddress := k.nftKeeper.GetOwner(ctx, swap.ClassId, swap.NftId)
	if moduleAddress.String() != ownerAddress.String() {
		return nil, types.ErrInsufficientPermission
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
