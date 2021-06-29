package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/blob/types"
	"strconv"
)

// GetBlobCount get the total number of TypeName.LowerCamel
func (k Keeper) GetBlobCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlobCountKey))
	byteKey := types.KeyPrefix(types.BlobCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to uint64
		panic("cannot decode count")
	}

	return count
}

// SetBlobCount set the total number of blob
func (k Keeper) SetBlobCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlobCountKey))
	byteKey := types.KeyPrefix(types.BlobCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendBlob appends a blob in the store with a new id and update the count
func (k Keeper) AppendBlob(
	ctx sdk.Context,
	blob types.Blob,
) uint64 {
	// Create the blob
	count := k.GetBlobCount(ctx)

	// Set the ID of the appended value
	blob.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlobKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&blob)
	store.Set(GetBlobIDBytes(blob.Id), appendedValue)

	// Update blob count
	k.SetBlobCount(ctx, count+1)

	return count
}

// SetBlob set a specific blob in the store
func (k Keeper) SetBlob(ctx sdk.Context, blob types.Blob) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlobKey))
	b := k.cdc.MustMarshalBinaryBare(&blob)
	store.Set(GetBlobIDBytes(blob.Id), b)
}

// GetBlob returns a blob from its id
func (k Keeper) GetBlob(ctx sdk.Context, id uint64) types.Blob {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlobKey))
	var blob types.Blob
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBlobIDBytes(id)), &blob)
	return blob
}

// HasBlob checks if the blob exists in the store
func (k Keeper) HasBlob(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlobKey))
	return store.Has(GetBlobIDBytes(id))
}

// GetBlobOwner returns the creator of the
func (k Keeper) GetBlobOwner(ctx sdk.Context, id uint64) string {
	return k.GetBlob(ctx, id).Creator
}

// RemoveBlob removes a blob from the store
func (k Keeper) RemoveBlob(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlobKey))
	store.Delete(GetBlobIDBytes(id))
}

// GetAllBlob returns all blob
func (k Keeper) GetAllBlob(ctx sdk.Context) (list []types.Blob) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlobKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Blob
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBlobIDBytes returns the byte representation of the ID
func GetBlobIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBlobIDFromBytes returns ID in uint64 format from a byte array
func GetBlobIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
