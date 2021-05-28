package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/wasm/types"
)

func (k msgServer) Deploy(goCtx context.Context, msg *types.MsgDeploy) (*types.MsgDeployResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeployResponse{}, nil
}
