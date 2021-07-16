package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/authority/types"
)

func (k msgServer) ProposeRevokeAccount(goCtx context.Context, msg *types.MsgProposeRevokeAccountRequest) (*types.MsgProposeRevokeAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if sender has enough rights to propose account.
	if !k.Keeper.HasRole(ctx, msg.Creator, types.AUTHORITY) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "MsgProposeRevokeAccount transaction should be signed by an account with the %s role", types.AUTHORITY)

	}
	// check that account exists
	if !k.Keeper.IsAccountPresent(ctx, msg.Address) {
		return nil, sdkerrors.Wrapf(types.ErrAccountAlreadyExists, "No account associated with the address=%s on the ledger", msg.Address)
	}
	// check that pending account revocation does not exist yet
	if k.Keeper.IsPendingAccountRevocationPresent(ctx, msg.Address) {
		return nil, sdkerrors.Wrapf(types.ErrAccountAlreadyExists, "Pending account revocation associated with the address=%s already exists on the ledger", msg.Address)
	}
	// if more than 1 trustee's approval is needed, create pending account revocation else delete the account.
	if AccountApprovalsCount(ctx, k.Keeper) > 1 {
		// create and store pending account revocation record
		revoc := types.PendingAccountRevocation{
			Address:   msg.Address,
			Approvals: []string{msg.Creator},
		}
		k.Keeper.SetPendingAccountRevocation(ctx, revoc)
	} else {

		acct := k.accountKeeper.NewAccount(ctx, &authtypes.BaseAccount{
			Address: msg.Address,
		})
		k.accountKeeper.RemoveAccount(ctx, acct)

		// delete account record
		k.Keeper.DeleteAccount(ctx, msg.Address)
	}
	return &types.MsgProposeRevokeAccountResponse{}, nil
}
