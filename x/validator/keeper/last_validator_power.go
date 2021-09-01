package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/validator/types"
)

// SetLastValidatorPower set a specific lastValidatorPower in the store from its index
func (k Keeper) SetLastValidatorPower(ctx sdk.Context, lastValidatorPower types.LastValidatorPower) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastValidatorPowerKey))
	b := k.cdc.MustMarshalBinaryBare(&lastValidatorPower)
	store.Set(types.KeyPrefix(lastValidatorPower.Index), b)
}

// GetLastValidatorPower returns a lastValidatorPower from its index
func (k Keeper) GetLastValidatorPower(ctx sdk.Context, index string) (val types.LastValidatorPower, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastValidatorPowerKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteLastValidatorPower removes a lastValidatorPower from the store
func (k Keeper) RemoveLastValidatorPower(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastValidatorPowerKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllLastValidatorPower returns all lastValidatorPower
func (k Keeper) GetAllLastValidatorPower(ctx sdk.Context) (list []types.LastValidatorPower) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastValidatorPowerKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LastValidatorPower
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
