package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/storage/types"
)

// SetRecryptAccount set a specific recryptAccount in the store from its index
func (k Keeper) SetRecryptAccount(ctx sdk.Context, recryptAccount types.RecryptAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecryptAccountKey))
	b := k.cdc.MustMarshalBinaryBare(&recryptAccount)
	store.Set(types.KeyPrefix(recryptAccount.Index), b)
}

// GetRecryptAccount returns a recryptAccount from its index
func (k Keeper) GetRecryptAccount(ctx sdk.Context, index string) (val types.RecryptAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecryptAccountKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteRecryptAccount removes a recryptAccount from the store
func (k Keeper) RemoveRecryptAccount(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecryptAccountKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllRecryptAccount returns all recryptAccount
func (k Keeper) GetAllRecryptAccount(ctx sdk.Context) (list []types.RecryptAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecryptAccountKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RecryptAccount
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
