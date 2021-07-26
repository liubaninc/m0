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
	certificates, found := k.GetCertificates(ctx, msg.Subject+"/"+msg.SubjectKeyID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "no X509 certificate associated with the "+
			"of subject=%v and subjectKeyID=%v already exists",
			msg.Subject, msg.SubjectKeyID)
	}

	if msg.Creator != certificates.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Only owner can revoke certificate")
	}

	certificate, found := k.GetCertificate(ctx, certificates.Items[0])
	if !found {
		panic("certificate not found")
	}

	if certificate.IsRoot {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "certificate with subject=%v and subjectKeyID=%v "+
			"is a root certificate.", msg.Subject, msg.SubjectKeyID)
	}

	revokedCertificates, _ := k.GetRevokedCertificates(ctx, msg.Subject+"/"+msg.SubjectKeyID)
	revokedCertificates.Creator = msg.Creator
	revokedCertificates.Items = append(revokedCertificates.Items, certificates.Items...)

	k.SetRevokedCertificates(ctx, revokedCertificates)
	k.RemoveCertificates(ctx, msg.Subject+"/"+msg.SubjectKeyID)

	issuer := certificate.Issuer
	authorityKeyID := certificate.AuthorityKeyID

	revokeChildCertificates(ctx, k.Keeper, msg.Subject+"/"+msg.SubjectKeyID)
	k.removeChildCertificateEntry(ctx, issuer, authorityKeyID, types.CertificateIdentifier{
		Subject:      msg.Subject,
		SubjectKeyID: msg.SubjectKeyID,
	})

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgRevokeCertResponse{}, nil
}

func revokeChildCertificates(ctx sdk.Context, keeper Keeper, index string) {
	// Get issuer's ChildCertificates record
	childCertificates, _ := keeper.GetChildCertificates(ctx, index)

	// For each child certificate subject/subjectKeyID combination
	for _, certIdentifier := range childCertificates.Items {
		// Revoke certificates with this subject/subjectKeyID combination
		certificates, _ := keeper.GetRevokedCertificates(ctx, certIdentifier)
		revokedCertificates, _ := keeper.GetRevokedCertificates(ctx, certIdentifier)
		revokedCertificates.Items = append(revokedCertificates.Items, certificates.Items...)
		keeper.SetRevokedCertificates(ctx, revokedCertificates)
		keeper.RemoveCertificates(ctx, certIdentifier)

		// Process child certificates recursively
		revokeChildCertificates(ctx, keeper, certIdentifier)
	}

	// Delete entire ChildCertificates record of issuer
	keeper.RemoveChildCertificates(ctx, index)
}
