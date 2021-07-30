package wasm

import (
	"github.com/liubaninc/m0/x/wasm/keeper"
	"github.com/liubaninc/m0/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the approveDeploy
	for _, elem := range genState.ApproveDeployList {
		k.SetApproveDeploy(ctx, *elem)
	}

	// Set all the proposeDeploy
	for _, elem := range genState.ProposeDeployList {
		k.SetProposeDeploy(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all approveDeploy
	approveDeployList := k.GetAllApproveDeploy(ctx)
	for _, elem := range approveDeployList {
		elem := elem
		genesis.ApproveDeployList = append(genesis.ApproveDeployList, &elem)
	}

	// Get all proposeDeploy
	proposeDeployList := k.GetAllProposeDeploy(ctx)
	for _, elem := range proposeDeployList {
		elem := elem
		genesis.ProposeDeployList = append(genesis.ProposeDeployList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
