package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/liubaninc/m0/x/wasm/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Upgrade(goCtx context.Context, msg *types.MsgUpgrade) (*types.MsgUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if state, found := k.GetContractState(ctx, msg.ContractName); !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "contract %s not exist", msg.ContractName)
	} else if state != types.Normarl {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "contract %s was not normal state", msg.ContractName)
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
	attrs = append(attrs, sdk.NewAttribute(types.AttributeKeyName, msg.ContractName))
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			types.EventTypeUpgrade,
			attrs...,
		),
	})
	return &types.MsgUpgradeResponse{}, nil
}
