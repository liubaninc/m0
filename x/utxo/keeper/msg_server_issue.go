package keeper

import (
	"context"
	"fmt"

	"github.com/liubaninc/m0/x/utxo/types"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	msgOffset := int32(ctx.Context().Value(baseapp.KeyMsgOffset).(int))
	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	attrs, err := k.Transfer(ctx, hash, msgOffset, msg.Creator, msg.Inputs, msg.Outputs)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			msg.Type(),
			attrs...,
		),
	})
	return &types.MsgIssueResponse{}, nil
}
