package keeper

import (
	"context"

	"github.com/liubaninc/m0/x/wasm/types"
)

func (k msgServer) Destroy(goCtx context.Context, msg *types.MsgDestroy) (*types.MsgDestroyResponse, error) {

	SetStatus(goCtx, k, msg.GetContractName(), DESTROY)

	return &types.MsgDestroyResponse{}, nil
}
