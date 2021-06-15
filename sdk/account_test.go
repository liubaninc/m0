package sdk

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func TestAccount(t *testing.T) {
	address := authtypes.NewModuleAddress(authtypes.FeeCollectorName).String()
	acct, err := testClient.GetAccount(address, 0)
	if err != nil {
		t.Fatal("GetAccount", address, err)
	}
	t.Log("GetAccount", address, string(testClient.JSONMarshaler.MustMarshalJSON(acct)))
}

func TestAccounts(t *testing.T) {
	accts, err := testClient.GetAccounts(nil, 0, 1, true)
	if err != nil {
		t.Fatal("GetAccounts", err)
	}
	for i := 0; i < int(accts.Pagination.Total); i++ {
		ats, err := testClient.GetAccounts(nil, uint64(i), 1, true)
		if err != nil {
			t.Fatal("GetAccounts", "offset", i, err)
		}
		t.Log("GetAccounts", "offset", i, string(testClient.JSONMarshaler.MustMarshalJSON(ats)))
	}
}

func TestAccounts2(t *testing.T) {
	accts, err := testClient.GetAccounts(nil, 0, 1, true)
	if err != nil {
		t.Fatal("GetAccounts", err)
	}
	t.Log("GetAccounts", string(testClient.JSONMarshaler.MustMarshalJSON(accts)))
	for accts.Pagination.NextKey != nil {
		accts, err = testClient.GetAccounts(accts.Pagination.NextKey, 0, 1, true)
		if err != nil {
			t.Fatal("GetAccounts", err)
		}
		t.Log("GetAccounts", string(testClient.JSONMarshaler.MustMarshalJSON(accts)))
	}
}

func TestAccountBalance(t *testing.T) {
	address := authtypes.NewModuleAddress(authtypes.FeeCollectorName).String()
	item, err := testClient.GetAccountBalance(address, "stake")
	if err != nil {
		t.Fatal("GetAccountBalance", err)
	}
	t.Log("GetAccountBalance", address, string(testClient.JSONMarshaler.MustMarshalJSON(item)))
}

func TestAccountBalances(t *testing.T) {
	address := authtypes.NewModuleAddress(authtypes.FeeCollectorName).String()
	items, err := testClient.GetAccountBalances(address, nil, 0, 1, true)
	if err != nil {
		t.Fatal("GetAccountBalances", err)
	}
	t.Log("GetAccountBalances", string(testClient.JSONMarshaler.MustMarshalJSON(items)))
	for items.Pagination.NextKey != nil {
		items, err = testClient.GetAccountBalances(address, items.Pagination.NextKey, 0, 1, true)
		if err != nil {
			t.Fatal("GetAccountBalances", err)
		}
		t.Log("GetAccountBalances", string(testClient.JSONMarshaler.MustMarshalJSON(items)))
	}
}

func TestTotalSupply(t *testing.T) {
	item, err := testClient.GetTotalSupply()
	if err != nil {
		t.Fatal("GetTotalSupply", err)
	}
	t.Log("GetTotalSupply", string(testClient.JSONMarshaler.MustMarshalJSON(item)))
}

func TestSupplyOf(t *testing.T) {
	denom := "stake"
	item, err := testClient.GetSupply(denom)
	if err != nil {
		t.Fatal("GetSupply", err)
	}
	t.Log("GetSupply", denom, string(testClient.JSONMarshaler.MustMarshalJSON(item)))
}
