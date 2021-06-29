package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/liubaninc/m0/x/blob/types"
)

func (k msgServer) CommentBlog(goCtx context.Context, msg *types.MsgCommentBlog) (*types.MsgCommentBlogResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Keeper.AppendComment(ctx, types.Comment{
		Creator: msg.Creator,
		BlobId:  msg.BlobId,
		Comment: msg.Comment,
	})

	return &types.MsgCommentBlogResponse{}, nil
}
