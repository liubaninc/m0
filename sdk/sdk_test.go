package sdk

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"os"
	"testing"
)

var (
	kr, _  = keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, "~/.m0", os.Stdin)
	testClient = MustNew("http://localhost:26657", kr)
	mnemonic   = "key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear"
	address = "mc19dzfuxxv8vjeajjq475ahgrl0meudwexdmrnye"
)

func TestGenerateAndBroadcastTx(t *testing.T) {
	msg, err := testClient.IssueMsg(address, address, "1000m0t", "test", "")
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", "msg", err)
	}
	res, err := testClient.GenerateAndBroadcastTx(address, "", "ss", 0, msg)
	if err != nil {
		t.Fatal("GenerateAndBroadcastTx", "tx", err)
	}
	if res.Code != 0 {
		t.Fatal("GenerateAndBroadcastTx", "tx", err)
	}
	t.Log("GenerateAndBroadcastTx", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}