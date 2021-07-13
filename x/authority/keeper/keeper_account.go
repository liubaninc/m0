package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/liubaninc/m0/x/authority/types"
)

// Check if account has assigned role.
func (k Keeper) HasRole(ctx sdk.Context, addr sdk.AccAddress, roleToCheck string) bool {
	account := k.GetAccount(ctx, addr)

	for _, role := range account.Roles {
		if role == roleToCheck {
			return true
		}
	}
	return false
}

// Get the Account record associated with an address.
func (k Keeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) (acc types.Account) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey(addr))

	b := store.Get(types.GetAccountKey(addr))
	if b == nil {
		return acc
	}
	k.cdc.MustUnmarshalBinaryBare(b, &acc)
	return acc
}

// Set Account record for an address.
func (k Keeper) SetAccount(ctx sdk.Context, acc types.Account) {
	address, _ := sdk.AccAddressFromBech32(acc.Address)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetAccountKey(address))
	bz := k.cdc.MustMarshalBinaryBare(&acc)
	store.Set(types.GetAccountKey(address), bz)
}

// Sets Pending Account record for an address.
func (k Keeper) SetPendingAccount(ctx sdk.Context, pendAcc types.PendingAccount) {
	address, _ := sdk.AccAddressFromBech32(pendAcc.Address)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetPendingAccountKey(address))
	store.Set(types.GetPendingAccountKey(address), k.cdc.MustMarshalBinaryBare(&pendAcc))
}

/*
	Account Number Counter
*/
func (k Keeper) GetNextAccountNumber(ctx sdk.Context) (unit64Value gogotypes.UInt64Value) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccountNumberCounterKey)
	bz := store.Get(types.AccountNumberCounterKey)
	var accNumber uint64
	if bz == nil {
		accNumber = 0
	} else {
		k.cdc.MustUnmarshalBinaryBare(bz, &gogotypes.UInt64Value{})
	}
	bz = k.cdc.MustMarshalBinaryBare(&gogotypes.UInt64Value{Value: accNumber + 1})

	store.Set(types.AccountNumberCounterKey, bz)
	return
}

// Check if the Account record associated with an address is present in the store or not.
func (k Keeper) IsAccountPresent(ctx sdk.Context, acc sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetAccountKey(acc))
}

// Check if the Pending Account record associated with an address is present in the store or not.
func (k Keeper) IsPendingAccountPresent(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has(types.GetPendingAccountKey(address))
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
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.AccountPrefix)

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
