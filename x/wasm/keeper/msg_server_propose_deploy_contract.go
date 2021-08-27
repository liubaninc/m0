package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func (k msgServer) ProposeDeployContract(goCtx context.Context, msg *types.MsgProposeDeployContract) (*types.MsgProposeDeployContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hash := tmhash.Sum(append(msg.ContractCodeHash, append([]byte(msg.ContractName), []byte(msg.ContractName)...)...))
	index := fmt.Sprintf("%X", hash)
	if _, found := k.GetProposeDeploy(ctx, index); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "index is proposing")
	}

	if _, found := k.GetApproveDeploy(ctx, index); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "index already approved")
	}

	proposeDeploy := types.ProposeDeploy{
		Creator:          msg.Creator,
		Index:            index,
		ContractName:     msg.ContractName,
		ContractCodeHash: msg.ContractCodeHash,
		InitArgs:         msg.InitArgs,
		Approval:         msg.Approval,
	}

	attrs := make([]sdk.Attribute, len(msg.Approval))
	for i, proposer := range msg.Approval {
		if proposer == msg.Creator {
			proposeDeploy.Approved = append(proposeDeploy.Approved, msg.Creator)
		}
		attrs[i] = sdk.NewAttribute(types.AttributeProposer, proposer)
	}
	attrs = append(attrs, sdk.NewAttribute(types.AttributeProposeID, proposeDeploy.Index))

	k.SetProposeDeploy(ctx, proposeDeploy)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			msg.Type(),
			attrs...,
		),
	})

	return &types.MsgProposeDeployContractResponse{}, nil
}
