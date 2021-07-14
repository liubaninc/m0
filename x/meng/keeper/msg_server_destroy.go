package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/meng/types"
)

func (k msgServer) Destroy(goCtx context.Context, msg *types.MsgDestroy) (*types.MsgDestroyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDestroyResponse{}, nil
}
