package internal

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type entity interface {
	codec.ProtoMarshaler
	GetId() uint64
}

func GetMaxID(ctx sdk.Context, storeKey storetypes.StoreKey, keyPrefix []byte) uint64 {
	store := prefix.NewStore(ctx.KVStore(storeKey), keyPrefix)
	bz := store.Get(keyPrefix)

	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func SetMaxID(ctx sdk.Context, storeKey storetypes.StoreKey, keyPrefix []byte, value uint64) {
	store := prefix.NewStore(ctx.KVStore(storeKey), keyPrefix)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, value)
	store.Set(keyPrefix, bz)
}

func SetData(ctx sdk.Context, cdc codec.BinaryCodec, storeKey storetypes.StoreKey, keyPrefix []byte, swap entity) {
	store := prefix.NewStore(ctx.KVStore(storeKey), keyPrefix)
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, swap.GetId())
	updatedValue := cdc.MustMarshal(swap)
	store.Set(byteKey, updatedValue)
}

func DeleteData(ctx sdk.Context, storeKey storetypes.StoreKey, keyPrefix []byte, data entity) {
	store := prefix.NewStore(ctx.KVStore(storeKey), keyPrefix)
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, data.GetId())
	store.Delete(byteKey)
}

func GetData(ctx sdk.Context, cdc codec.BinaryCodec, storeKey storetypes.StoreKey, keyPrefix []byte, id uint64, swap entity) bool {
	store := prefix.NewStore(ctx.KVStore(storeKey), keyPrefix)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	value := store.Get(bz)
	if value == nil {
		return false
	}
	cdc.MustUnmarshal(value, swap)
	return true
}
