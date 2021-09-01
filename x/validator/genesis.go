package validator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/validator/keeper"
	"github.com/liubaninc/m0/x/validator/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	// this line is used by starport scaffolding # genesis/module/init
	// Set all the lastValidatorPower
	for _, elem := range genState.LastValidatorPowerList {
		k.SetLastValidatorPower(ctx, *elem)
	}

	// Set all the validator
	for _, elem := range genState.ValidatorList {
		k.SetValidator(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export
	// Get all lastValidatorPower
	lastValidatorPowerList := k.GetAllLastValidatorPower(ctx)
	for _, elem := range lastValidatorPowerList {
		elem := elem
		genesis.LastValidatorPowerList = append(genesis.LastValidatorPowerList, &elem)
	}

	// Get all validato
	validatorList := k.GetAllValidator(ctx)
	for _, elem := range validatorList {
		elem := elem
		genesis.ValidatorList = append(genesis.ValidatorList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
