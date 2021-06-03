package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/liubaninc/m0/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Upgrade(goCtx context.Context, msg *types.MsgUpgrade) (*types.MsgUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	txHash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	if ok, err := VerifyTxRWSets(ctx, k.Keeper, &types.MsgInvoke{
		Creator:          msg.Creator,
		InputsExt:        msg.InputsExt,
		OutputsExt:       msg.OutputsExt,
		ContractRequests: []*types.InvokeRequest{msg.ConvertInvokeRequest()},
	}); err != nil {
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
	attrs = append(attrs,
		sdk.NewAttribute(types.AttributeKeyName, msg.ContractName),
	)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type()),
		),
		sdk.NewEvent(types.EventTypeUpgradeContract, attrs...),
	})
	return &types.MsgUpgradeResponse{}, nil
}
