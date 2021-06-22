package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/mibc/types"
	"strconv"
)

// GetItxCount get the total number of TypeName.LowerCamel
func (k Keeper) GetItxCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItxCountKey))
	byteKey := types.KeyPrefix(types.ItxCountKey)
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

// SetItxCount set the total number of itx
func (k Keeper) SetItxCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItxCountKey))
	byteKey := types.KeyPrefix(types.ItxCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendItx appends a itx in the store with a new id and update the count
func (k Keeper) AppendItx(
	ctx sdk.Context,
	itx types.Itx,
) uint64 {
	// Create the itx
	count := k.GetItxCount(ctx)

	// Set the ID of the appended value
	itx.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItxKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&itx)
	store.Set(GetItxIDBytes(itx.Id), appendedValue)

	// Update itx count
	k.SetItxCount(ctx, count+1)

	return count
}

// SetItx set a specific itx in the store
func (k Keeper) SetItx(ctx sdk.Context, itx types.Itx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItxKey))
	b := k.cdc.MustMarshalBinaryBare(&itx)
	store.Set(GetItxIDBytes(itx.Id), b)
}

// GetItx returns a itx from its id
func (k Keeper) GetItx(ctx sdk.Context, id uint64) types.Itx {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItxKey))
	var itx types.Itx
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetItxIDBytes(id)), &itx)
	return itx
}

// HasItx checks if the itx exists in the store
func (k Keeper) HasItx(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItxKey))
	return store.Has(GetItxIDBytes(id))
}

// GetItxOwner returns the creator of the
func (k Keeper) GetItxOwner(ctx sdk.Context, id uint64) string {
	return k.GetItx(ctx, id).Creator
}

// RemoveItx removes a itx from the store
func (k Keeper) RemoveItx(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItxKey))
	store.Delete(GetItxIDBytes(id))
}

// GetAllItx returns all itx
func (k Keeper) GetAllItx(ctx sdk.Context) (list []types.Itx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItxKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Itx
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetItxIDBytes returns the byte representation of the ID
func GetItxIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetItxIDFromBytes returns ID in uint64 format from a byte array
func GetItxIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
