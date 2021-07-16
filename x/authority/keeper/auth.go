package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type ValidateAuthRolesDecorator struct {
	keeper Keeper
}

func NewValidateAuthRolesDecorator(keeper Keeper) ValidateAuthRolesDecorator {
	return ValidateAuthRolesDecorator{keeper: keeper}
}

func (vad ValidateAuthRolesDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if ctx.BlockHeight() == 0 {
		return next(ctx, tx, simulate)
	}
	for _, msg := range tx.GetMsgs() {
		route := msg.Route()
		for _, acc := range msg.GetSigners() {
			if !vad.keeper.HasRole(ctx, acc.String(), route) {
				return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "invalid transaction auth")
			}
		}
	}
	return next(ctx, tx, simulate)
}
