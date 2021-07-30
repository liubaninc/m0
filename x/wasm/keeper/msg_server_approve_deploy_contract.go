package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/wasm/types"
)

func (k msgServer) ApproveDeployContract(goCtx context.Context, msg *types.MsgApproveDeployContract) (*types.MsgApproveDeployContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposeDeploy, found := k.GetProposeDeploy(ctx, msg.Index)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "index is not proposing")
	}

	isProposer := false
	for _, proposer := range proposeDeploy.Approval {
		if proposer == msg.Creator {
			isProposer = true
			break
		}
	}

	if !isProposer {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid proposer")
	}

	for _, proposer := range proposeDeploy.Approved {
		if proposer == msg.Creator {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "already approved")
		}
	}

	proposeDeploy.Approved = append(proposeDeploy.Approved, msg.Creator)

	if len(proposeDeploy.Approval) == len(proposeDeploy.Approved) {
		k.RemoveProposeDeploy(ctx, proposeDeploy.Index)
		k.SetApproveDeploy(ctx, types.ApproveDeploy{
			Creator:      proposeDeploy.Creator,
			Index:        proposeDeploy.Index,
			ContractName: proposeDeploy.ContractName,
			ContractCode: proposeDeploy.ContractCode,
			InitArgs:     proposeDeploy.InitArgs,
			Approval:     proposeDeploy.Approval,
		})
	} else {
		k.SetProposeDeploy(ctx, proposeDeploy)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			msg.Type(),
			sdk.NewAttribute(types.AttributeIndexName, proposeDeploy.Index),
		),
	})

	return &types.MsgApproveDeployContractResponse{}, nil
}
