package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/validator/types"
)

func (k msgServer) LeaveValidator(goCtx context.Context, msg *types.MsgLeaveValidator) (*types.MsgLeaveValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	minValidator := k.MinValidators(ctx)
	if uint16(len(k.GetAllValidator(ctx))) == minValidator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "the minimum number of validators has been reached")
	}

	if !k.HasValidator(ctx, msg.Creator) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "address %s is not a validator", msg.Creator)
	}

	validator := k.GetValidator(ctx, msg.Creator)
	validator.Power = 0
	k.SetValidator(ctx, validator)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeLeaveValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, validator.GetConsAddr().String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgLeaveValidatorResponse{}, nil
}
