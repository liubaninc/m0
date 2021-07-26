package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/peer/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (k Keeper) IDPeerFilter(ctx sdk.Context, index string) abci.ResponseQuery {
	if len(k.GetAllPeerID(ctx)) == 0 {
		return abci.ResponseQuery{
			Code: abci.CodeTypeOK,
		}
	}

	code := abci.CodeTypeOK
	var err error
	peerID, found := k.GetPeerID(ctx, index)
	if !found {
		err = fmt.Errorf("id %s not found", index)
		code = 1
	}
	_ = peerID

	return abci.ResponseQuery{
		Code: code,
		Log:  err.Error(),
	}
}

// SetPeerID set a specific peerID in the store from its index
func (k Keeper) SetPeerID(ctx sdk.Context, peerID types.PeerID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeerIDKey))
	b := k.cdc.MustMarshalBinaryBare(&peerID)
	store.Set(types.KeyPrefix(peerID.Index), b)
}

// GetPeerID returns a peerID from its index
func (k Keeper) GetPeerID(ctx sdk.Context, index string) (val types.PeerID, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeerIDKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeletePeerID removes a peerID from the store
func (k Keeper) RemovePeerID(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeerIDKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllPeerID returns all peerID
func (k Keeper) GetAllPeerID(ctx sdk.Context) (list []types.PeerID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeerIDKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PeerID
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
