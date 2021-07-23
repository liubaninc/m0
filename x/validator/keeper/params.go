package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/validator/types"
	"time"
)

// UnbondingTime
func (k Keeper) UnbondingTime(ctx sdk.Context) (res time.Duration) {
	k.paramspace.Get(ctx, types.KeyUnbondingTime, &res)
	return
}

// MaxValidators - Maximum number of validators
func (k Keeper) MaxValidators(ctx sdk.Context) (res uint32) {
	k.paramspace.Get(ctx, types.KeyMaxValidators, &res)
	return
}

// Quorum - Quorum percentage to add or remove validators
func (k Keeper) MinValidators(ctx sdk.Context) (res uint32) {
	k.paramspace.Get(ctx, types.KeyMinValidators, &res)
	return
}

// GetParams returns the total set of poa parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the poa parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}

// HistoricalEntries = number of historical info entries
// to persist in store
func (k Keeper) HistoricalEntries(ctx sdk.Context) (res uint32) {
	k.paramspace.Get(ctx, types.KeyHistoricalEntries, &res)
	return
}
