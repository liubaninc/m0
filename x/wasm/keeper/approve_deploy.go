package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/wasm/types"
)

// SetApproveDeploy set a specific approveDeploy in the store from its index
func (k Keeper) SetApproveDeploy(ctx sdk.Context, approveDeploy types.ApproveDeploy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApproveDeployKey))
	b := k.cdc.MustMarshalBinaryBare(&approveDeploy)
	store.Set(types.KeyPrefix(approveDeploy.Index), b)
}

// GetApproveDeploy returns a approveDeploy from its index
func (k Keeper) GetApproveDeploy(ctx sdk.Context, index string) (val types.ApproveDeploy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApproveDeployKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteApproveDeploy removes a approveDeploy from the store
func (k Keeper) RemoveApproveDeploy(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApproveDeployKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllApproveDeploy returns all approveDeploy
func (k Keeper) GetAllApproveDeploy(ctx sdk.Context) (list []types.ApproveDeploy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ApproveDeployKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ApproveDeploy
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
