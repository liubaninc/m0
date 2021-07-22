package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/permission/types"
)

func (k msgServer) SetPermission(goCtx context.Context, msg *types.MsgSetPermission) (*types.MsgSetPermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	k.SetAccount(ctx, types.Account{
		Creator: msg.Creator,
		Address: msg.Address,
		Perms:   msg.Perms,
	})

	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	if acc := k.accountKeeper.GetAccount(ctx, addr); acc == nil {
		defer telemetry.IncrCounter(1, "new", "account")
		k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, addr))
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetPermission,
			sdk.NewAttribute(types.AttributeKeyAddress, msg.Address),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgSetPermissionResponse{}, nil
}
