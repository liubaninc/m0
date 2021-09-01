package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/permission/types"
)

// SetAccount set a specific account in the store from its index
func (k Keeper) SetAccount(ctx sdk.Context, account types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	b := k.cdc.MustMarshalBinaryBare(&account)
	store.Set(types.KeyPrefix(account.Address), b)
}

// GetAccount returns a account from its index
func (k Keeper) GetAccount(ctx sdk.Context, index string) (val types.Account, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteAccount removes a account from the store
func (k Keeper) RemoveAccount(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllAccount returns all account
func (k Keeper) GetAllAccount(ctx sdk.Context) (list []types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Account
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) HasPerm(ctx sdk.Context, addr string, perm string) bool {
	if !k.GetParams(ctx).Enabled {
		return true
	}

	account, _ := k.GetAccount(ctx, addr)
	for _, role := range account.Perms {
		if role == types.AllPermissions {
			return true
		}
		if role == perm {
			return true
		}
	}
	return false
}
