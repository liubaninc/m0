package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/pki/types"
)

func (k msgServer) FreezeCert(goCtx context.Context, msg *types.MsgFreezeCert) (*types.MsgFreezeCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get list of certificates for Subject / Subject Key Id combination
	certificates, found := k.GetCertificates(ctx, msg.Subject, msg.SubjectKeyID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "no X509 certificate associated with the "+
			"of subject=%v and subjectKeyID=%v already exists",
			msg.Subject, msg.SubjectKeyID)
	}

	if msg.Creator != certificates.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only owner can revoke certificate")
	}

	certificate, found := k.GetCertificate(ctx, certificates.Items[0].Issuer, certificates.Items[0].SerialNumber)
	if !found {
		panic("certificate not found")
	}

	if certificate.IsRoot {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "certificate with subject=%v and subjectKeyID=%v "+
			"is a root certificate.", msg.Subject, msg.SubjectKeyID)
	}

	certificates.Disable = true

	k.SetCertificates(ctx, certificates)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(msg.Type(),
			sdk.NewAttribute(types.AttributeKeyCertificate, certificates.Identifier.Index()),
		),
	})

	return &types.MsgFreezeCertResponse{}, nil
}
