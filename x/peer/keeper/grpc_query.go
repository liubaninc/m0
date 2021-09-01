package keeper

import (
	"github.com/liubaninc/m0/x/peer/types"
)

var _ types.QueryServer = Keeper{}
