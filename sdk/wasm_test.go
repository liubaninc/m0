package sdk

import (
	"io/ioutil"
	"testing"
)

func TestDeploy(t *testing.T) {
	code, err := ioutil.ReadFile("../x/wasm/xmodel/contractsdk/cpp/build/counter.wasm")
	if err != nil {
		t.Fatal("invalid code file", code, err)
	}
	res, err := testClient.BroadcastDeployTx(address, "ccc", code, "{\"creator\":\"someone\"}", "test", "", "")
	if err != nil {
		t.Fatal("deploy", "tx", err)
	}
	t.Log("deploy", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}

func TestUpgrade(t *testing.T) {
	code, err := ioutil.ReadFile("../x/wasm/xmodel/contractsdk/cpp/build/counter.wasm")
	if err != nil {
		t.Fatal("invalid code file", code, err)
	}
	res, err := testClient.BroadcastUpgradeTx(address, "ccc", code, "test", "", "")
	if err != nil {
		t.Fatal("upgrade", "tx", err)
	}
	t.Log("upgrade", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}

func TestFreeze(t *testing.T) {
	res, err := testClient.BroadcastFreezeTx(address, "ccc", "", "")
	if err != nil {
		t.Fatal("freeze", "tx", err)
	}
	t.Log("freeze", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}

func TestUnfreeze(t *testing.T) {
	res, err := testClient.BroadcastUnfreezeTx(address, "ccc", "", "")
	if err != nil {
		t.Fatal("unfreeze", "tx", err)
	}
	t.Log("unfreeze", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
}

func TestUndeploy(t *testing.T) {
	res, err := testClient.BroadcastUndeployTx(address, "ccc", "", "")
	if err != nil {
		t.Fatal("undeploy", "tx", err)
	}
	t.Log("undeploy", "result", string(testClient.JSONMarshaler.MustMarshalJSON(res)))
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
