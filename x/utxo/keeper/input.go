package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/utxo/types"
)

// SetInput set a specific input in the store from its index
func (k Keeper) SetInput(ctx sdk.Context, input types.Input) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InputKey))
	b := k.cdc.MustMarshalBinaryBare(&input)
	store.Set(types.KeyPrefix(input.Index()), b)
}

// GetInput returns a input from its index
func (k Keeper) GetInput(ctx sdk.Context, name string) (val types.Input, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InputKey))

	b := store.Get(types.KeyPrefix(name))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteInput removes a input from the store
func (k Keeper) RemoveInput(ctx sdk.Context, name string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InputKey))
	store.Delete(types.KeyPrefix(name))
}

// GetAllInput returns all input
func (k Keeper) GetAllInput(ctx sdk.Context) (list []types.Input) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InputKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Input
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SelectUtxos(ctx sdk.Context, addr sdk.AccAddress, totalNeed sdk.Coins, lock int64) (inputs []*types.Input, total sdk.Coins, err error) {
	if totalNeed.IsZero() {
		return
	}
	k.clearExpiredLocks()

	curBlockHeight := ctx.BlockHeight()
	total = sdk.NewCoins()
	for _, coin := range totalNeed {
		foundEnough := false
		foundTotal := sdk.ZeroInt()
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InputKey))
		iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(fmt.Sprintf("%X_%s", addr.Bytes(), coin.Denom)))
		defer iterator.Close()
		for ; iterator.Valid(); iterator.Next() {
			key := string(iterator.Key())
			var val types.Input
			k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)

			if k.isLocked(key) {
				k.Logger(ctx).Debug("select utxo, skip locked", "key", key)
				continue
			}
			if val.FrozenHeight > curBlockHeight || val.FrozenHeight == -1 {
				k.Logger(ctx).Debug("select utxo, skip frozen", "key", key, "fheight", val.FrozenHeight)
				continue
			}
			if lock > 0 && !k.tryLockKey(key, lock) {
				k.Logger(ctx).Debug("select utxos, skip unlock", "key", val.Index())
				continue
			}
			inputs = append(inputs, &val)
			foundTotal = foundTotal.Add(val.Amount.Amount)
			if foundTotal.GTE(coin.Amount) {
				foundEnough = true
				break
			}
		}
		if !foundEnough {
			if lock > 0 {
				for _, input := range inputs {
					k.unlockKey(input.Index())
				}
			}
			return nil, nil, sdkerrors.Wrapf(types.ErrNoEnoughUTXO, "address %s expect %s get %s", addr.String(), coin, sdk.NewCoin(coin.Denom, foundTotal))
		}
		total = total.Add(sdk.NewCoin(coin.Denom, foundTotal))
	}
	return
}
