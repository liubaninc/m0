package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/storage/types"
)

func (k msgServer) CreateShareEncryptStorage(goCtx context.Context, msg *types.MsgCreateShareEncryptStorage) (*types.MsgCreateShareEncryptStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.GetEncryptStorage(ctx, msg.Key); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "key already exist")
	}

	k.SetShareEncryptStorage(ctx, types.ShareEncryptStorage{
		Creator:    msg.Creator,
		Index:      msg.Key,
		ShareIndex: msg.ShareKey,
		Envelope:   msg.Envelope,
		Sharer:     msg.Sharer,
	})

	return &types.MsgCreateShareEncryptStorageResponse{}, nil
}
