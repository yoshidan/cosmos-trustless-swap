package keeper

import (
	"github.com/yoshidan/cosmos-trustless-swap/x/internal"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetMaxSaleID(ctx sdk.Context) uint64 {
	return internal.GetMaxID(ctx, k.storeKey, types.KeyPrefix(types.MaxSaleIDKey))
}

func (k Keeper) SetMaxSaleID(ctx sdk.Context, value uint64) {
	internal.SetMaxID(ctx, k.storeKey, types.KeyPrefix(types.MaxSaleIDKey), value)
}

func (k Keeper) AppendSale(ctx sdk.Context, sale types.Sale) uint64 {
	maxID := k.GetMaxSaleID(ctx)
	sale.Id = maxID + 1
	k.SetSale(ctx, sale)
	k.SetMaxSaleID(ctx, sale.Id)
	return sale.Id
}

func (k Keeper) SetSale(ctx sdk.Context, sale types.Sale) {
	internal.SetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.SaleKey), &sale)
}

func (k Keeper) DeleteSale(ctx sdk.Context, sale types.Sale) {
	internal.DeleteData(ctx, k.storeKey, types.KeyPrefix(types.SaleKey), &sale)
}

func (k Keeper) GetSale(ctx sdk.Context, id uint64) (val types.Sale, found bool) {
	return val, internal.GetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.SaleKey), id, &val)
}

func (k Keeper) GetMaxNFTSaleID(ctx sdk.Context) uint64 {
	return internal.GetMaxID(ctx, k.storeKey, types.KeyPrefix(types.MaxNFTSaleIDKey))
}

func (k Keeper) SetMaxNFTSaleID(ctx sdk.Context, value uint64) {
	internal.SetMaxID(ctx, k.storeKey, types.KeyPrefix(types.MaxNFTSaleIDKey), value)
}

func (k Keeper) AppendNFTSale(ctx sdk.Context, sale types.NFTSale) uint64 {
	maxID := k.GetMaxNFTSaleID(ctx)
	sale.Id = maxID + 1
	k.SetNFTSale(ctx, sale)
	k.SetMaxNFTSaleID(ctx, sale.Id)
	return sale.Id
}

func (k Keeper) SetNFTSale(ctx sdk.Context, sale types.NFTSale) {
	internal.SetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.NFTSaleKey), &sale)
}

func (k Keeper) DeleteNFTSale(ctx sdk.Context, sale types.NFTSale) {
	internal.DeleteData(ctx, k.storeKey, types.KeyPrefix(types.NFTSaleKey), &sale)
}

func (k Keeper) GetNFTSale(ctx sdk.Context, id uint64) (val types.NFTSale, found bool) {
	return val, internal.GetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.NFTSaleKey), id, &val)
}
