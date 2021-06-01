package keeper

import (
	"context"
	"github.com/liubaninc/m0/x/wasm/types"
)

func (k Keeper) PreExec(c context.Context, request *types.InvokeRPCRequest) (*types.InvokeRPCResponse, error) {
	return &types.InvokeRPCResponse{}, nil
}
