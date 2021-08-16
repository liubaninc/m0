package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/peer/types"
	"github.com/liubaninc/m0/x/pki/x509"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (k Keeper) IDPeerFilter(ctx sdk.Context, index string) abci.ResponseQuery {
	if ctx.BlockHeight() == 0 || !k.GetParams(ctx).Enabled {
		return abci.ResponseQuery{
			Code: abci.CodeTypeOK,
		}
	}

	peerID, found := k.GetPeerID(ctx, index)
	if !found {
		return abci.ResponseQuery{
			Code: 1,
			Log:  fmt.Sprintf("node id %s not found in peer module", index),
		}
	}

	cert, found := k.pkiKeeper.GetCertificate(ctx, peerID.CertIssuer, peerID.CertSerialNum)
	if !found {
		return abci.ResponseQuery{
			Code: 1,
			Log:  fmt.Sprintf("certificate %s/%s not found in pki module", peerID.CertIssuer, peerID.CertSerialNum),
		}
	}

	if _, found := k.pkiKeeper.GetRevokedCertificates(ctx, cert.Subject, cert.SubjectKeyID); found {
		return abci.ResponseQuery{
			Code: 1,
			Log:  fmt.Sprintf("certificate %s/%s is revoked in pki module", peerID.CertIssuer, peerID.CertSerialNum),
		}
	}

	if certs, _ := k.pkiKeeper.GetCertificates(ctx, cert.Subject, cert.SubjectKeyID); certs.Disable {
		return abci.ResponseQuery{
			Code: 1,
			Log:  fmt.Sprintf("certificate %s/%s is frozen in pki module", peerID.CertIssuer, peerID.CertSerialNum),
		}
	}

	x509Cert, err := x509.DecodeX509Certificate(cert.PemCert)
	if err != nil {
		panic(err)
	}

	if x509Cert.Certificate.Subject.CommonName != index {
		return abci.ResponseQuery{
			Code: 1,
			Log:  fmt.Sprintf("invalid certificate, mismatch CommonName"),
		}
	}

	if ctx.BlockHeader().Time.After(x509Cert.Certificate.NotAfter) {
		return abci.ResponseQuery{
			Code: 1,
			Log:  fmt.Sprintf("certificate %s/%s has expired in pki module", peerID.CertIssuer, peerID.CertSerialNum),
		}
	}

	return abci.ResponseQuery{
		Code: abci.CodeTypeOK,
	}
}

// SetPeerID set a specific peerID in the store from its index
func (k Keeper) SetPeerID(ctx sdk.Context, peerID types.PeerID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeerIDKey))
	b := k.cdc.MustMarshalBinaryBare(&peerID)
	store.Set(types.KeyPrefix(peerID.Index), b)
}

// GetPeerID returns a peerID from its index
func (k Keeper) GetPeerID(ctx sdk.Context, index string) (val types.PeerID, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeerIDKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeletePeerID removes a peerID from the store
func (k Keeper) RemovePeerID(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeerIDKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllPeerID returns all peerID
func (k Keeper) GetAllPeerID(ctx sdk.Context) (list []types.PeerID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PeerIDKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PeerID
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
