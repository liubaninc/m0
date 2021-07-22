package keeper

import (
	"context"
	"github.com/liubaninc/m0/x/authority/types"
)

func (k msgServer) ApproveModifyAccount(goCtx context.Context, msg *types.MsgApproveModifyAccountRequest) (*types.MsgApproveModifyAccountResponse, error) {
	_, err := k.ApproveAddAccount(goCtx, &types.MsgApproveAddAccountRequest{
		Creator: msg.Creator,
		Address: msg.Address,
	})

	return &types.MsgApproveModifyAccountResponse{}, err
}
