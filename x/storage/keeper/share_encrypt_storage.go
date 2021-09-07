package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/storage/types"
)

// SetShareEncryptStorage set a specific shareEncryptStorage in the store from its index
func (k Keeper) SetShareEncryptStorage(ctx sdk.Context, shareEncryptStorage types.ShareEncryptStorage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShareEncryptStorageKey))
	b := k.cdc.MustMarshalBinaryBare(&shareEncryptStorage)
	store.Set(types.KeyPrefix(shareEncryptStorage.Index), b)
}

// GetShareEncryptStorage returns a shareEncryptStorage from its index
func (k Keeper) GetShareEncryptStorage(ctx sdk.Context, index string) (val types.ShareEncryptStorage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShareEncryptStorageKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteShareEncryptStorage removes a shareEncryptStorage from the store
func (k Keeper) RemoveShareEncryptStorage(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShareEncryptStorageKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllShareEncryptStorage returns all shareEncryptStorage
func (k Keeper) GetAllShareEncryptStorage(ctx sdk.Context) (list []types.ShareEncryptStorage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShareEncryptStorageKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ShareEncryptStorage
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
