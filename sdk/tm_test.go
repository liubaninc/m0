package sdk

import (
	"encoding/json"
	"testing"
)

func TestTx(t *testing.T) {
	hash := "D6E2B26B0EFE390CE4C3E3F46400B50644B6806EE72EDBF733D182F1E8495769"
	res, err := testClient.GetTx(hash)
	if err != nil {
		t.Fatal("GetTx", err)
	}
	bts, _ := json.Marshal(res)
	t.Log("GetTx", hash, string(bts))
}
