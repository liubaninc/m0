package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/pki/x509"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/pki/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) verifyCertificate(ctx sdk.Context, x509Certificate *x509.X509Certificate) (string, string, error) {
	if x509Certificate.IsSelfSigned() {
		// in this system a certificate is self-signed if and only if it is a root certificate
		if err := x509Certificate.Verify(x509Certificate); err == nil {
			return x509Certificate.Subject, x509Certificate.SubjectKeyID, nil
		}
	} else {
		//parentCertificates := keeper.GetApprovedCertificates(ctx, x509Certificate.Issuer, x509Certificate.AuthorityKeyID)
		//
		//for _, cert := range parentCertificates.Items {
		//	parentX509Certificate, err := x509.DecodeX509Certificate(cert.PemCert)
		//	if err != nil {
		//		continue
		//	}
		//
		//	// verify certificate against parent
		//	if err := x509Certificate.Verify(parentX509Certificate); err != nil {
		//		continue
		//	}
		//
		//	// verify parent certificate
		//	if subject, subjectKeyID, err := verifyCertificate(ctx, keeper, parentX509Certificate); err == nil {
		//		return subject, subjectKeyID, nil
		//	}
		//}
	}

	return "", "", sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Certificate verification failed for certificate with subject=%v and subjectKeyID=%v",
		x509Certificate.Subject, x509Certificate.SubjectKeyID)
}
