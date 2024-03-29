package sdk

import (
	"testing"
)

func TestDeploy(t *testing.T) {
	res, err := testClient.BroadcastDeployTx(address, "ccc", "../x/wasm/xmodel/contractsdk/cpp/build/counter.wasm", "{\"creator\":\"someone\"}", "test", "", "")
	if err != nil {
		t.Fatal("deploy", "tx", err)
	}
	t.Log("deploy", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}

func TestUpgrade(t *testing.T) {
	res, err := testClient.BroadcastUpgradeTx(address, "ccc", "../x/wasm/xmodel/contractsdk/cpp/build/counter.wasm", "test", "", "")
	if err != nil {
		t.Fatal("upgrade", "tx", err)
	}
	t.Log("upgrade", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}

func TestInvoke(t *testing.T) {
	res, err := testClient.BroadcastInvokeTx(address, "ccc", "increase", "{\"key\":\"someone\"}", "", "test", "", "")
	if err != nil {
		t.Fatal("invoke", "tx", err)
	}
	t.Log("invoke", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))

	resc, err := testClient.Query("ccc", "get", "{\"key\":\"someone\"}")
	if err != nil {
		t.Fatal("query", "tx", err)
	}
	t.Log("invoke", "result", string(testClient.JSONMarshaler.MustMarshalJSON(resc)))

}
