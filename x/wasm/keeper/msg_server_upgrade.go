package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/wasm/types"
)

func (k msgServer) Upgrade(goCtx context.Context, msg *types.MsgUpgrade) (*types.MsgUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpgradeResponse{}, nil
}
