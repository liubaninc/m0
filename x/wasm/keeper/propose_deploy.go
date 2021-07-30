package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/wasm/types"
)

// SetProposeDeploy set a specific proposeDeploy in the store from its index
func (k Keeper) SetProposeDeploy(ctx sdk.Context, proposeDeploy types.ProposeDeploy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposeDeployKey))
	b := k.cdc.MustMarshalBinaryBare(&proposeDeploy)
	store.Set(types.KeyPrefix(proposeDeploy.Index), b)
}

// GetProposeDeploy returns a proposeDeploy from its index
func (k Keeper) GetProposeDeploy(ctx sdk.Context, index string) (val types.ProposeDeploy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposeDeployKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteProposeDeploy removes a proposeDeploy from the store
func (k Keeper) RemoveProposeDeploy(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposeDeployKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllProposeDeploy returns all proposeDeploy
func (k Keeper) GetAllProposeDeploy(ctx sdk.Context) (list []types.ProposeDeploy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposeDeployKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProposeDeploy
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
