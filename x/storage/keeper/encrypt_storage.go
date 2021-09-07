package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/storage/types"
)

// SetEncryptStorage set a specific encryptStorage in the store from its index
func (k Keeper) SetEncryptStorage(ctx sdk.Context, encryptStorage types.EncryptStorage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptStorageKey))
	b := k.cdc.MustMarshalBinaryBare(&encryptStorage)
	store.Set(types.KeyPrefix(encryptStorage.Index), b)
}

// GetEncryptStorage returns a encryptStorage from its index
func (k Keeper) GetEncryptStorage(ctx sdk.Context, index string) (val types.EncryptStorage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptStorageKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteEncryptStorage removes a encryptStorage from the store
func (k Keeper) RemoveEncryptStorage(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptStorageKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllEncryptStorage returns all encryptStorage
func (k Keeper) GetAllEncryptStorage(ctx sdk.Context) (list []types.EncryptStorage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EncryptStorageKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EncryptStorage
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
