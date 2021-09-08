package keeper

import (
	"bytes"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/golang/protobuf/proto"
	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/bridge"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetContract(ctx sdk.Context, contract string) (*types.Contract, error) {
	verData, err := k.GetLastVersionedData(ctx, kernel.Contract2AccountBucket, []byte(contract))
	if err != nil {
		return nil, err
	}
	if verData == nil || types.IsDelFlag(verData.PureData.Value) {
		return nil, nil
	}

	account := string(verData.PureData.Value)
	verData, err = k.GetLastVersionedData(ctx, "contract", bridge.ContractCodeDescKey(contract))
	if err != nil {
		return nil, err
	}
	var desc xmodel.WasmCodeDesc
	if err := proto.Unmarshal(verData.PureData.Value, &desc); err != nil {
		return nil, err
	}
	verData, err = k.GetLastVersionedData(ctx, "contract", bridge.ContractNumberKey(contract))
	if err != nil {
		return nil, err
	}
	number, err := strconv.ParseUint(string(verData.PureData.Value), 10, 64)
	_ = number
	return &types.Contract{
		Name:      contract,
		Initiator: account,
		Desc:      &desc,
	}, nil
}

func (k Keeper) GetAllContract(ctx sdk.Context) (list []types.Contract) {
	store := ctx.KVStore(k.storeKey)
	prefix := append([]byte(types.ExtUtxoTablePrefix), types.MakeRawKey(kernel.Contract2AccountBucket, nil)...)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		verData, err := k.GetVersionedData(ctx, iterator.Value())
		if err != nil {
			panic(err)
		}
		account := string(verData.PureData.Value)
		name := string(iterator.Key()[len(prefix):])

		verData, err = k.GetLastVersionedData(ctx, "contract", bridge.ContractCodeDescKey(name))
		if err != nil {
			panic(err)
		}
		var desc xmodel.WasmCodeDesc
		if err := proto.Unmarshal(verData.PureData.Value, &desc); err != nil {
			panic(err)
		}
		verData, err = k.GetLastVersionedData(ctx, "contract", bridge.ContractNumberKey(name))
		if err != nil {
			panic(err)
		}
		number, err := strconv.ParseUint(string(verData.PureData.Value), 10, 64)
		_ = number
		contract := types.Contract{
			Name:      name,
			Initiator: account,
			Desc:      &desc,
		}
		list = append(list, contract)
	}
	return
}

func (k Keeper) SetContractState(ctx sdk.Context, name string, state uint16) {
	if state != types.Normarl && state != types.Freeze {
		panic("Incorrect state")
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatesKey))
	store.Set(types.KeyPrefix(name), []byte{byte(state)})
}

func (k Keeper) GetContractState(ctx sdk.Context, name string) (state uint16, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatesKey))

	b := store.Get(types.KeyPrefix(name))
	if b == nil {
		return 0, false
	}

	state = uint16(b[0])
	return state, true
}

func (k Keeper) RemoveContractState(ctx sdk.Context, name string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatesKey))

	store.Delete(types.KeyPrefix(name))
}

func (k Keeper) RemoveContract(ctx sdk.Context, contract *types.Contract, hash string, msgOffset int32) {

	k.DelLastVersionedData(ctx, kernel.Contract2AccountBucket, []byte(contract.Name))
	k.DelLastVersionedData(ctx, kernel.Account2ContractBucket, []byte(contract.Initiator+kernel.Account2ContractSeparator+contract.Name))
	k.DelLastVersionedData(ctx, "contract", bridge.ContractCodeKey(contract.Name))
	k.DelLastVersionedData(ctx, "contract", bridge.ContractCodeDescKey(contract.Name))
	k.DelLastVersionedData(ctx, "contract", bridge.ContractNumberKey(contract.Name))

	iter := k.IterVersioned(ctx, contract.Name, nil, nil)
	for ; iter.Valid(); iter.Next() {
		key := bytes.TrimPrefix(bytes.TrimPrefix(iter.Key(), []byte(types.ExtUtxoTablePrefix)), []byte(types.ExtUtxoDelTablePrefix))
		bucket, key, err := types.ParseRawKey(key)
		if err != nil {
			panic(err)
		}
		k.DelLastVersionedData(ctx, bucket, key)
	}
}
