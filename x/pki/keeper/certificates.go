package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/pki/types"
)

// SetCertificates set a specific certificates in the store from its index
func (k Keeper) SetCertificates(ctx sdk.Context, certificates types.Certificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificatesKey))
	b := k.cdc.MustMarshalBinaryBare(&certificates)
	store.Set(types.KeyPrefix(certificates.Identifier.Index()), b)
}

// GetCertificates returns a certificates from its index
func (k Keeper) GetCertificates(ctx sdk.Context, subject string, subjectKeyID string) (val types.Certificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificatesKey))

	val.Identifier = types.CertificatesIdentifier{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}
	b := store.Get(types.KeyPrefix(val.Identifier.Index()))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteCertificates removes a certificates from the store
func (k Keeper) RemoveCertificates(ctx sdk.Context, subject string, subjectKeyID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificatesKey))
	store.Delete(types.KeyPrefix(types.CertificatesIdentifier{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}.Index()))
}

// GetAllCertificates returns all certificates
func (k Keeper) GetAllCertificates(ctx sdk.Context) (list []types.Certificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CertificatesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Certificates
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetCertificates set a specific certificates in the store from its index
func (k Keeper) SetRevokedCertificates(ctx sdk.Context, certificates types.Certificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKey))
	b := k.cdc.MustMarshalBinaryBare(&certificates)
	store.Set(types.KeyPrefix(certificates.Identifier.Index()), b)
}

// GetCertificates returns a certificates from its index
func (k Keeper) GetRevokedCertificates(ctx sdk.Context, subject string, subjectKeyID string) (val types.Certificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKey))

	val.Identifier = types.CertificatesIdentifier{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}
	b := store.Get(types.KeyPrefix(val.Identifier.Index()))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteCertificates removes a certificates from the store
func (k Keeper) RemoveRevokedCertificates(ctx sdk.Context, subject string, subjectKeyID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKey))
	store.Delete(types.KeyPrefix(types.CertificatesIdentifier{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}.Index()))
}

// GetAllCertificates returns all certificates
func (k Keeper) GetAllRevokedCertificates(ctx sdk.Context) (list []types.Certificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevokedCertificatesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Certificates
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetCertificates set a specific certificates in the store from its index
func (k Keeper) SetChildCertificates(ctx sdk.Context, certificates types.ChildCertificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKey))
	b := k.cdc.MustMarshalBinaryBare(&certificates)
	store.Set(types.KeyPrefix(certificates.Identifier.Index()), b)
}

// GetCertificates returns a certificates from its index
func (k Keeper) GetChildCertificates(ctx sdk.Context, subject string, subjectKeyID string) (val types.ChildCertificates, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKey))

	val.Identifier = types.CertificatesIdentifier{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}
	b := store.Get(types.KeyPrefix(val.Identifier.Index()))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// DeleteCertificates removes a certificates from the store
func (k Keeper) RemoveChildCertificates(ctx sdk.Context, subject string, subjectKeyID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKey))
	store.Delete(types.KeyPrefix(types.CertificatesIdentifier{
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}.Index()))
}

// GetAllCertificates returns all certificates
func (k Keeper) GetAllChildCertificates(ctx sdk.Context) (list []types.Certificates) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChildCertificatesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Certificates
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) addChildCertificateEntry(ctx sdk.Context, issuer string, authorityKeyID string, certIdentifier types.CertificatesIdentifier) {
	childCertificates, _ := k.GetChildCertificates(ctx, issuer, authorityKeyID)

	for _, existingIdentifier := range childCertificates.Items {
		if existingIdentifier.Index() == certIdentifier.Index() {
			return
		}
	}

	childCertificates.Items = append(childCertificates.Items, certIdentifier)
	k.SetChildCertificates(ctx, childCertificates)
}

func (k Keeper) removeChildCertificateEntry(ctx sdk.Context, issuer string, authorityKeyID string, certIdentifier types.CertificatesIdentifier) {
	childCertificates, _ := k.GetChildCertificates(ctx, issuer, authorityKeyID)

	certIDIndex := -1
	for i, existingIdentifier := range childCertificates.Items {
		if existingIdentifier.Index() == certIdentifier.Index() {
			certIDIndex = i

			break
		}
	}

	if certIDIndex == -1 {
		return
	}

	childCertificates.Items =
		append(childCertificates.Items[:certIDIndex], childCertificates.Items[certIDIndex+1:]...)

	if len(childCertificates.Items) > 0 {
		k.SetChildCertificates(ctx, childCertificates)
	} else {
		k.RemoveChildCertificates(ctx, issuer, authorityKeyID)
	}
}
