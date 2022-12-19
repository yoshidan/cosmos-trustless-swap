package keeper

import (
	"encoding/binary"
	"swap/x/swap/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetMaxSwapID(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MaxSwapIDKey))
	byteKey := []byte(types.MaxSwapIDKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetMaxSwapID(ctx sdk.Context, value uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MaxSwapIDKey))
	byteKey := []byte(types.MaxSwapIDKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, value)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendSwap(ctx sdk.Context, swap types.Swap) uint64 {
	maxID := k.GetMaxSwapID(ctx)
	swap.Id = maxID + 1
	k.SetSwap(ctx, swap)
	k.SetMaxSwapID(ctx, swap.Id)
	return swap.Id
}

func (k Keeper) SetSwap(ctx sdk.Context, swap types.Swap) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SwapKey))
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, swap.Id)
	updatedValue := k.cdc.MustMarshal(&swap)
	store.Set(byteKey, updatedValue)
}

func (k Keeper) GetSwap(ctx sdk.Context, id uint64) (val types.Swap, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SwapKey))

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	value := store.Get(bz)
	if value == nil {
		return types.Swap{}, false
	}
	k.cdc.MustUnmarshal(value, &val)
	return val, true
}
