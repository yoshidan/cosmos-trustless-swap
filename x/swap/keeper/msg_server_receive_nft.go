package keeper

import (
	"context"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ReceiveNFT(goCtx context.Context, msg *types.MsgReceiveNFT) (*types.MsgReceiveNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	swap, found := k.GetNFTSwap(ctx, msg.Sender, msg.Id)
	if !found {
		return nil, types.ErrSwapNotFound
	}

	if swap.Receiver != msg.Creator {
		return nil, types.ErrInsufficientPermission
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

	receiver, err := sdk.AccAddressFromBech32(swap.Receiver)
	if err != nil {
		return nil, err
	}
	amountToReceive, err := sdk.ParseCoinNormalized(swap.AmountToReceive)
	if err != nil {
		return nil, err
	}

	if err = k.nftKeeper.Transfer(ctx, swap.ClassId, swap.NftId, receiver); err != nil {
		return nil, err
	}
	if err = k.bankKeeper.SendCoins(ctx, receiver, sender, sdk.NewCoins(amountToReceive)); err != nil {
		return nil, err
	}
	k.DeleteNFTSwap(ctx, swap)
	return &types.MsgReceiveNFTResponse{}, nil
}
