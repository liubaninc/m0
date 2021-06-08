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

func (k msgServer) Deploy(goCtx context.Context, msg *types.MsgDeploy) (*types.MsgDeployResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	msgOffset := int32(ctx.Context().Value("msg-index").(int))
	txHash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	if err := k.utxoKeeper.Transfer(ctx, txHash, msgOffset, msg.Creator, msg.Inputs, msg.Outputs); err != nil {
		return nil, err
	}
	if err := k.RWSet(ctx, txHash, msgOffset, msg.Creator, msg.InputsExt, msg.OutputsExt, []*types.InvokeRequest{msg.ConvertInvokeRequest()}); err != nil {
		return nil, err
	}
	defer telemetry.IncrCounter(1, "new", "contract")
	argsStr, _ := json.Marshal(msg.Args)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
		),
		sdk.NewEvent(types.EventTypeContract, sdk.NewAttribute(types.AttributeKeyName, msg.ContractName),
			sdk.NewAttribute(types.AttributeKeyArg, string(argsStr))),
	})
	return &types.MsgDeployResponse{}, nil
}
