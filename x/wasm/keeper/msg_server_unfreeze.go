package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/wasm/types"
)

func (k msgServer) Unfreeze(goCtx context.Context, msg *types.MsgUnfreeze) (*types.MsgUnfreezeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if state, found := k.GetContractState(ctx, msg.ContractName); !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "contract not exist")
	} else if state == types.Normarl {
		return nil, nil
	}

	c, err := k.GetContract(ctx, msg.ContractName)
	if err != nil {
		panic(err)
	}
	if c.Initiator != msg.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "contract invalid owner")
	}

	k.SetContractState(ctx, msg.ContractName, types.Normarl)

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

	return &types.MsgUnfreezeResponse{}, nil
}
