package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/pki/types"
)

func (k msgServer) RevokeCert(goCtx context.Context, msg *types.MsgRevokeCert) (*types.MsgRevokeCertResponse, error) {
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

	revokedCertificates, _ := k.GetRevokedCertificates(ctx, msg.Subject, msg.SubjectKeyID)
	revokedCertificates.Creator = msg.Creator
	revokedCertificates.Items = append(revokedCertificates.Items, certificates.Items...)

	k.SetRevokedCertificates(ctx, revokedCertificates)
	k.RemoveCertificates(ctx, msg.Subject, msg.SubjectKeyID)

	issuer := certificate.Issuer
	authorityKeyID := certificate.AuthorityKeyID

	k.removeChildCertificateEntry(ctx, issuer, authorityKeyID, certificates.Identifier)

	revokeChildCertificates(ctx, k.Keeper, msg.Subject, msg.SubjectKeyID)

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

	return &types.MsgRevokeCertResponse{}, nil
}

func revokeChildCertificates(ctx sdk.Context, keeper Keeper, subject string, subjectKeyID string) {
	// Get issuer's ChildCertificates record
	childCertificates, _ := keeper.GetChildCertificates(ctx, subject, subjectKeyID)

	// For each child certificate subject/subjectKeyID combination
	for _, certIdentifier := range childCertificates.Items {
		// Revoke certificates with this subject/subjectKeyID combination
		certificates, _ := keeper.GetCertificates(ctx, subject, subjectKeyID)
		revokedCertificates, _ := keeper.GetRevokedCertificates(ctx, subject, subjectKeyID)
		revokedCertificates.Items = append(revokedCertificates.Items, certificates.Items...)
		keeper.SetRevokedCertificates(ctx, revokedCertificates)
		keeper.RemoveCertificates(ctx, subject, subjectKeyID)

		// Process child certificates recursively
		revokeChildCertificates(ctx, keeper, certIdentifier.Subject, certIdentifier.SubjectKeyID)
	}

	// Delete entire ChildCertificates record of issuer
	keeper.RemoveChildCertificates(ctx, subject, subjectKeyID)
}
