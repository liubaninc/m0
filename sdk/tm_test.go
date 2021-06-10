package sdk

import (
	"encoding/json"
	"testing"
)

func TestTx(t *testing.T) {
	hash := "1070B88E8C6CB288FFCF741368D4CA6536F302C59851326639C46013EC5A5333"
	res, err := testClient.GetTx(hash)
	if err != nil {
		t.Fatal("GetTx", err)
	}
	bts, _ := json.Marshal(res)
	t.Log("GetTx", hash, string(bts))
}