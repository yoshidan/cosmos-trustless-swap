package keeper

import (
	"swap/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetMaxNFTSwapID(ctx sdk.Context) uint64 {
	return k.getMaxID(ctx, types.MaxNFTSwapIDKey)
}

func (k Keeper) SetMaxNFTSwapID(ctx sdk.Context, value uint64) {
	k.setMaxID(ctx, types.MaxNFTSwapIDKey, value)
}

func (k Keeper) AppendNFTSwap(ctx sdk.Context, swap types.NFTSwap) uint64 {
	maxID := k.GetMaxNFTSwapID(ctx)
	swap.Id = maxID + 1
	k.SetNFTSwap(ctx, swap)
	k.SetMaxNFTSwapID(ctx, swap.Id)
	return swap.Id
}

func (k Keeper) SetNFTSwap(ctx sdk.Context, swap types.NFTSwap) {
	k.setData(ctx, types.NFTSwapKey, &swap)
}

func (k Keeper) DeleteNFTSwap(ctx sdk.Context, swap types.NFTSwap) {
	k.deleteData(ctx, types.NFTSwapKey, &swap)
}

func (k Keeper) GetNFTSwap(ctx sdk.Context, id uint64) (val types.NFTSwap, found bool) {
	return val, k.getData(ctx, types.NFTSwapKey, id, &val)
}
