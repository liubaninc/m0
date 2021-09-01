package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/peer/types"
)

func (k msgServer) CreatePeerID(goCtx context.Context, msg *types.MsgCreatePeerID) (*types.MsgCreatePeerIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.pkiKeeper.GetCertificate(ctx, msg.CertIssuer, msg.CertSerialNum); !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("certificate %s/%s not found in pki module", msg.CertIssuer, msg.CertSerialNum))
	}

	// Check if the value already exists
	_, isFound := k.GetPeerID(ctx, msg.Index)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("nodeID %v already set", msg.Index))
	}

	var peerID = types.PeerID{
		Index:         msg.Index,
		Creator:       msg.Creator,
		CertIssuer:    msg.CertIssuer,
		CertSerialNum: msg.CertSerialNum,
	}

	k.SetPeerID(
		ctx,
		peerID,
	)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			msg.Type(),
			sdk.NewAttribute(types.AttributeKeyPeer, msg.Index),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgCreatePeerIDResponse{}, nil
}

func (k msgServer) UpdatePeerID(goCtx context.Context, msg *types.MsgUpdatePeerID) (*types.MsgUpdatePeerIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPeerID(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("nodeID %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if _, found := k.pkiKeeper.GetCertificate(ctx, msg.CertIssuer, msg.CertSerialNum); !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("certificate %s/%s not found in pki module", msg.CertIssuer, msg.CertSerialNum))
	}

	var peerID = types.PeerID{
		Index:         msg.Index,
		Creator:       msg.Creator,
		CertIssuer:    msg.CertIssuer,
		CertSerialNum: msg.CertSerialNum,
	}

	k.SetPeerID(ctx, peerID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			msg.Type(),
			sdk.NewAttribute(types.AttributeKeyPeer, msg.Index),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgUpdatePeerIDResponse{}, nil
}

func (k msgServer) DeletePeerID(goCtx context.Context, msg *types.MsgDeletePeerID) (*types.MsgDeletePeerIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPeerID(ctx, msg.Index)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("nodeID %v not set", msg.Index))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePeerID(ctx, msg.Index)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			msg.Type(),
			sdk.NewAttribute(types.AttributeKeyPeer, msg.Index),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgDeletePeerIDResponse{}, nil
}
