package keeper

import (
	"context"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	ownerAddress := k.nftKeeper.GetOwner(ctx, swap.ClassId, swap.NftId)
	if k.accountKeeper.GetModuleAddress(types.ModuleName).String() != ownerAddress.String() {
		return nil, types.ErrInvalidSwapData
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
