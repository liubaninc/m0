package recrypt

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/sm2"
	curve25519 "github.com/miguelsandro/curve25519-go/axlsign"
	"testing"
)

func TestRecrypt(t *testing.T) {
	proxy := sm2.GenPrivKey()
	alice := sm2.GenPrivKey()
	bob := sm2.GenPrivKey()

	proxy_re := curve25519.GenerateKeyPair(proxy.Key)
	alice_re := curve25519.GenerateKeyPair(alice.Key)
	bob_re := curve25519.GenerateKeyPair(bob.Key)

	m := "Hello, Proxy Re-Encryption"
	fmt.Println("origin message:", m)

	// 加密 alice公钥 + proxy私钥
	chiperContent, envelope, err := encrypt(proxy_re.PrivateKey, alice_re.PublicKey, m)
	if err != nil {
		t.Errorf("Encrypt %s", err)
	}
	fmt.Printf("alice encrypt: chiper %x envelope %x\n", chiperContent, envelope)

	// 解密 alice私钥 + proxy公钥
	content, err := dencrypt(alice_re.PrivateKey, proxy_re.PublicKey, envelope, chiperContent)
	if err != nil {
		t.Errorf("Dencrypt %s", err)
	}
	fmt.Println("alice dencrypt:", string(content))

	// 代理重加密 alice公钥 + bob公钥 + proxy私钥
	reEnvelope, err := reEnvelope(proxy_re, alice_re.PublicKey, bob_re.PublicKey, envelope)
	if err != nil {
		t.Errorf("ReEnvelope %s", err)
	}
	fmt.Printf("bob reEnvelope: %x\n", reEnvelope)

	// bob 私钥解密 + proxy公钥
	content, err = dencrypt(bob_re.PrivateKey, proxy_re.PublicKey, reEnvelope, chiperContent)
	if err != nil {
		t.Errorf("Dencrypt %s", err)
	}
	fmt.Println("bob dencrypt:", string(content))

}

func TestRecrypt2(t *testing.T) {
	alice := sm2.GenPrivKey()
	bob := sm2.GenPrivKey()
	carol := sm2.GenPrivKey()

	alice_pub := curve25519.GenerateKeyPair(alice.Key).PublicKey
	bob_pub := curve25519.GenerateKeyPair(bob.Key).PublicKey
	carol_pub := curve25519.GenerateKeyPair(carol.Key).PublicKey

	m := "Hello, Proxy Re-Encryption"
	fmt.Println("origin message:", m)

	// alice 加密存储
	chiperContent, envelope, err := Encrypt(alice, alice_pub, m)
	if err != nil {
		t.Errorf("Encrypt %s", err)
	}
	fmt.Printf("alice encrypt: chiper %x envelope %x\n", chiperContent, envelope)

	// alice解密
	content, err := Dencrypt(alice, alice_pub, envelope, chiperContent)
	if err != nil {
		t.Errorf("Dencrypt %s", err)
	}
	fmt.Println("alice dencrypt:", string(content))

	// alice 分享给 bob
	reEnvelope, err := ReEnvelope(alice, alice_pub, bob_pub, envelope)
	if err != nil {
		t.Errorf("ReEnvelope %s", err)
	}
	fmt.Printf("bob reEnvelope: %x\n", reEnvelope)

	// bob 解密
	content, err = Dencrypt(bob, alice_pub, reEnvelope, chiperContent)
	if err != nil {
		t.Errorf("Dencrypt %s", err)
	}
	fmt.Println("bob dencrypt:", string(content))

	// bob 分享给 carol
	rereEnvelope, err := ReEnvelope(bob, alice_pub, carol_pub, reEnvelope)
	if err != nil {
		t.Errorf("ReEnvelope %s", err)
	}
	fmt.Printf("carol reEnvelope: %x\n", reEnvelope)

	// carol 解密
	content, err = Dencrypt(carol, bob_pub, rereEnvelope, chiperContent)
	if err != nil {
		t.Errorf("Dencrypt %s", err)
	}
	fmt.Println("carol dencrypt:", string(content))
}
