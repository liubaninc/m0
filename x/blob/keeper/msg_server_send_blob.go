package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/blob/types"
)

func (k msgServer) SendBlob(goCtx context.Context, msg *types.MsgSendBlob) (*types.MsgSendBlobResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Keeper.AppendBlob(ctx, types.Blob{
		Creator: msg.Creator,
		Title:   msg.Title,
		Content: msg.Content,
	})

	return &types.MsgSendBlobResponse{}, nil
}
