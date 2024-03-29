package types

const (
	// ModuleName defines the module name
	ModuleName = "pki"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_pki"

	// this line is used by starport scaffolding # ibc/keys/name
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	CertificatesKey        = "Certificates-value-"
	RevokedCertificatesKey = "Revoked-certificates-value-"
	ChildCertificatesKey   = "Child-certificates-value-"
	CertificateKey         = "Certificate-value-"
)
