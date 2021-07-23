package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/validator/types"
	tmstrings "github.com/tendermint/tendermint/libs/strings"
)

func (k msgServer) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	maxValidator := k.MaxValidators(ctx)
	if uint32(len(k.GetAllValidator(ctx))) == maxValidator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "the maximum number of validators has been reached")
	}

	if cp := ctx.ConsensusParams(); cp != nil {
		pk := sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, msg.PubKey)
		if !tmstrings.StringInSlice(pk.Type(), ctx.ConsensusParams().Validator.PubKeyTypes) {
			return nil, sdkerrors.Wrapf(
				sdkerrors.ErrUnknownRequest,
				"public key type %s is not supported. Supported types: [%s]", pk.Type(), ctx.ConsensusParams().Validator.PubKeyTypes,
			)
		}
	}

	// check if a validator with a given address already exists
	if k.HasValidator(ctx, msg.Creator) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "address %s is already a validator", msg.Creator)
	}

	validator := types.Validator{
		Creator:     msg.Creator,
		PubKey:      msg.PubKey,
		Power:       1,
		Description: msg.Description,
	}

	// check if a validator with a given address already exists
	if k.HasValidatorByConsAddr(ctx, validator.GetConsAddr().String()) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "public key %s is already a validator", msg.PubKey)
	}

	k.SetValidator(ctx, validator)
	k.SetValidatorByConsAddr(ctx, validator)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, validator.GetConsAddr().String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgCreateValidatorResponse{}, nil
}
