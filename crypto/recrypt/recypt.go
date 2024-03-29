package recrypt

import (
	"encoding/hex"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	ed25519 "github.com/miguelsandro/curve25519-go/axlsign"
	"math/rand"
	"time"
)

func generateEncryptKey() []byte {
	rand.Seed(time.Now().UnixNano())
	var seed = make([]uint8, 32)
	for i := 0; i < len(seed); i++ {
		seed[i] = uint8(rand.Int() % (256))
	}
	return seed
}

// 加密
func encrypt(privateKey, publicKey []byte, content string) (string, string, error) {
	//随机加密密码加密内容
	contentKey := generateEncryptKey()
	cipherContent, err := AESCBCPKCS7Encrypt(contentKey, []byte(content))
	if err != nil {
		return "", "", err
	}

	//生成加密密码信封
	sk := ed25519.SharedKey(privateKey, publicKey)
	envelope, err := AESCBCPKCS7Encrypt(sk, contentKey)
	if err != nil {
		return "", "", err
	}
	return hex.EncodeToString(cipherContent), hex.EncodeToString(envelope), nil
}

// 解密
func dencrypt(privateKey, publicKey []byte, envelope string, cipherContent string) (string, error) {
	//拆开加密密码信封
	sk := ed25519.SharedKey(privateKey, publicKey)
	envelopeBytes, err := hex.DecodeString(envelope)
	if err != nil {
		return "", err
	}
	contentKey, err := AESCBCPKCS7Decrypt(sk, envelopeBytes)
	if err != nil {
		return "", err
	}
	//加密密码解密内容
	cipherContentBytes, err := hex.DecodeString(cipherContent)
	if err != nil {
		return "", err
	}
	content, err := AESCBCPKCS7Decrypt(contentKey, cipherContentBytes)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

//  代理重加密
func reEnvelope(priv ed25519.Keys, oldPublicKey, newPublicKey []byte, envelope string) (string, error) {
	//拆开加密密码信封
	sk := ed25519.SharedKey(priv.PrivateKey, oldPublicKey)
	envelopeBytes, err := hex.DecodeString(envelope)
	if err != nil {
		return "", err
	}
	contentKey, err := AESCBCPKCS7Decrypt(sk, envelopeBytes)
	if err != nil {
		return "", err
	}

	//重新生成加密密码信封
	sk = ed25519.SharedKey(priv.PrivateKey, newPublicKey)
	reEnvelope, err := AESCBCPKCS7Encrypt(sk, contentKey)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(reEnvelope), nil
}

func Encrypt(key cryptotypes.PrivKey, publicKey []byte, content string) (string, string, error) {
	priv := ed25519.GenerateKeyPair(key.Bytes())
	return encrypt(priv.PrivateKey, publicKey, content)
}

func Dencrypt(key cryptotypes.PrivKey, publicKey []byte, envelope string, cipherContent string) (string, error) {
	priv := ed25519.GenerateKeyPair(key.Bytes())
	return dencrypt(priv.PrivateKey, publicKey, envelope, cipherContent)
}

func ReEnvelope(key cryptotypes.PrivKey, oldPublicKey, newPublicKey []byte, envelope string) (string, error) {
	priv := ed25519.GenerateKeyPair(key.Bytes())
	return reEnvelope(priv, oldPublicKey, newPublicKey, envelope)
}

func GetPubKey(key cryptotypes.PrivKey) []byte {
	priv := ed25519.GenerateKeyPair(key.Bytes())
	return priv.PublicKey
}
