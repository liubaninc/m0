package keeper

import (
	"github.com/liubaninc/m0/x/utxo/types"
)

var _ types.QueryServer = Keeper{}
