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

func (k msgServer) Invoke(goCtx context.Context, msg *types.MsgInvoke) (*types.MsgInvokeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, cr := range msg.ContractRequests {
		if state, found := k.GetContractState(ctx, cr.ContractName); !found {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "contract %s not exist", cr.ContractName)
		} else if state != types.Normarl {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "contract %s was not normal state", cr.ContractName)
		}
	}

	msgOffset := int32(ctx.Context().Value(baseapp.KeyMsgOffset).(int))
	txHash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	attrs, err := k.utxoKeeper.Transfer(ctx, txHash, msgOffset, msg.Creator, msg.Inputs, msg.Outputs)
	if err != nil {
		return nil, err
	}
	if err := k.RWSet(ctx, txHash, msgOffset, msg.Creator, msg.InputsExt, msg.OutputsExt, msg.ContractRequests); err != nil {
		return nil, err
	}
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
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			types.EventTypeInvoke,
			attrs...,
		),
	})
	return &types.MsgInvokeResponse{}, nil
}
