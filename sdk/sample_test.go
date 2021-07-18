package sdk

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestSample(t *testing.T) {
	mnemonic := "key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear"
	kr := keyring.NewInMemory()
	// 导入助记词
	info, err := kr.NewAccount("alice", mnemonic, "", sdk.GetConfig().GetFullFundraiserPath(), hd.Secp256k1)
	if err != nil {
		t.Fatal("TestSample", "new account", err)
	}
	address, _ := sdk.AccAddressFromBech32(info.GetAddress().String())

	sdkClient := MustNew("http://localhost:26657", kr)
	// 发行资产
	msg, err := sdkClient.IssueMsg(address.String(), []string{address.String()}, []string{"1000m0t"}, "test", "")
	if err != nil {
		t.Fatal("TestSample", "issue msg", err)
	}
	// 构建签名交易并发送
	res, err := sdkClient.GenerateAndBroadcastTx(address.String(), "", "ss", 0, msg)
	if err != nil {
		t.Fatal("TestSample", "broadcast tx", err)
	}
	if res.Code != 0 {
		t.Fatal("TestSample", "result", err)
	}
	t.Log("TestSample", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}

func TestSample2(t *testing.T) {
	mnemonic := "key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear"
	kr := keyring.NewInMemory()
	// 导入助记词
	info, err := kr.NewAccount("alice", mnemonic, "", sdk.GetConfig().GetFullFundraiserPath(), hd.Secp256k1)
	if err != nil {
		t.Fatal("TestSample2", "new account", err)
	}
	address, _ := sdk.AccAddressFromBech32(info.GetAddress().String())

	sdkClient := MustNew("http://localhost:26657", kr)
	// 发行交易
	msg, err := sdkClient.IssueMsg(address.String(), []string{address.String()}, []string{"1000m0t"}, "test", "")
	if err != nil {
		t.Fatal("TestSample2", "issue msg", err)
	}
	// 构建交易
	tx, err := sdkClient.GenerateTx(address.String(), "", "ss", 0, msg)
	if err != nil {
		t.Fatal("TestSample2", "build tx", err)
	}
	// 签名交易
	if err := sdkClient.SignTx(address.String(), "", tx, true); err != nil {
		t.Fatal("TestSample2", "sign tx", err)
	}
	// 广播交易
	res, err := sdkClient.BroadcastTx(tx.GetTx())
	if err != nil {
		t.Fatal("TestSample2", "broadcast tx", err)
	}
	if res.Code != 0 {
		t.Fatal("TestSample2", "result", err)
	}
	t.Log("TestSample2", string(sdkClient.JSONMarshaler.MustMarshalJSON(res)))
}
