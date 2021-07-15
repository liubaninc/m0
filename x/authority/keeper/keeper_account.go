package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/liubaninc/m0/x/authority/types"
)

// Check if account has assigned role.
func (k Keeper) HasRole(ctx sdk.Context, addr string, roleToCheck string) bool {
	account := k.GetAccount(ctx, addr)

	for _, role := range account.Roles {
		if role == roleToCheck {
			return true
		}
	}
	return false
}

// Get the Account record associated with an address.
func (k Keeper) GetAccount(ctx sdk.Context, addr string) (acc types.Account) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountPrefix))

	b := store.Get(types.KeyPrefix(addr))
	if b == nil {
		return acc
	}
	k.cdc.MustUnmarshalBinaryBare(b, &acc)
	return acc
}

// Set Account record for an address.
func (k Keeper) SetAccount(ctx sdk.Context, acc types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountPrefix))
	bz := k.cdc.MustMarshalBinaryBare(&acc)
	store.Set(types.KeyPrefix(acc.Address), bz)
}

// Deletes the Account from the store.
func (k Keeper) DeleteAccount(ctx sdk.Context, address string) {
	if !k.IsAccountPresent(ctx, address) {
		return
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountPrefix))
	store.Delete(types.KeyPrefix(address))
}

// Check if the Account record associated with an address is present in the store or not.
func (k Keeper) IsAccountPresent(ctx sdk.Context, address string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountPrefix))

	return store.Has(types.KeyPrefix(address))
}

// Gets the Pending Account record associated with an address.
func (k Keeper) GetPendingAccount(ctx sdk.Context, address string) (acc types.PendingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountPrefix))
	bz := store.Get(types.KeyPrefix(address))

	if bz == nil {
		return acc
	}
	k.cdc.MustUnmarshalBinaryBare(bz, &acc)

	return acc
}

// Sets Pending Account record for an address.
func (k Keeper) SetPendingAccount(ctx sdk.Context, pendAcc types.PendingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountPrefix))
	store.Set(types.KeyPrefix(pendAcc.Address), k.cdc.MustMarshalBinaryBare(&pendAcc))
}

// Deletes the Pending Account from the store.
func (k Keeper) DeletePendingAccount(ctx sdk.Context, address string) {
	if !k.IsPendingAccountPresent(ctx, address) {
		return
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountPrefix))
	store.Delete(types.KeyPrefix(address))
}

// Check if the Pending Account record associated with an address is present in the store or not.
func (k Keeper) IsPendingAccountPresent(ctx sdk.Context, address string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountPrefix))

	return store.Has(types.KeyPrefix(address))
}

// Gets the Pending Account Revocation record associated with an address.
func (k Keeper) GetPendingAccountRevocation(ctx sdk.Context, address string) (acc types.PendingAccountRevocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationPrefix))
	bz := store.Get(types.KeyPrefix(address))

	if bz == nil {
		return acc
	}

	k.cdc.MustUnmarshalBinaryBare(bz, &acc)

	return acc
}

// Sets Pending Account Revocation record for an address.
func (k Keeper) SetPendingAccountRevocation(ctx sdk.Context, revoc types.PendingAccountRevocation) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationPrefix))
	store.Set(types.KeyPrefix(revoc.Address), k.cdc.MustMarshalBinaryBare(&revoc))
}

// Deletes the Pending Account Revocation from the store.
func (k Keeper) DeletePendingAccountRevocation(ctx sdk.Context, address string) {
	if !k.IsPendingAccountRevocationPresent(ctx, address) {
		return
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationPrefix))
	store.Delete(types.KeyPrefix(address))
}

/*
	Account Number Counter
*/
func (k Keeper) GetNextAccountNumber(ctx sdk.Context) (unit64Value gogotypes.UInt64Value) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountNumberCounterKey))
	bz := store.Get(types.KeyPrefix(types.AccountNumberCounterKey))

	if bz == nil {
		unit64Value.Value = 0
	} else {
		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &unit64Value)
	}
	bz = k.cdc.MustMarshalBinaryLengthPrefixed(&unit64Value)

	store.Set(types.KeyPrefix(types.AccountNumberCounterKey), bz)
	return
}

// Check if the Pending Account Revocation record associated with an address is present in the store or not.
func (k Keeper) IsPendingAccountRevocationPresent(ctx sdk.Context, address string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationPrefix))

	return store.Has(types.KeyPrefix(address))
}

// Count account with assigned role.
func (k Keeper) CountAccountsWithRole(ctx sdk.Context, roleToCount string) int {
	res := 0

	k.IterateAccounts(ctx, func(account types.Account) (stop bool) {
		for _, role := range account.Roles {

			if role == roleToCount {
				res++

				return false
			}
		}

		return false
	})

	return res
}

// Iterate over all stored accounts.
func (k Keeper) IterateAccounts(ctx sdk.Context, process func(types.Account) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountPrefix))
	iter := sdk.KVStorePrefixIterator(store, []byte{})
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var account types.Account

		k.cdc.MustUnmarshalBinaryBare(val, &account)

		if process(account) {
			return
		}

		iter.Next()
	}
}
