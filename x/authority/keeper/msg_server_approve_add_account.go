package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/liubaninc/m0/x/authority/types"
)

func (k msgServer) ApproveAddAccount(goCtx context.Context, msg *types.MsgApproveAddAccountRequest) (*types.MsgApproveAddAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("+++++++++++++++" + msg.Creator + "++++++++++++++++++++")
	// check if sender has enough rights to approve account.
	if !k.Keeper.HasRole(ctx, msg.Creator, types.AUTHORITY) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "MsgApproveAddAccount transaction should be signed by an account with the %s role", types.AUTHORITY)

	}
	// check if pending account exists
	if !k.Keeper.IsPendingAccountPresent(ctx, msg.Address) {
		return nil, sdkerrors.Wrapf(types.ErrAccountAlreadyExists, "No pending account associated with the address=%s on the ledger", msg.Address)
	}

	// get pending account
	pendAcc := k.Keeper.GetPendingAccount(ctx, msg.Address)

	// check if pending account already has approval from signer
	if HasApprovalFrom(msg.Creator, pendAcc.Approvals) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "Pending account associated with the address=%s already has approval from=%s", msg.Address, msg.Creator)
	}
	// append approval
	pendAcc.Approvals = append(pendAcc.Approvals, msg.Creator)

	// check if pending account has enough approvals
	if len(pendAcc.Approvals) == AccountApprovalsCount(ctx, k.Keeper) {
		// create approved account, assign account number and store it
		account := types.Account{
			Address: pendAcc.Address,
			PubKey:  pendAcc.PubKey,
			Roles:   pendAcc.Roles,
		}
		acct := k.accountKeeper.NewAccount(ctx, &authtypes.BaseAccount{
			Address: account.Address,
			PubKey: account.PubKey,
		})
		account.AccountNumber = acct.GetAccountNumber()
		k.Keeper.SetAccount(ctx, account)

		// delete pending account record
		k.Keeper.DeletePendingAccount(ctx, msg.Address)
	} else {
		// update pending account record
		k.Keeper.SetPendingAccount(ctx, pendAcc)
	}

	return &types.MsgApproveAddAccountResponse{}, nil
}
