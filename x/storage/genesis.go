package storage

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/storage/keeper"
	"github.com/liubaninc/m0/x/storage/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the recryptAccount
	for _, elem := range genState.RecryptAccountList {
		k.SetRecryptAccount(ctx, *elem)
	}

	// Set all the recryptAccount
	for _, elem := range genState.RecryptAccountList {
		k.SetRecryptAccount(ctx, *elem)
	}

	// Set all the shareEncryptStorage
	for _, elem := range genState.ShareEncryptStorageList {
		k.SetShareEncryptStorage(ctx, *elem)
	}

	// Set all the encryptStorage
	for _, elem := range genState.EncryptStorageList {
		k.SetEncryptStorage(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all recryptAccount
	recryptAccountList := k.GetAllRecryptAccount(ctx)
	for _, elem := range recryptAccountList {
		elem := elem
		genesis.RecryptAccountList = append(genesis.RecryptAccountList, &elem)
	}

	// Get all shareEncryptStorage
	shareEncryptStorageList := k.GetAllShareEncryptStorage(ctx)
	for _, elem := range shareEncryptStorageList {
		elem := elem
		genesis.ShareEncryptStorageList = append(genesis.ShareEncryptStorageList, &elem)
	}

	// Get all encryptStorage
	encryptStorageList := k.GetAllEncryptStorage(ctx)
	for _, elem := range encryptStorageList {
		elem := elem
		genesis.EncryptStorageList = append(genesis.EncryptStorageList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
