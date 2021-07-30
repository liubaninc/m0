package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/tendermint/tendermint/crypto/tmhash"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/wasm/types"
)

func (k msgServer) Undeploy(goCtx context.Context, msg *types.MsgUndeploy) (*types.MsgUndeployResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	c, err := k.GetContract(ctx, msg.ContractName)
	if err != nil {
		panic(err)
	}
	if c.Initiator != msg.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "contract %s invalid owner", msg.ContractName)
	}

	tinputs := k.utxoKeeper.GetAllInputByAddress(ctx, authtypes.NewModuleAddress(msg.ContractName))
	if n := len(tinputs); n > 0 {
		totalCoins := sdk.NewCoins()
		inputs := make([]*utxotypes.Input, len(tinputs))
		for i, input := range tinputs {
			totalCoins = totalCoins.Add(input.Amount)
			inputs[i] = &input
		}

		outputs := make([]*utxotypes.Output, totalCoins.Len())
		for i, coin := range totalCoins {
			outputs[i] = &utxotypes.Output{
				ToAddr: msg.Creator,
				Amount: coin,
			}
		}
		msgOffset := int32(ctx.Context().Value(baseapp.KeyMsgOffset).(int))
		txHash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
		k.utxoKeeper.Transfer(ctx, txHash, msgOffset, msg.Creator, inputs, outputs)
	}

	k.RemoveContract(ctx, msg.ContractName)
	k.RemoveContractState(ctx, msg.ContractName)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			msg.Type(),
			sdk.NewAttribute(types.AttributeKeyName, msg.ContractName),
		),
	})

	return &types.MsgUndeployResponse{}, nil
}
