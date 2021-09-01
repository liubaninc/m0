package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/permission/types"
)

// SetParams sets the auth module's parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}

// GetParams gets the auth module's parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return
}
