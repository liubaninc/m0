package sdk

import (
	"testing"
)

func TestToken(t *testing.T) {
	res, err := testClient.GetToken("m0token")
	if err != nil {
		t.Fatal("token", err)
	}
	t.Log("token", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}

func TestIssue(t *testing.T) {
	res, err := testClient.BroadcastIssueTx(address, []string{address}, []string{"1000m0t"}, "test", "", "")
	if err != nil {
		t.Fatal("issue", "tx", err)
	}
	t.Log("issue", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}

func TestDestroy(t *testing.T) {
	res, err := testClient.BroadcastDestroyTx(address, "1000m0t", "test", "", "")
	if err != nil {
		t.Fatal("destroy", "tx", err)
	}
	t.Log("destroy", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}

func TestSend(t *testing.T) {
	res, err := testClient.BroadcastIssueTx(address, []string{address}, []string{"1000m0t"}, "test", "", "")
	if err != nil {
		t.Fatal("send", "tx", err)
	}
	t.Log("send", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}
