package keeper

import (
	"context"
	"fmt"
	"github.com/liubaninc/m0/x/utxo/types"
	"github.com/tendermint/tendermint/crypto/tmhash"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	msgOffset:= int32(ctx.Context().Value("msg-index").(int))
	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	if err := k.Transfer(ctx, hash, msgOffset, msg.Creator, msg.Inputs, msg.Outputs); err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
		),
	})
	return &types.MsgSendResponse{}, nil
}
