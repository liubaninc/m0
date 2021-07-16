package keeper

import (
	"context"
	"fmt"

	utxokeeper "github.com/liubaninc/m0/x/utxo/keeper"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/bridge"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/wasm/xvm"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/liubaninc/m0/x/wasm/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute

		utxoKeeper utxokeeper.Keeper
		vmMgr      *contract.VMManager
		Ctx        sdk.Context
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	utxoKeeper utxokeeper.Keeper,
) *Keeper {
	keeper := &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
		utxoKeeper: utxoKeeper,
		vmMgr:      contract.NewVMManager(),
	}
	// 合约虚拟机
	basedir := viper.GetString(flags.FlagHome)
	bridge, err := bridge.New(&bridge.XBridgeConfig{
		Basedir: basedir,
		VMConfigs: map[bridge.ContractType]bridge.VMConfig{
			bridge.TypeWasm: &xvm.WasmConfig{
				Driver: "xvm",
				XVM: xvm.XVMConfig{
					OptLevel: 1,
				},
			},
		},
		Config: bridge.ContractConfig{
			EnableUpgrade: true,
		},
	})
	if err != nil {
		panic(err)
	}
	bridge.RegisterToXCore(keeper.vmMgr.RegisterVM)

	// 系统合约
	kernel, err := kernel.NewKernel(bridge)
	if err != nil {
		panic(err)
	}
	keeper.vmMgr.RegisterVM(kernel.GetName(), kernel)
	return keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

const DEPLOY = "0"
const FREEZE = "1"
const UNFREEZE = "2"
const DESTROY = "3"

/**
部署 	0
冻结 	1 freeze
解冻		2 unfreeze
销毁		3 destroy
*/
func SetStatus(ctx context.Context, k msgServer, contractName string, status string) string {
	ctx2 := sdk.UnwrapSDKContext(ctx)
	store := ctx2.KVStore(k.storeKey)
	nameByte := []byte(contractName)
	statusByte := []byte(status)
	store.Set(nameByte, statusByte)

	fmt.Sprintf("change %s status to %s", contractName, status)

	return ""
}
