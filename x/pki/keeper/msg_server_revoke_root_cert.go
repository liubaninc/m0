package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/pki/types"
)

func (k msgServer) RevokeRootCert(goCtx context.Context, msg *types.MsgRevokeRootCert) (*types.MsgRevokeRootCertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get list of certificates for Subject / Subject Key Id combination
	certificates, found := k.GetCertificates(ctx, msg.Subject+"/"+msg.SubjectKeyID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "no X509 certificate associated with the " +
			"of subject=%v and subjectKeyID=%v already exists",
			msg.Subject, msg.SubjectKeyID)
	}

	if msg.Creator != certificates.Creator {
		return nil,  sdkerrors.Wrapf(sdkerrors.ErrUnauthorized,"Only owner can revoke root certificate")
	}

	certificate, found := k.GetCertificate(ctx, certificates.Items[0])
	if !found {
		panic("certificate not found")
	}

	if !certificate.IsRoot {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "certificate with subject=%v and subjectKeyID=%v "+
				"is not a root certificate.", msg.Subject, msg.SubjectKeyID)
	}

	revokedCertificates, _ := k.GetRevokedCertificates(ctx, msg.Subject+"/"+msg.SubjectKeyID)
	revokedCertificates.Creator = msg.Creator
	revokedCertificates.Items = append(revokedCertificates.Items, certificates.Items...)

	k.SetRevokedCertificates(ctx, revokedCertificates)
	k.RemoveCertificates(ctx, msg.Subject+"/"+msg.SubjectKeyID)

	revokeChildCertificates(ctx, k.Keeper,  msg.Subject + "/" + msg.SubjectKeyID)


	return &types.MsgRevokeRootCertResponse{}, nil
}
