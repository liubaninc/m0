package keeper

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/telemetry"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/liubaninc/m0/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// GetAllInput returns all input
func (k Keeper) GetAllInputByAddress(ctx sdk.Context, addr sdk.AccAddress) (list []types.Input) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InputKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(fmt.Sprintf("%X_", addr.Bytes())))

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

func (k Keeper) Transfer(ctx sdk.Context, hash string, msgOffset int32, creator string, inputs []*types.Input, outputs []*types.Output) ([]sdk.Attribute, error) {
	totalIn := sdk.NewCoins()
	attrsIn := make([]sdk.Attribute, len(inputs))
	for index, input := range inputs {
		tinput, found := k.GetInput(ctx, input.Index())
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrUTXONotFound, "index %s in inputs", index)
		}
		if !tinput.Amount.Equal(input.Amount) {
			return nil, sdkerrors.Wrapf(types.ErrUTXONotMismatch, "index %d in inputs, amount expect %s get %s", index, input.Amount, tinput.Amount)
		}
		if tinput.FrozenHeight != input.FrozenHeight {
			return nil, sdkerrors.Wrapf(types.ErrUTXONotMismatch, "index %d in inputs, frozen height expect %s get %s", index, input.FrozenHeight, tinput.FrozenHeight)
		}
		if input.FrozenHeight == -1 || input.FrozenHeight > ctx.BlockHeight() {
			return nil, sdkerrors.Wrapf(types.ErrUTXOFrozen, "index %d in inputs, frozen height expect %s get %s", index, input.FrozenHeight, ctx.BlockHeight())
		}

		attrsIn[index] = sdk.NewAttribute(types.AttributeKeySender, input.FromAddr)
		totalIn = totalIn.Add(input.Amount)
		k.RemoveInput(ctx, input.Index())
		// bank
		addr, _ := sdk.AccAddressFromBech32(input.FromAddr)
		if err := k.bankKeeper.SubtractCoins(ctx, addr, sdk.NewCoins(input.Amount)); err != nil {
			return nil, err
		}
	}

	totalOut := sdk.NewCoins()
	attrsOut := make([]sdk.Attribute, len(outputs))
	for index, output := range outputs {
		attrsOut[index] = sdk.NewAttribute(types.AttributeKeyRecipient, output.ToAddr)
		totalOut = totalOut.Add(output.Amount)
		k.SetInput(ctx, types.Input{
			RefTx:        hash,
			RefMsg:       msgOffset,
			RefOffset:    int32(index),
			FromAddr:     output.ToAddr,
			Amount:       output.Amount,
			FrozenHeight: output.FrozenHeight,
		})
		// bank
		addr, _ := sdk.AccAddressFromBech32(output.ToAddr)
		if addr.Equals(authtypes.NewModuleAddress(authtypes.FeeCollectorName)) {
			// TODO chaogaofeng
		} else if err := k.bankKeeper.AddCoins(ctx, addr, sdk.NewCoins(output.Amount)); err != nil {
			return nil, err
		}
		acc := k.accountKeeper.GetAccount(ctx, addr)
		if acc == nil {
			defer telemetry.IncrCounter(1, "new", "account")
			k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, addr))
		}
	}

	attrs := make([]sdk.Attribute, totalOut.Len())
	for index, c := range totalOut {
		attrs[index] = sdk.NewAttribute(types.AttributeKeyDenom, c.Denom)
	}

	if totalIn.IsEqual(totalOut) {
		// send

		// attrs = append(attrs, sdk.NewAttribute(sdk.AttributeKeyAmount, totalIn.String()))
	} else if totalIn.IsAllGTE(totalOut) {
		// destroy
		changeCoins := totalIn.Sub(totalOut)
		for _, coin := range changeCoins {
			token, found := k.GetToken(ctx, coin.Denom)
			if !found {
				panic("token not found")
			}
			circulating, _ := sdk.NewIntFromString(token.Circulating)
			token.Circulating = circulating.Sub(coin.Amount).String()
			k.SetToken(ctx, token)
		}
		supply := k.bankKeeper.GetSupply(ctx)
		supply.Deflate(changeCoins)
		k.bankKeeper.SetSupply(ctx, supply)

		// attrs = append(attrs, sdk.NewAttribute(sdk.AttributeKeyAmount, totalIn.String()))
		attrs = append(attrs, sdk.NewAttribute(types.AttributeKeyAmountChanged, changeCoins.String()))
	} else {
		// issue/reissue
		changeCoins := totalOut.Sub(totalIn)
		for _, coin := range changeCoins {
			token, found := k.GetToken(ctx, coin.Denom)
			if found {
				if strings.Compare(token.Issuer, creator) != 0 {
					return nil, sdkerrors.Wrapf(types.ErrInvalidIssuer, "except %s get %s", token.Issuer, creator)
				}
				supply, _ := sdk.NewIntFromString(token.Supply)
				token.Supply = supply.Add(coin.Amount).String()
				circulating, _ := sdk.NewIntFromString(token.Circulating)
				token.Circulating = circulating.Add(coin.Amount).String()
			} else {
				token = types.Token{
					Name:        coin.Denom,
					Issuer:      creator,
					IssueTx:     hash,
					Supply:      coin.Amount.String(),
					Circulating: coin.Amount.String(),
				}
				defer telemetry.IncrCounter(1, "new", "token")
			}
			k.SetToken(ctx, token)
		}
		supply := k.bankKeeper.GetSupply(ctx)
		supply.Inflate(changeCoins)
		k.bankKeeper.SetSupply(ctx, supply)

		// attrs = append(attrs, sdk.NewAttribute(sdk.AttributeKeyAmount, totalOut.String()))
		attrs = append(attrs, sdk.NewAttribute(types.AttributeKeyAmountChanged, changeCoins.String()))
	}
	return append(append(attrsIn, attrsOut...), attrs...), nil
}
