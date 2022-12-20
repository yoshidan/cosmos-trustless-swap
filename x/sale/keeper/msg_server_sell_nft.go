package keeper

import (
	"context"

	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SellNFT(goCtx context.Context, msg *types.MsgSellNFT) (*types.MsgSellNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Creator)
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

	if err = k.nftKeeper.Transfer(ctx, msg.ClassId, msg.NftId, moduleAddress); err != nil {
		return nil, err
	}

	k.SetNFTSale(ctx, types.NFTSale{
		Id:      msg.Id,
		Creator: msg.Creator,
		ClassId: msg.ClassId,
		NftId:   msg.NftId,
		Price:   msg.Price,
	})

	return &types.MsgSellNFTResponse{}, nil
}
