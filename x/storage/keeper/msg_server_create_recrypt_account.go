package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/storage/types"
)

func (k msgServer) CreateRecryptAccount(goCtx context.Context, msg *types.MsgCreateRecryptAccount) (*types.MsgCreateRecryptAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.SetRecryptAccount(ctx, types.RecryptAccount{
		Creator: msg.Creator,
		Index:   msg.Creator,
		Pubkey:  msg.Pubkey,
	})

	return &types.MsgCreateRecryptAccountResponse{}, nil
}
