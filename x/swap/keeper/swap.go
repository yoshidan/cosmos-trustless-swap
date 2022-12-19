package keeper

import (
	"encoding/binary"
	"swap/x/swap/types"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type entity interface {
	codec.ProtoMarshaler
	GetId() uint64
}

func (k Keeper) getMaxID(ctx sdk.Context, key string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(key))
	byteKey := []byte(key)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) setMaxID(ctx sdk.Context, key string, value uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(key))
	byteKey := []byte(key)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, value)
	store.Set(byteKey, bz)
}

func (k Keeper) setData(ctx sdk.Context, key string, swap entity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(key))
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, swap.GetId())
	updatedValue := k.cdc.MustMarshal(swap)
	store.Set(byteKey, updatedValue)
}

func (k Keeper) deleteData(ctx sdk.Context, key string, swap entity) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(key))
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, swap.GetId())
	store.Delete(byteKey)
}

func (k Keeper) getData(ctx sdk.Context, key string, id uint64, swap entity) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(key))

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	value := store.Get(bz)
	if value == nil {
		return false
	}
	k.cdc.MustUnmarshal(value, swap)
	return true
}

func (k Keeper) GetMaxSwapID(ctx sdk.Context) uint64 {
	return k.getMaxID(ctx, types.MaxSwapIDKey)
}

func (k Keeper) SetMaxSwapID(ctx sdk.Context, value uint64) {
	k.setMaxID(ctx, types.MaxSwapIDKey, value)
}

func (k Keeper) AppendSwap(ctx sdk.Context, swap types.Swap) uint64 {
	maxID := k.GetMaxSwapID(ctx)
	swap.Id = maxID + 1
	k.SetSwap(ctx, swap)
	k.SetMaxSwapID(ctx, swap.Id)
	return swap.Id
}

func (k Keeper) SetSwap(ctx sdk.Context, swap types.Swap) {
	k.setData(ctx, types.SwapKey, &swap)
}

func (k Keeper) DeleteSwap(ctx sdk.Context, swap types.Swap) {
	k.deleteData(ctx, types.SwapKey, &swap)
}

func (k Keeper) GetSwap(ctx sdk.Context, id uint64) (val types.Swap, found bool) {
	return val, k.getData(ctx, types.SwapKey, id, &val)
}
