package keeper

import (
	"context"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/liubaninc/m0/x/authority/types"
)

func (k msgServer) ProposeAddAccount(goCtx context.Context, msg *types.MsgProposeAddAccountRequest) (*types.MsgProposeAddAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if sender has enough rights to propose account.
	if ctx.BlockHeight() != 0 && !k.Keeper.HasRole(ctx, msg.Creator, types.AUTHORITY) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "MsgProposeAddAccount transaction should be signed by an account with the %s role", types.AUTHORITY)

	}

	// check if active account already exists.
	if k.Keeper.IsAccountPresent(ctx, msg.Address) {
		return nil, sdkerrors.Wrapf(types.ErrAccountAlreadyExists, "Account associated with the address=%s already exists on the ledger", msg.Address)
	}

	// check if pending account already exists.
	if k.Keeper.IsPendingAccountPresent(ctx, msg.Address) {
		return nil, sdkerrors.Wrapf(types.ErrAccountAlreadyExists, "Pending account associated with the address=%v already exists on the ledger", msg.Address)
	}

	any, err := codectypes.NewAnyWithValue(sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, msg.PublicKey))
	if err != nil {
		panic(err)
	}

	// if more than 1 trustee's approval is needed, create pending account else create an active account.
	if AccountApprovalsCount(ctx, k.Keeper) > 1 {
		// create and store pending account.
		account := types.PendingAccount{
			Address:   msg.Address,
			PubKey:    any,
			Roles:     msg.Roles,
			Approvals: []string{msg.Creator},
		}
		k.Keeper.SetPendingAccount(ctx, account)
	} else {
		// create account, assign account number and store it
		account := types.Account{
			Address: msg.Address,
			PubKey:  any,
			Roles:   msg.Roles,
		}
		acct := k.accountKeeper.NewAccount(ctx, &authtypes.BaseAccount{
			Address: account.Address,
			PubKey:  account.PubKey,
		})
		k.accountKeeper.SetAccount(ctx, acct)

		account.AccountNumber = acct.GetAccountNumber()
		k.Keeper.SetAccount(ctx, account)
	}
	return &types.MsgProposeAddAccountResponse{}, nil
}