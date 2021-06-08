package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/liubaninc/m0/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Invoke(goCtx context.Context, msg *types.MsgInvoke) (*types.MsgInvokeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	msgOffset := int32(ctx.Context().Value("msg-index").(int))
	txHash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	if err := k.utxoKeeper.Transfer(ctx, txHash, msgOffset, msg.Creator, msg.Inputs, msg.Outputs); err != nil {
		return nil, err
	}
	if err := k.RWSet(ctx, txHash, msgOffset, msg.Creator, msg.InputsExt, msg.OutputsExt, msg.ContractRequests); err != nil {
		return nil, err
	}
	defer telemetry.IncrCounter(1, "invoke", "contract")
	var attrs []sdk.Attribute
	for _, request := range msg.ContractRequests {
		args, _ := json.Marshal(request.Args)
		attrs = append(attrs,
			sdk.NewAttribute(types.AttributeKeyName, request.ContractName),
			sdk.NewAttribute(types.AttributeKeyMethod, request.MethodName),
			sdk.NewAttribute(types.AttributeKeyArg, string(args)),
		)
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
		),
		sdk.NewEvent(types.EventTypeContract, attrs...),
	})
	return &types.MsgInvokeResponse{}, nil
}
