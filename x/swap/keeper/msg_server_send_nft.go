package keeper

import (
	"context"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendNFT(goCtx context.Context, msg *types.MsgSendNFT) (*types.MsgSendNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}
	ownerAddress := k.nftKeeper.GetOwner(ctx, msg.ClassId, msg.NftId)
	if sender.String() != ownerAddress.String() {
		return nil, types.ErrInsufficientPermission
	}
	moduleAddress := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAddress == nil {
		return nil, sdkerrors.ErrInvalidAddress
	}
	if _, found := k.GetSwap(ctx, msg.Creator, msg.Id); found {
		return nil, types.ErrSwapExists
	}
	if err = k.nftKeeper.Transfer(ctx, msg.ClassId, msg.NftId, moduleAddress); err != nil {
		return nil, err
	}

	k.SetNFTSwap(ctx, types.NFTSwap{
		Id:              msg.Id,
		Creator:         msg.Creator,
		Receiver:        msg.Receiver,
		ClassId:         msg.ClassId,
		NftId:           msg.NftId,
		AmountToReceive: msg.AmountToReceive,
	})

	return &types.MsgSendNFTResponse{}, nil
}
