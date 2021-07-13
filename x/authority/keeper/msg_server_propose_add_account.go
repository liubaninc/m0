package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/authority/types"
	"math"
)

func (k msgServer) ProposeAddAccount(goCtx context.Context, msg *types.MsgProposeAddAccountRequest) (*types.MsgProposeAddAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	address, _ := sdk.AccAddressFromBech32(msg.Address)

	// check if sender has enough rights to propose account.
	if !k.Keeper.HasRole(ctx, address, types.AUTHORITY) {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "MsgProposeAddAccount transaction should be signed by an account with the %s role", types.AUTHORITY)

	}

	// check if active account already exists.
	if k.Keeper.IsAccountPresent(ctx, address) {
		return nil, sdkerrors.Wrapf(types.ErrAccountAlreadyExists, "Account associated with the address=%v already exists on the ledger", address)
	}

	// check if pending account already exists.
	if k.Keeper.IsPendingAccountPresent(ctx, address) {
		return nil, sdkerrors.Wrapf(types.ErrAccountAlreadyExists, "Pending account associated with the address=%v already exists on the ledger", address)
	}

	/*	// parse the key.
		pubKey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, msg.PublicKey)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, err.Error())
		}
	*/
	// if more than 1 trustee's approval is needed, create pending account else create an active account.
	if AccountApprovalsCount(ctx, k.Keeper) > 1 {
		// create and store pending account.
		account := types.PendingAccount{
			Address:   msg.Address,
			PubKey:    msg.PublicKey,
			Roles:     msg.Roles,
			Approvals: []string{msg.Creator},
		}
		k.Keeper.SetPendingAccount(ctx, account)
	} else {
		// create account, assign account number and store it
		account := types.Account{
			Address: msg.Address,
			PubKey:  msg.PublicKey,
			Roles:   msg.Roles,
		}
		account.AccountNumber = k.Keeper.GetNextAccountNumber(ctx).Value
		k.Keeper.SetAccount(ctx, account)
	}
	return &types.MsgProposeAddAccountResponse{}, nil
}

func AccountApprovalsCount(ctx sdk.Context, keeper Keeper) int {
	return int(math.Round(types.AccountApprovalPercent * float64(keeper.CountAccountsWithRole(ctx, types.AUTHORITY))))
}
