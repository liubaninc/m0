package keeper

import (
	"context"

	"github.com/liubaninc/m0/x/wasm/types"
)

func (k msgServer) Unfreeze(goCtx context.Context, msg *types.MsgUnfreeze) (*types.MsgUnfreezeResponse, error) {

	SetStatus(goCtx, k, msg.GetContractName(), UNFREEZE)

	return &types.MsgUnfreezeResponse{}, nil
}
