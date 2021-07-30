package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/liubaninc/m0/x/wasm/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Deploy(goCtx context.Context, msg *types.MsgDeploy) (*types.MsgDeployResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hash := tmhash.Sum(append(msg.ContractCode, append([]byte(msg.ContractName), []byte(msg.ContractName)...)...))
	index := fmt.Sprintf("%X", hash)
	if _, found := k.GetProposeDeploy(ctx, index); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "contract is proposing")
	}

	msgOffset := int32(ctx.Context().Value(baseapp.KeyMsgOffset).(int))
	txHash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	attrs, err := k.utxoKeeper.Transfer(ctx, txHash, msgOffset, msg.Creator, msg.Inputs, msg.Outputs)
	if err != nil {
		return nil, err
	}
	if err := k.RWSet(ctx, txHash, msgOffset, msg.Creator, msg.InputsExt, msg.OutputsExt, []*types.InvokeRequest{msg.ConvertInvokeRequest()}); err != nil {
		return nil, err
	}
	k.SetContractState(ctx, msg.ContractName, types.Normarl)

	argsStr, _ := json.Marshal(msg.Args)
	attrs = append(attrs,
		sdk.NewAttribute(types.AttributeKeyName, msg.ContractName),
		sdk.NewAttribute(types.AttributeKeyArg, string(argsStr)),
	)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			types.EventTypeDeploy,
			attrs...,
		),
	})
	return &types.MsgDeployResponse{}, nil
}
