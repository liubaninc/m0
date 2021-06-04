package keeper

import (
	"container/list"
	"fmt"
	"sync"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/liubaninc/m0/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc           codec.Marshaler
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		paramSubspace paramtypes.Subspace
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper

		mutexMem    sync.Mutex             // 内存锁定状态互斥锁
		lockKeys    map[string]*LockedItem // 内存锁定状态
		lockKeyList *list.List             // 按锁定的过期时间先后顺序，方便过期清理
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	paramstore paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramSubspace: paramstore,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		lockKeys:      make(map[string]*LockedItem),
		lockKeyList:   list.New(),
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
