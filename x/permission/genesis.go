package permission

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/permission/keeper"
	"github.com/liubaninc/m0/x/permission/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	// this line is used by starport scaffolding # genesis/module/init
	// Set all the account
	for _, elem := range genState.AccountList {
		k.SetAccount(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export
	// Get all account
	accountList := k.GetAllAccount(ctx)
	for _, elem := range accountList {
		elem := elem
		genesis.AccountList = append(genesis.AccountList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
