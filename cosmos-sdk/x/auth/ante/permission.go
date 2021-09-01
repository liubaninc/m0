package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type ValidatePermissonDecorator struct {
	keeper PermissionKeeper
}

func NewValidatePermissonDecorator(keeper PermissionKeeper) ValidatePermissonDecorator {
	return ValidatePermissonDecorator{keeper: keeper}
}

func (vad ValidatePermissonDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		route := msg.Route()
		for _, acc := range msg.GetSigners() {
			if !vad.keeper.HasPerm(ctx, acc.String(), route) {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s had no perm %s", acc.String(), route)
			}
		}
	}
	return next(ctx, tx, simulate)
}
