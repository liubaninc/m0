package types

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

const (
	PendingAccountPrefix           = "Pending-Account-Prefix-"            // prefix for each key to a pending account
	AccountPrefix                  = "Account-Prefix-"                    // prefix for each key to an account
	PendingAccountRevocationPrefix = "Pending-Account-Revocation-Prefix-" // prefix for each key to a pending account revocation

	AccountNumberCounterKey = "Global-Account-Number-" // key for account number counter
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
