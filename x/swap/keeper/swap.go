package keeper

import (
	"github.com/yoshidan/cosmos-trustless-swap/x/internal"
	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetMaxSwapID(ctx sdk.Context) uint64 {
	return internal.GetMaxID(ctx, k.storeKey, types.KeyPrefix(types.MaxSwapIDKey))
}

func (k Keeper) SetMaxSwapID(ctx sdk.Context, value uint64) {
	internal.SetMaxID(ctx, k.storeKey, types.KeyPrefix(types.MaxSwapIDKey), value)
}

func (k Keeper) AppendSwap(ctx sdk.Context, swap types.Swap) uint64 {
	maxID := k.GetMaxSwapID(ctx)
	swap.Id = maxID + 1
	k.SetSwap(ctx, swap)
	k.SetMaxSwapID(ctx, swap.Id)
	return swap.Id
}

func (k Keeper) SetSwap(ctx sdk.Context, swap types.Swap) {
	internal.SetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.SwapKey), &swap)
}

func (k Keeper) DeleteSwap(ctx sdk.Context, swap types.Swap) {
	internal.DeleteData(ctx, k.storeKey, types.KeyPrefix(types.SwapKey), &swap)
}

func (k Keeper) GetSwap(ctx sdk.Context, id uint64) (val types.Swap, found bool) {
	return val, internal.GetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.SwapKey), id, &val)
}

func (k Keeper) GetMaxNFTSwapID(ctx sdk.Context) uint64 {
	return internal.GetMaxID(ctx, k.storeKey, types.KeyPrefix(types.MaxNFTSwapIDKey))
}

func (k Keeper) SetMaxNFTSwapID(ctx sdk.Context, value uint64) {
	internal.SetMaxID(ctx, k.storeKey, types.KeyPrefix(types.MaxNFTSwapIDKey), value)
}

func (k Keeper) AppendNFTSwap(ctx sdk.Context, swap types.NFTSwap) uint64 {
	maxID := k.GetMaxNFTSwapID(ctx)
	swap.Id = maxID + 1
	k.SetNFTSwap(ctx, swap)
	k.SetMaxNFTSwapID(ctx, swap.Id)
	return swap.Id
}

func (k Keeper) SetNFTSwap(ctx sdk.Context, swap types.NFTSwap) {
	internal.SetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.NFTSwapKey), &swap)
}

func (k Keeper) DeleteNFTSwap(ctx sdk.Context, swap types.NFTSwap) {
	internal.DeleteData(ctx, k.storeKey, types.KeyPrefix(types.NFTSwapKey), &swap)
}

func (k Keeper) GetNFTSwap(ctx sdk.Context, id uint64) (val types.NFTSwap, found bool) {
	return val, internal.GetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.NFTSwapKey), id, &val)
}
