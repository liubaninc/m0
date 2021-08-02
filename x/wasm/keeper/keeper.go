package keeper

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

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
		paramspace paramtypes.Subspace
		utxoKeeper utxokeeper.Keeper
		vmMgr      *contract.VMManager
		Ctx        sdk.Context
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	paramspace paramtypes.Subspace,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	utxoKeeper utxokeeper.Keeper,
) *Keeper {
	keeper := &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
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
