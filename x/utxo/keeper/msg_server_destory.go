package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/liubaninc/m0/x/utxo/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Destroy(goCtx context.Context, msg *types.MsgDestroy) (*types.MsgDestroyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	msgOffset := int32(ctx.Context().Value("msg-index").(int))
	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	t := time.Now()
	defer func() {
		k.Logger(ctx).Debug("handler", "route", msg.Route(), "msg", msg.Type(), "hash", hash, "index", msgOffset, "elapsed", time.Now().Sub(t).String())
	}()
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

	return &types.MsgDestroyResponse{}, nil
}
