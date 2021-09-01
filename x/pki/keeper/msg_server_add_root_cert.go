package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/pki/x509"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/pki/types"
)

func (k msgServer) AddRootCert(goCtx context.Context, msg *types.MsgAddRootCert) (*types.MsgAddRootCertResponse, error) {
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
		// subjectKeyID. Since new certificate is self-signed, we have to ensure that the exisiting certificates are
		// self-signed too, consequently are root certificates.
		if !cert.IsRoot {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Issuer and authorityKeyID of new certificate with subject=%v and subjectKeyID=%v "+
				"must be the same as ones of existing certificates with the same subject and subjectKeyID",
				x509Certificate.Subject, x509Certificate.SubjectKeyID)
		}
	}

	// verify certificate
	_, _, err = k.verifyCertificate(ctx, x509Certificate)
	if err != nil {
		return nil, err
	}

	rootCertificate := types.NewRootCertificate(
		msg.Certificate,
		x509Certificate.Subject,
		x509Certificate.SubjectKeyID,
		x509Certificate.SerialNumber,
		x509Certificate.Issuer,
		msg.Creator,
	)
	identifier := types.CertificateIdentifier{
		Issuer:       x509Certificate.Issuer,
		SerialNumber: x509Certificate.SerialNumber,
	}
	certificates.Creator = msg.Creator
	certificates.Items = append(certificates.Items, identifier)

	k.SetCertificate(ctx, rootCertificate)
	k.SetCertificates(ctx, certificates)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, msg.Route()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(msg.Type(),
			sdk.NewAttribute(types.AttributeKeyCertificate, identifier.Index()),
		),
	})
	return &types.MsgAddRootCertResponse{}, nil
}
