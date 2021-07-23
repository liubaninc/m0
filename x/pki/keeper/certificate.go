package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/pki/types"
)

// SetCertificate set a specific certificate in the store from its index
func (k Keeper) SetCertificate(ctx sdk.Context, certificate types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateKey))
	b := k.cdc.MustMarshalBinaryBare(&certificate)
	store.Set(types.KeyPrefix(certificate.Index()), b)
}

// GetCertificate returns a certificate from its index
func (k Keeper) GetCertificate(ctx sdk.Context, index string) (val types.Certificate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteCertificate removes a certificate from the store
func (k Keeper) RemoveCertificate(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllCertificate returns all certificate
func (k Keeper) GetAllCertificate(ctx sdk.Context) (list []types.Certificate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificateKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Certificate
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
