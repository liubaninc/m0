package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/authority/types"
)

func (k msgServer) ApproveRevokeAccount(goCtx context.Context, msg *types.MsgApproveRevokeAccountRequest) (*types.MsgApproveRevokeAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that sender has enough rights to approve account revocation
	if !k.Keeper.HasRole(ctx, msg.Creator, types.AUTHORITY) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "MsgApproveRevokeAccount transaction should be signed by an account with the %s role", types.AUTHORITY)

	}
	if !k.Keeper.IsPendingAccountRevocationPresent(ctx, msg.Address) {
		return nil, sdkerrors.Wrapf(types.ErrAccountAlreadyExists, "No pending account revocation associated with the address=%s on the ledger", msg.Address)
	}
	// get pending account revocation
	revoc := k.Keeper.GetPendingAccountRevocation(ctx, msg.Address)

	// check if pending account revocation already has approval from signer
	if HasApprovalFrom(msg.Creator, revoc.Approvals) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "Pending account associated with the address=%s already has approval from=%s", msg.Address, msg.Creator)
	}
	// append approval
	revoc.Approvals = append(revoc.Approvals, msg.Creator)

	// check if pending account revocation has enough approvals
	if len(revoc.Approvals) == AccountApprovalsCount(ctx, k.Keeper) {
		// delete account record
		k.Keeper.DeleteAccount(ctx, msg.Address)
		// delete pending account revocation record
		k.Keeper.DeletePendingAccountRevocation(ctx, msg.Address)
	} else {
		// update pending account revocation record
		k.Keeper.SetPendingAccountRevocation(ctx, revoc)
	}
	return &types.MsgApproveRevokeAccountResponse{}, nil
}
