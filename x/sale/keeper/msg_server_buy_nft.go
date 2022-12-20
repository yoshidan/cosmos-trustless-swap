package keeper

import (
	"context"

	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyNFT(goCtx context.Context, msg *types.MsgBuyNFT) (*types.MsgBuyNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sale, found := k.GetNFTSale(ctx, msg.Seller, msg.Id)
	if !found {
		return nil, types.ErrSaleNotFound
	}

	moduleAddress := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if moduleAddress == nil {
		return nil, sdkerrors.ErrInvalidAddress
	}

	ownerAddress := k.nftKeeper.GetOwner(ctx, sale.ClassId, sale.NftId)
	if moduleAddress.String() != ownerAddress.String() {
		return nil, types.ErrInsufficientPermission
	}

	seller, err := sdk.AccAddressFromBech32(sale.Creator)
	if err != nil {
		return nil, err
	}

	buyer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	price, err := sdk.ParseCoinNormalized(sale.Price)
	if err != nil {
		return nil, err
	}

	if err = k.nftKeeper.Transfer(ctx, sale.ClassId, sale.NftId, buyer); err != nil {
		return nil, err
	}
	if err = k.bankKeeper.SendCoins(ctx, buyer, seller, sdk.NewCoins(price)); err != nil {
		return nil, err
	}
	k.DeleteNFTSale(ctx, sale)

	return &types.MsgBuyNFTResponse{}, nil
}