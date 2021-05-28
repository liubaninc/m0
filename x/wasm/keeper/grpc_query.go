package keeper

import (
	"github.com/liubaninc/m0/x/wasm/types"
)

var _ types.QueryServer = Keeper{}
