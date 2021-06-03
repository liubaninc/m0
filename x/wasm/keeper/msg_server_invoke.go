package keeper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/liubaninc/m0/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Invoke(goCtx context.Context, msg *types.MsgInvoke) (*types.MsgInvokeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	txHash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	if ok, err := VerifyTxRWSets(ctx, k.Keeper, msg); err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("verifyTxRWSets failed")
	}
	// TODO chaogaofeng
	msgOffset := int32(0)
	for offset, outputExt := range msg.OutputsExt {
		if outputExt.Bucket == types.TransientBucket {
			continue
		}
		k.SetVersionedData(ctx, &xmodel.VersionedData{
			RefTxid:      []byte(txHash),
			RefMsgOffset: int32(msgOffset),
			RefOffset:    int32(offset),
			PureData: &xmodel.PureData{
				Bucket: outputExt.Bucket,
				Key:    []byte(outputExt.Key),
				Value:  outputExt.Value,
			},
		})
	}

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
			sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type()),
		),
		sdk.NewEvent(types.EventTypeInvokeContract, attrs...),
	})
	return &types.MsgInvokeResponse{}, nil
}
