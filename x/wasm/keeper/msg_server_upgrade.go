package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/liubaninc/m0/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Upgrade(goCtx context.Context, msg *types.MsgUpgrade) (*types.MsgUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	msgOffset := int32(ctx.Context().Value("msg-index").(int))
	txHash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	t := time.Now()
	defer func() {
		k.Logger(ctx).Debug("handler", "route", msg.Route(), "msg", msg.Type(), "hash", txHash, "index", msgOffset, "elapsed", time.Now().Sub(t).String())
	}()
	if err := k.utxoKeeper.Transfer(ctx, txHash, msgOffset, msg.Creator, msg.Inputs, msg.Outputs); err != nil {
		return nil, err
	}
	if err := k.RWSet(ctx, txHash, msgOffset, msg.Creator, msg.InputsExt, msg.OutputsExt, []*types.InvokeRequest{msg.ConvertInvokeRequest()}); err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
		),
		sdk.NewEvent(types.EventTypeContract, sdk.NewAttribute(types.AttributeKeyName, msg.ContractName)),
	})
	return &types.MsgUpgradeResponse{}, nil
}
