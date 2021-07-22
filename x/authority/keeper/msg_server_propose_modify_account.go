package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/authority/types"
)

func (k msgServer) ProposeModifyAccount(goCtx context.Context, msg *types.MsgProposeModifyAccountRequest) (*types.MsgProposeModifyAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if sender has enough rights to propose account.
	if !k.Keeper.HasRole(ctx, msg.Creator, types.AUTHORITY) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "MsgProposeModifyAccount transaction should be signed by an account with the %s role", types.AUTHORITY)

	}

	// check that account exists
	if !k.Keeper.IsAccountPresent(ctx, msg.Address) {
		return nil, sdkerrors.Wrapf(types.ErrAccountAlreadyExists, "No account associated with the address=%s on the ledger", msg.Address)
	}

	// check if pending account already exists.
	if k.Keeper.IsPendingAccountPresent(ctx, msg.Address) {
		return nil, sdkerrors.Wrapf(types.ErrAccountAlreadyExists, "Pending account associated with the address=%v already exists on the ledger", msg.Address)
	}

	// get  account
	account := k.Keeper.GetAccount(ctx, msg.Address)

	//Check if the roles are the same
	if len(account.Roles) == len(msg.Roles) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Account roles: [%v] already exists on the ledger", msg.Address, account.Roles)

	}

	// modify rolesØ
	account.Roles = msg.Roles

	// if more than 1 trustee's approval is needed, create pending account else create an active account.
	if AccountApprovalsCount(ctx, k.Keeper) > 1 {
		// create and store pending account.
		account := types.PendingAccount{
			Address:   msg.Address,
			PubKey:    account.PubKey,
			Roles:     msg.Roles,
			Approvals: []string{msg.Creator},
		}
		k.Keeper.SetPendingAccount(ctx, account)
	} else {
		// create account, assign account number and store it
		account := types.Account{
			Address: msg.Address,
			PubKey:  account.PubKey,
			Roles:   msg.Roles,
		}
		acct := k.accountKeeper.NewAccount(ctx, &authtypes.BaseAccount{
			Address: account.Address,
			PubKey:  account.PubKey,
		})
		//cover account record
		k.accountKeeper.SetAccount(ctx, acct)
		account.AccountNumber = acct.GetAccountNumber()
		k.Keeper.SetAccount(ctx, account)
	}

	return &types.MsgProposeModifyAccountResponse{}, nil
}
