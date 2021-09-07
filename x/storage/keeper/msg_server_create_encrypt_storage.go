package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/storage/types"
)

func (k msgServer) CreateEncryptStorage(goCtx context.Context, msg *types.MsgCreateEncryptStorage) (*types.MsgCreateEncryptStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.GetEncryptStorage(ctx, msg.Key); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "key already exist")
	}

	k.SetEncryptStorage(ctx, types.EncryptStorage{
		Creator:   msg.Creator,
		Index:     msg.Key,
		Md5:       msg.Md5,
		Encrypted: msg.Encrypted,
		Envelope:  msg.Envelope,
	})

	return &types.MsgCreateEncryptStorageResponse{}, nil
}
