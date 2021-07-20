package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/validator/types"
)

func (k msgServer) EditValidator(goCtx context.Context, msg *types.MsgEditValidator) (*types.MsgEditValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasValidator(ctx, msg.Creator) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "address %s is not a validator", msg.Creator)
	}

	validator := k.GetValidator(ctx, msg.Creator)
	if msg.Description.Moniker != types.DoNotModifyDesc {
		validator.Description.Moniker = msg.Description.Moniker
	}
	if msg.Description.Website != types.DoNotModifyDesc {
		validator.Description.Website = msg.Description.Website
	}
	if msg.Description.Details != types.DoNotModifyDesc {
		validator.Description.Details = msg.Description.Details
	}
	if msg.Description.Identity != types.DoNotModifyDesc {
		validator.Description.Identity = msg.Description.Identity
	}
	k.SetValidator(ctx, validator)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, validator.GetConsAddr().String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgEditValidatorResponse{}, nil
}
