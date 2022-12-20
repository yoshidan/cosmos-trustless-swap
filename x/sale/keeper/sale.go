package keeper

import (
	"github.com/yoshidan/cosmos-trustless-swap/x/internal"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetSale(ctx sdk.Context, sale types.Sale) {
	internal.SetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.SaleKey), &sale)
}

func (k Keeper) DeleteSale(ctx sdk.Context, sale types.Sale) {
	internal.DeleteData(ctx, k.storeKey, types.KeyPrefix(types.SaleKey), &sale)
}

func (k Keeper) GetSale(ctx sdk.Context, creator string, id uint64) (val types.Sale, found bool) {
	return val, internal.GetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.SaleKey), creator, id, &val)
}

func (k Keeper) SetNFTSale(ctx sdk.Context, sale types.NFTSale) {
	internal.SetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.NFTSaleKey), &sale)
}

func (k Keeper) DeleteNFTSale(ctx sdk.Context, sale types.NFTSale) {
	internal.DeleteData(ctx, k.storeKey, types.KeyPrefix(types.NFTSaleKey), &sale)
}

func (k Keeper) GetNFTSale(ctx sdk.Context, creator string, id uint64) (val types.NFTSale, found bool) {
	return val, internal.GetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.NFTSaleKey), creator, id, &val)
}
