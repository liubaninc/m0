package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/validator/types"
	// this line is used by starport scaffolding # ibc/keeper/import
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute
		paramspace paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	paramspace paramtypes.Subspace,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
		paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
