package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"swap/x/swap/types"
)

func (k Keeper) GetSwapCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SwapCountKey))
	byteKey := []byte(types.SwapCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetSwapCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SwapCountKey))
	byteKey := []byte(types.SwapCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendSwap(ctx sdk.Context, swap types.Swap) uint64 {
	count := k.GetSwapCount(ctx)
	swap.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SwapKey))
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, swap.Id)
	appendedValue := k.cdc.MustMarshal(&swap)
	store.Set(byteKey, appendedValue)
	k.SetSwapCount(ctx, count+1)
	return swap.Id
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
