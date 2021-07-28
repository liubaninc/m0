package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/pki/x509"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/pki/types"
)

func (k msgServer) AddCert(goCtx context.Context, msg *types.MsgAddCert) (*types.MsgAddCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	x509Certificate, err := x509.DecodeX509Certificate(msg.Certificate)
	if err != nil {
		panic(err)
	}

	// check if certificate with Issuer/Serial Number combination already exists
	if _, found := k.GetCertificate(ctx, x509Certificate.Issuer, x509Certificate.SerialNumber); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "certificate associated with the combination of issuer=%v and serialNumber=%v already exists already exist", x509Certificate.Issuer, x509Certificate.SerialNumber)
	}

	// Get list of certificates for Subject / Subject Key Id combination
	certificates, _ := k.GetCertificates(ctx, x509Certificate.Subject, x509Certificate.SubjectKeyID)
	if len(certificates.Items) > 0 {
		if msg.Creator != certificates.Creator {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Only owner of existing certificates with subject=%v and subjectKeyID=%v "+
				"can add new certificate with the same subject and subjectKeyID",
				x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}

		cert, found := k.GetCertificate(ctx, certificates.Items[0].Issuer, certificates.Items[0].SerialNumber)
		if !found {
			panic("certificate not found")
		}
		// Issuer and authorityKeyID must be the same as ones of exisiting certificates with the same subject and
		// subjectKeyID. Since new certificate is not self-signed, we have to ensure that the exisiting certificates
		// are not self-signed too, consequently are non-root certificates, before to match issuer and authorityKeyID.
		if cert.IsRoot || x509Certificate.Issuer != cert.Issuer ||
			x509Certificate.AuthorityKeyID != cert.AuthorityKeyID {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Issuer and authorityKeyID of new certificate with subject=%v and subjectKeyID=%v "+
				"must be the same as ones of existing certificates with the same subject and subjectKeyID",
				x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}
	}

	// verify certificate
	rootCertificateSubject, rootCertificateSubjectKeyID, err := k.verifyCertificate(ctx, x509Certificate)
	if err != nil {
		return nil, err
	}

	certificate := types.NewNonRootCertificate(
		msg.Certificate,
		x509Certificate.Subject,
		x509Certificate.SubjectKeyID,
		x509Certificate.SerialNumber,
		x509Certificate.Issuer,
		x509Certificate.AuthorityKeyID,
		rootCertificateSubject,
		rootCertificateSubjectKeyID,
		msg.Creator,
	)
	identifier := types.CertificateIdentifier{
		Issuer:       certificate.Issuer,
		SerialNumber: certificate.SerialNumber,
	}
	certificates.Creator = msg.Creator
	certificates.Items = append(certificates.Items, identifier)

	k.SetCertificate(ctx, certificate)
	k.SetCertificates(ctx, certificates)

	k.addChildCertificateEntry(ctx, certificate.Issuer, certificate.AuthorityKeyID, certificates.Identifier)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(msg.Type(),
			sdk.NewAttribute(types.AttributeKeyCertificate, identifier.Index()),
		),
	})
	return &types.MsgAddCertResponse{}, nil
}
