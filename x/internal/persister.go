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
	GetCreator() string
}

func SetData(ctx sdk.Context, cdc codec.BinaryCodec, storeKey storetypes.StoreKey, keyPrefix []byte, swap entity) {
	store := prefix.NewStore(ctx.KVStore(storeKey), append(keyPrefix, swap.GetCreator()...))
	idKey := make([]byte, 8)
	binary.BigEndian.PutUint64(idKey, swap.GetId())
	updatedValue := cdc.MustMarshal(swap)
	store.Set(idKey, updatedValue)
}

func DeleteData(ctx sdk.Context, storeKey storetypes.StoreKey, keyPrefix []byte, data entity) {
	store := prefix.NewStore(ctx.KVStore(storeKey), append(keyPrefix, data.GetCreator()...))
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, data.GetId())
	store.Delete(byteKey)
}

func GetData(ctx sdk.Context, cdc codec.BinaryCodec, storeKey storetypes.StoreKey, keyPrefix []byte, creator string, id uint64, swap entity) bool {
	store := prefix.NewStore(ctx.KVStore(storeKey), append(keyPrefix, creator...))

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)

	value := store.Get(bz)
	if value == nil {
		return false
	}
	cdc.MustUnmarshal(value, swap)
	return true
}
