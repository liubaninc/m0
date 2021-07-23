package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/liubaninc/m0/x/validator/types"
	"strconv"
)

// GetHistoricalInfo gets the historical info at a given height
func (k Keeper) GetHistoricalInfo(ctx sdk.Context, height int64) (val stakingtypes.HistoricalInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HistoricalInfoKey))

	value := store.Get([]byte(strconv.FormatInt(height, 10)))
	if value == nil {
		return stakingtypes.HistoricalInfo{}, false
	}

	k.cdc.MustUnmarshalBinaryBare(value, &val)
	return val, true
}

// SetHistoricalInfo sets the historical info at a given height
func (k Keeper) SetHistoricalInfo(ctx sdk.Context, height int64, hi *stakingtypes.HistoricalInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HistoricalInfoKey))

	value := k.cdc.MustMarshalBinaryBare(hi)
	store.Set([]byte(strconv.FormatInt(height, 10)), value)
}

// DeleteHistoricalInfo deletes the historical info at a given height
func (k Keeper) DeleteHistoricalInfo(ctx sdk.Context, height int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HistoricalInfoKey))

	store.Delete([]byte(strconv.FormatInt(height, 10)))
}

// IterateHistoricalInfo provides an interator over all stored HistoricalInfo
//  objects. For each HistoricalInfo object, cb will be called. If the cb returns
// true, the iterator will close and stop.
func (k Keeper) IterateHistoricalInfo(ctx sdk.Context, cb func(stakingtypes.HistoricalInfo) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.HistoricalInfoKey))

	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var histInfo stakingtypes.HistoricalInfo
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &histInfo)
		if cb(histInfo) {
			break
		}
	}
}

// GetAllHistoricalInfo returns all stored HistoricalInfo objects.
func (k Keeper) GetAllHistoricalInfo(ctx sdk.Context) []stakingtypes.HistoricalInfo {
	var infos []stakingtypes.HistoricalInfo

	k.IterateHistoricalInfo(ctx, func(histInfo stakingtypes.HistoricalInfo) bool {
		infos = append(infos, histInfo)
		return false
	})

	return infos
}

// TrackHistoricalInfo saves the latest historical-info and deletes the oldest
// heights that are below pruning height
func (k Keeper) TrackHistoricalInfo(ctx sdk.Context) {
	entryNum := k.HistoricalEntries(ctx)

	// Prune store to ensure we only have parameter-defined historical entries.
	// In most cases, this will involve removing a single historical entry.
	// In the rare scenario when the historical entries gets reduced to a lower value k'
	// from the original value k. k - k' entries must be deleted from the store.
	// Since the entries to be deleted are always in a continuous range, we can iterate
	// over the historical entries starting from the most recent version to be pruned
	// and then return at the first empty entry.
	for i := ctx.BlockHeight() - int64(entryNum); i >= 0; i-- {
		_, found := k.GetHistoricalInfo(ctx, i)
		if found {
			k.DeleteHistoricalInfo(ctx, i)
		} else {
			break
		}
	}

	// if there is no need to persist historicalInfo, return
	if entryNum == 0 {
		return
	}

	// Create HistoricalInfo struct
	lastValidatorPowers := k.GetAllLastValidatorPower(ctx)
	historicalEntry := stakingtypes.HistoricalInfo{
		Header: ctx.BlockHeader(),
		Valset: func() (lastVals []stakingtypes.Validator) {
			for _, lastValidatorPower := range lastValidatorPowers {
				validator := k.GetValidator(ctx, lastValidatorPower.Index)
				pk := sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, validator.PubKey)
				pkAny, err := codectypes.NewAnyWithValue(pk)
				if err != nil {
					panic(err)
				}

				lastVals = append(lastVals, stakingtypes.Validator{
					OperatorAddress: validator.Creator,
					ConsensusPubkey: pkAny,
					Jailed:          validator.Jailed,
				})
			}
			return
		}(),
	}

	// Set latest HistoricalInfo at current height
	k.SetHistoricalInfo(ctx, ctx.BlockHeight(), &historicalEntry)
}
