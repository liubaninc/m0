package utxo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/utxo/keeper"
	"github.com/liubaninc/m0/x/utxo/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	for _, elem := range genState.TokenList {
		k.SetToken(ctx, *elem)
	}

	for _, input := range genState.InputList {
		k.SetInput(ctx, *input)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Params = k.GetParams(ctx)

	tokenList := k.GetAllToken(ctx)
	for _, elem := range tokenList {
		elem := elem
		genesis.TokenList = append(genesis.TokenList, &elem)
	}

	inputList := k.GetAllInput(ctx)
	for _, elem := range inputList {
		elem := elem
		genesis.InputList = append(genesis.InputList, &elem)
	}

	return genesis
}
