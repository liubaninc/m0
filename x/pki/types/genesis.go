package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		CertificateList:        []*Certificate{},
		ChildCertificatesList:  []*ChildCertificates{},
		RevokeCertificatesList: []*Certificates{},
		CertificatesList:       []*Certificates{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in certificate
	certificateIndexMap := make(map[string]bool)

	for _, elem := range gs.CertificateList {
		identifier := CertificateIdentifier{
			Issuer:       elem.Issuer,
			SerialNumber: elem.SerialNumber,
		}
		if _, ok := certificateIndexMap[identifier.Index()]; ok {
			return fmt.Errorf("duplicated index for certificate")
		}
		certificateIndexMap[identifier.Index()] = true
	}
	// Check for duplicated index in certificates
	certificatesIndexMap := make(map[string]bool)

	for _, elem := range gs.CertificatesList {
		if _, ok := certificatesIndexMap[elem.Identifier.Index()]; ok {
			return fmt.Errorf("duplicated index for certificates")
		}
		certificatesIndexMap[elem.Identifier.Index()] = true
	}

	// Check for duplicated index in certificates
	revokeCertificatesIndexMap := make(map[string]bool)

	for _, elem := range gs.RevokeCertificatesList {
		if _, ok := revokeCertificatesIndexMap[elem.Identifier.Index()]; ok {
			return fmt.Errorf("duplicated index for certificates")
		}
		revokeCertificatesIndexMap[elem.Identifier.Index()] = true
	}

	// Check for duplicated index in certificates
	childCertificatesIndexMap := make(map[string]bool)

	for _, elem := range gs.ChildCertificatesList {
		if _, ok := childCertificatesIndexMap[elem.Identifier.Index()]; ok {
			return fmt.Errorf("duplicated index for certificates")
		}
		childCertificatesIndexMap[elem.Identifier.Index()] = true
	}
	return nil
}
