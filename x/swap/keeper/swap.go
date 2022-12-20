package keeper

import (
	"github.com/yoshidan/cosmos-trustless-swap/x/internal"
	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetSwap(ctx sdk.Context, swap types.Swap) {
	internal.SetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.SwapKey), &swap)
}

func (k Keeper) DeleteSwap(ctx sdk.Context, swap types.Swap) {
	internal.DeleteData(ctx, k.storeKey, types.KeyPrefix(types.SwapKey), &swap)
}

func (k Keeper) GetSwap(ctx sdk.Context, creator string, id uint64) (val types.Swap, found bool) {
	return val, internal.GetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.SwapKey), creator, id, &val)
}

func (k Keeper) SetNFTSwap(ctx sdk.Context, swap types.NFTSwap) {
	internal.SetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.NFTSwapKey), &swap)
}

func (k Keeper) DeleteNFTSwap(ctx sdk.Context, swap types.NFTSwap) {
	internal.DeleteData(ctx, k.storeKey, types.KeyPrefix(types.NFTSwapKey), &swap)
}

func (k Keeper) GetNFTSwap(ctx sdk.Context, creator string, id uint64) (val types.NFTSwap, found bool) {
	return val, internal.GetData(ctx, k.cdc, k.storeKey, types.KeyPrefix(types.NFTSwapKey), creator, id, &val)
}
