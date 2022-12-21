package keeper

import (
	"context"

	"cosmossdk.io/errors"
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
		return nil, errors.Wrapf(types.ErrInsufficientPermission, "classId = %s, nftId = %s", msg.ClassId, msg.NftId)
	}
	moduleAddress := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAddress == nil {
		return nil, sdkerrors.ErrInvalidAddress
	}
	if _, found := k.GetNFTSale(ctx, msg.Creator, msg.Id); found {
		return nil, errors.Wrapf(types.ErrSaleExists, "seller = %s, id = %d", msg.Creator, msg.Id)
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
