package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/golang/protobuf/proto"
	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/bridge"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"
	"strconv"
)

func (k Keeper) GetContract(ctx sdk.Context, contract string) (*types.Contract, error) {
	verData, err := k.GetLastVersionedData(ctx, kernel.Contract2AccountBucket, []byte(contract))
	if err != nil {
		return nil, err
	}
	if verData == nil {
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
	return &types.Contract{
		Name:         contract,
		Initiator:    account,
		Number:       number,
		Address:      authtypes.NewModuleAddress(contract).String(),
		WasmCodeDesc: &desc,
	}, nil
}

func (k Keeper) IterateContracts(ctx sdk.Context, page uint, limit uint, cb func(item *types.Contract) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	prefix := append([]byte(types.ExtUtxoTablePrefix), types.MakeRawKey(kernel.Contract2AccountBucket, nil)...)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	if page > 0 && limit > 0 {
		iterator = sdk.KVStorePrefixIteratorPaginated(store, prefix, page, limit)
	}
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
		contract := &types.Contract{
			Name:         name,
			Number:       number,
			Initiator:    account,
			Address:      authtypes.NewModuleAddress(name).String(),
			WasmCodeDesc: &desc,
		}
		if cb(contract) {
			break
		}
	}
}

func (k Keeper) IterateAccountContracts(ctx sdk.Context, account string, page uint, limit uint, cb func(item *types.Contract) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	prefix := append([]byte(types.ExtUtxoTablePrefix), types.MakeRawKey(kernel.Account2ContractBucket, []byte(account+kernel.Account2ContractSeparator))...)
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	if page > 0 && limit > 0 {
		iterator = sdk.KVStorePrefixIteratorPaginated(store, prefix, page, limit)
	}
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
		contract := &types.Contract{
			Name:         name,
			Initiator:    account,
			Number:       number,
			Address:      authtypes.NewModuleAddress(name).String(),
			WasmCodeDesc: &desc,
		}
		if cb(contract) {
			break
		}
	}
}
