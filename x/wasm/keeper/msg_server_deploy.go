package keeper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/tendermint/tendermint/crypto/tmhash"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/liubaninc/m0/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Deploy(goCtx context.Context, msg *types.MsgDeploy) (*types.MsgDeployResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	msgIndex := int32(ctx.Context().Value("msg-index").(int))

	txHash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	if ok, err := VerifyTxRWSets(ctx, k.Keeper, &types.MsgInvoke{
		Creator:          msg.Creator,
		InputsExt:        msg.InputsExt,
		OutputsExt:       msg.OutputsExt,
		ContractRequests: []*types.InvokeRequest{msg.ConvertInvokeRequest()},
	}); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrRWSet, "verify %s", err)
	} else if !ok {
		return nil, errors.New("verifyTxRWSets failed")
	}

	for offset, outputExt := range msg.OutputsExt {
		if outputExt.Bucket == types.TransientBucket {
			continue
		}
		k.SetVersionedData(ctx, &xmodel.VersionedData{
			RefTxid:      []byte(txHash),
			RefMsgOffset: msgIndex,
			RefOffset:    int32(offset),
			PureData: &xmodel.PureData{
				Bucket: outputExt.Bucket,
				Key:    []byte(outputExt.Key),
				Value:  outputExt.Value,
			},
		})
	}

	var attrs []sdk.Attribute
	args, _ := json.Marshal(msg.Args)
	attrs = append(attrs,
		sdk.NewAttribute(types.AttributeKeyName, msg.ContractName),
		sdk.NewAttribute(types.AttributeKeyArg, string(args)),
	)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type()),
		),
		sdk.NewEvent(types.EventTypeDeployContract, attrs...),
	})
	return &types.MsgDeployResponse{}, nil
}
