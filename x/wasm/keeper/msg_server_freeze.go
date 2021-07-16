package keeper

import (
	"context"

	"github.com/liubaninc/m0/x/wasm/types"
)

func (k msgServer) Freeze(goCtx context.Context, msg *types.MsgFreeze) (*types.MsgFreezeResponse, error) {

	SetStatus(goCtx, k, msg.GetContractName(), FREEZE)

	return &types.MsgFreezeResponse{}, nil
}
