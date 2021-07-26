package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/wasm/types"
)

func (k msgServer) Freeze(goCtx context.Context, msg *types.MsgFreeze) (*types.MsgFreezeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if state, found := k.GetContractState(ctx, msg.ContractName); !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "contract %s not exist", msg.ContractName)
	} else if state == types.Freeze {
		return nil, nil
	}

	c, err := k.GetContract(ctx, msg.ContractName)
	if err != nil {
		panic(err)
	}
	if c.Initiator != msg.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "contract %s invalid owner", msg.ContractName)
	}

	k.SetContractState(ctx, msg.ContractName, types.Freeze)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			types.EventTypeFreeze,
			sdk.NewAttribute(types.AttributeKeyName, msg.ContractName),
		),
	})

	return &types.MsgFreezeResponse{}, nil
}
