package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "authority"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_authority"

	// this line is used by starport scaffolding # ibc/keys/name
)

var (
	PendingAccountPrefix           = []byte{0x01} // prefix for each key to a pending account
	AccountPrefix                  = []byte{0x02} // prefix for each key to an account
	PendingAccountRevocationPrefix = []byte{0x03} // prefix for each key to a pending account revocation

	AccountNumberCounterKey = []byte("globalAccountNumber") // key for account number counter
)

// Key builder for Account.
func GetAccountKey(addr sdk.AccAddress) []byte {
	return append(AccountPrefix, addr.Bytes()...)
}

// Key builder for Pending Account.
func GetPendingAccountKey(addr sdk.AccAddress) []byte {
	return append(PendingAccountPrefix, addr.Bytes()...)
}
