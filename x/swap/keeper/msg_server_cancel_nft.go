package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"

	"swap/x/swap/types"

	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelNFT(goCtx context.Context, msg *types.MsgCancelNFT) (*types.MsgCancelNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	swap, found := k.GetNFTSwap(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrSwapNotFound, "id = %d", msg.Id)
	}

	if swap.Sender != msg.Creator {
		return nil, types.ErrInsufficientPermission
	}

	moduleAddress := k.accountKeeper.GetModuleAddress(nft.ModuleName)
	if moduleAddress == nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	ownerAddress := k.nftKeeper.GetOwner(ctx, swap.ClassId, swap.NftId)
	if moduleAddress.String() != ownerAddress.String() {
		return nil, types.ErrInsufficientPermission
	}

	sender, err := sdk.AccAddressFromBech32(swap.Sender)
	if err != nil {
		return nil, err
	}

	if err = k.nftKeeper.Transfer(ctx, swap.ClassId, swap.NftId, sender); err != nil {
		return nil, err
	}
	k.DeleteNFTSwap(ctx, swap)
	return &types.MsgCancelNFTResponse{}, nil
}
