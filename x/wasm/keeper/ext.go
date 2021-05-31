package keeper

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
)

func (k Keeper) SetVersionedData(ctx sdk.Context, data *xmodel.VersionedData) {
	store := ctx.KVStore(k.storeKey)
	bucketAndKey := types.MakeRawKey(data.PureData.Bucket, data.PureData.Key)
	versioned := types.MakeVersion(data.RefTxid, data.RefMsgOffset, data.RefOffset)
	if types.IsDelFlag(data.PureData.Value) {
		store.Delete(append([]byte(types.ExtUtxoTablePrefix), bucketAndKey...))
		store.Set(append([]byte(types.ExtUtxoDelTablePrefix), bucketAndKey...), []byte(versioned))
	} else {
		store.Set(append([]byte(types.ExtUtxoTablePrefix), bucketAndKey...), []byte(versioned))
	}
	bz := k.cdc.MustMarshalBinaryBare(data)
	store.Set([]byte(types.ExtUtxoValueTablePrefix+versioned), bz)
}

func (k Keeper) GetVersionedData(ctx sdk.Context, versioned []byte) (*xmodel.VersionedData, error) {
	if _, _, _, err := types.ParseVersion(string(versioned)); err != nil {
		return nil, err
	}
	store := ctx.KVStore(k.storeKey)
	verData := xmodel.VersionedData{}
	val := store.Get(append([]byte(types.ExtUtxoValueTablePrefix), versioned...))
	if err := k.cdc.UnmarshalBinaryBare(val, &verData); err != nil {
		return nil, err
	}
	return &verData, nil
}

func (k Keeper) GetLastVersionedData(ctx sdk.Context, bucket string, key []byte) (*xmodel.VersionedData, error) {
	store := ctx.KVStore(k.storeKey)
	bucketAndKey := types.MakeRawKey(bucket, key)

	versioned := store.Get(append([]byte(types.ExtUtxoTablePrefix), bucketAndKey...))
	if versioned == nil {
		versioned = store.Get(append([]byte(types.ExtUtxoDelTablePrefix), bucketAndKey...))
	}
	if versioned == nil {
		return nil, nil
	}
	return k.GetVersionedData(ctx, versioned)
}

func (k Keeper) IterVersioned(ctx sdk.Context, bucket string, startKey []byte, endKey []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	rawStartKey := types.MakeRawKey(bucket, startKey)
	if len(endKey) == 0 {
		return storetypes.KVStoreReversePrefixIterator(store, append([]byte(types.ExtUtxoTablePrefix), rawStartKey...))
	}
	rawEndKey := types.MakeRawKey(bucket, endKey)
	return store.Iterator(append([]byte(types.ExtUtxoTablePrefix), rawStartKey...), append([]byte(types.ExtUtxoTablePrefix), rawEndKey...))
}
