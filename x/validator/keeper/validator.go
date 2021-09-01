package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/validator/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// SetValidator set a specific validator in the store
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	b := k.cdc.MustMarshalBinaryBare(&validator)
	store.Set([]byte(validator.Creator), b)
}

// GetValidator returns a validator from its id
func (k Keeper) GetValidator(ctx sdk.Context, address string) types.Validator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	var validator types.Validator
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(address)), &validator)
	return validator
}

// HasValidator checks if the validator exists in the store
func (k Keeper) HasValidator(ctx sdk.Context, address string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	return store.Has([]byte(address))
}

// HasValidator checks if the validator exists in the store
func (k Keeper) HasValidatorByConsAddr(ctx sdk.Context, address string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorConsAddrKey))
	return store.Has([]byte(address))
}

// Get a validator by consensus address
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, address string) (validator types.Validator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorConsAddrKey))

	opAddr := store.Get([]byte(address))
	if opAddr == nil {
		return validator
	}

	return k.GetValidator(ctx, string(opAddr))
}

// Set validator consensus address
func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, validator types.Validator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorConsAddrKey))

	store.Set([]byte(validator.GetConsAddr().String()), []byte(validator.Creator))
}

// RemoveValidator removes a validator from the store
func (k Keeper) RemoveValidatorByConsAddr(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorConsAddrKey))
	store.Delete([]byte(address))
}

// GetValidatorOwner returns the creator of the
func (k Keeper) GetValidatorOwner(ctx sdk.Context, address string) string {
	return k.GetValidator(ctx, address).Creator
}

// RemoveValidator removes a validator from the store
func (k Keeper) RemoveValidator(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	store.Delete([]byte(address))
}

// GetAllValidator returns all validator
func (k Keeper) GetAllValidator(ctx sdk.Context) (list []types.Validator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Validator
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ApplyAndReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate, err error) {
	// Iterate over validators.
	validators := k.GetAllValidator(ctx)
	for _, validator := range validators {
		// power on the last height.
		lastValidatorPower, _ := k.GetLastValidatorPower(ctx, validator.Creator)

		// if last power was more then 0 and potential power 0 it
		// means that validator was jailed or removed within the block.
		if lastValidatorPower.Power > 0 && validator.GetPower() == 0 {
			updates = append(updates, validator.ABCIValidatorUpdateZero())

			// set validator power on lookup index.
			k.RemoveLastValidatorPower(ctx, validator.Creator)
			k.RemoveValidator(ctx, validator.Creator)
			k.RemoveValidatorByConsAddr(ctx, validator.GetConsAddr().String())
		}

		// if last power was 0 and potential power more then 0 it means that validator was added in the current block.
		if lastValidatorPower.Power == 0 && validator.GetPower() > 0 {
			k.SetLastValidatorPower(ctx, types.LastValidatorPower{
				Index: validator.Creator,
				Power: validator.Power,
			})

			updates = append(updates, validator.ABCIValidatorUpdate())
		}
	}

	return updates, nil
}
