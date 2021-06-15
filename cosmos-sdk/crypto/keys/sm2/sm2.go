package sm2

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"io"
	"math/big"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

//-------------------------------------

const (
	PrivKeyName = "tendermint/PrivKeySm2"
	PubKeyName  = "tendermint/PubKeySm2"
	// PubKeySize is is the size, in bytes, of public keys as used in this package.
	PubKeySize = 33
	// PrivKeySize is the size, in bytes, of private keys as used in this package.
	PrivKeySize = 32
	// Size of an Edwards25519 signature. Namely the size of a compressed
	// Edwards25519 point, and a field element. Both of which are 32 bytes.
	SignatureSize = 64
	// SeedSize is the size, in bytes, of private key seeds. These are the
	// private key representations used by RFC 8032.
	SeedSize = 32

	keyType = "sm2"
)

var _ cryptotypes.PrivKey = &PrivKey{}
var _ codec.AminoMarshaler = &PrivKey{}

func (privKey *PrivKey) ToSM2PrivateKey() *sm2.PrivateKey {
	k := new(big.Int).SetBytes(privKey.Bytes())
	c := sm2.P256Sm2()
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return priv
}

// Bytes returns the privkey byte format.
func (privKey *PrivKey) Bytes() []byte {
	return privKey.Key
}

// Sign produces a signature on the provided message.
// This assumes the privkey is wellformed in the golang format.
// The first 32 bytes should be random,
// corresponding to the normal sm2 private key.
// The latter 32 bytes should be the compressed public key.
// If these conditions aren't met, Sign will panic or produce an
// incorrect signature.
func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	r, s, err := sm2.Sm2Sign(privKey.ToSM2PrivateKey(), msg, nil, crypto.CReader())
	if err != nil {
		return nil, err
	}
	R := r.Bytes()
	S := s.Bytes()
	sig := make([]byte, SignatureSize)
	copy(sig[32-len(R):32], R[:])
	copy(sig[64-len(S):64], S[:])
	return sig, nil
}

// PubKey gets the corresponding public key from the private key.
//
// Panics if the private key is not initialized.
func (privKey *PrivKey) PubKey() cryptotypes.PubKey {
	priv := privKey.ToSM2PrivateKey()
	compPubkey := sm2.Compress(&priv.PublicKey)

	pubkeyBytes := make([]byte, PubKeySize)
	copy(pubkeyBytes, compPubkey[:])
	return &PubKey{Key: pubkeyBytes}
}

// Equals - you probably don't need to use this.
// Runs in constant time based on length of the keys.
func (privKey *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	if privKey.Type() != other.Type() {
		return false
	}

	return subtle.ConstantTimeCompare(privKey.Bytes(), other.Bytes()) == 1
}

func (privKey *PrivKey) Type() string {
	return keyType
}

// MarshalAmino overrides Amino binary marshalling.
func (privKey PrivKey) MarshalAmino() ([]byte, error) {
	return privKey.Key, nil
}

// UnmarshalAmino overrides Amino binary marshalling.
func (privKey *PrivKey) UnmarshalAmino(bz []byte) error {
	if len(bz) != PrivKeySize {
		return fmt.Errorf("invalid privkey size")
	}
	privKey.Key = bz

	return nil
}

// MarshalAminoJSON overrides Amino JSON marshalling.
func (privKey PrivKey) MarshalAminoJSON() ([]byte, error) {
	// When we marshal to Amino JSON, we don't marshal the "key" field itself,
	// just its contents (i.e. the key bytes).
	return privKey.MarshalAmino()
}

// UnmarshalAminoJSON overrides Amino JSON marshalling.
func (privKey *PrivKey) UnmarshalAminoJSON(bz []byte) error {
	return privKey.UnmarshalAmino(bz)
}

// GenPrivKey generates a new sm2 private key.
// It uses OS randomness in conjunction with the current global random seed
// in tendermint/libs/common to generate the private key.
func GenPrivKey() *PrivKey {
	return genPrivKey(crypto.CReader())
}

// genPrivKey generates a new sm2 private key using the provided reader.
func genPrivKey(rand io.Reader) *PrivKey {
	seed := make([]byte, SeedSize)

	_, err := io.ReadFull(rand, seed)
	if err != nil {
		panic(err)
	}

	privKey, err := sm2.GenerateKey(rand)
	if err != nil {
		panic(err)
	}
	return &PrivKey{Key: privKey.D.Bytes()}
}

// GenPrivKeyFromSecret hashes the secret with SHA2, and uses
// that 32 byte output to create the private key.
// NOTE: secret should be the output of a KDF like bcrypt,
// if it's derived from user input.
func GenPrivKeyFromSecret(secret []byte) *PrivKey {
	one := new(big.Int).SetInt64(1)
	secHash := sha256.Sum256(secret)

	k := new(big.Int).SetBytes(secHash[:])
	n := new(big.Int).Sub(sm2.P256Sm2().Params().N, one)
	k.Mod(k, n)
	k.Add(k, one)

	return &PrivKey{Key: k.Bytes()[:]}
}

//-------------------------------------

var _ cryptotypes.PubKey = &PubKey{}
var _ codec.AminoMarshaler = &PubKey{}

// Address is the SHA256-20 of the raw pubkey bytes.
func (pubKey *PubKey) Address() crypto.Address {
	if len(pubKey.Key) != PubKeySize {
		panic("pubkey is incorrect size")
	}
	return crypto.Address(tmhash.SumTruncated(pubKey.Key))
}

// Bytes returns the PubKey byte format.
func (pubKey *PubKey) Bytes() []byte {
	return pubKey.Key
}

func (pubKey *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	// make sure we use the same algorithm to sign
	if len(sig) != SignatureSize {
		return false
	}

	publicKey := sm2.Decompress(pubKey.Key[:])
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])

	return sm2.Sm2Verify(publicKey, msg, nil, r, s)
}

func (pubKey *PubKey) String() string {
	return fmt.Sprintf("PubKeySM2{%X}", pubKey.Key)
}

func (pubKey *PubKey) Type() string {
	return keyType
}

func (pubKey *PubKey) Equals(other cryptotypes.PubKey) bool {
	if pubKey.Type() != other.Type() {
		return false
	}

	return subtle.ConstantTimeCompare(pubKey.Bytes(), other.Bytes()) == 1
}

// MarshalAmino overrides Amino binary marshalling.
func (pubKey PubKey) MarshalAmino() ([]byte, error) {
	return pubKey.Key, nil
}

// UnmarshalAmino overrides Amino binary marshalling.
func (pubKey *PubKey) UnmarshalAmino(bz []byte) error {
	if len(bz) != PubKeySize {
		return errors.Wrap(errors.ErrInvalidPubKey, "invalid pubkey size")
	}
	pubKey.Key = bz

	return nil
}

// MarshalAminoJSON overrides Amino JSON marshalling.
func (pubKey PubKey) MarshalAminoJSON() ([]byte, error) {
	// When we marshal to Amino JSON, we don't marshal the "key" field itself,
	// just its contents (i.e. the key bytes).
	return pubKey.MarshalAmino()
}

// UnmarshalAminoJSON overrides Amino JSON marshalling.
func (pubKey *PubKey) UnmarshalAminoJSON(bz []byte) error {
	return pubKey.UnmarshalAmino(bz)
}
