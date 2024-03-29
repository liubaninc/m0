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
		RecryptAccountList:      []*RecryptAccount{},
		ShareEncryptStorageList: []*ShareEncryptStorage{},
		EncryptStorageList:      []*EncryptStorage{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in recryptAccount
	recryptAccountIndexMap := make(map[string]bool)

	for _, elem := range gs.RecryptAccountList {
		if _, ok := recryptAccountIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for recryptAccount")
		}
		recryptAccountIndexMap[elem.Index] = true
	}

	// Check for duplicated index in shareEncryptStorage
	shareEncryptStorageIndexMap := make(map[string]bool)

	for _, elem := range gs.ShareEncryptStorageList {
		if _, ok := shareEncryptStorageIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for shareEncryptStorage")
		}
		shareEncryptStorageIndexMap[elem.Index] = true
	}
	// Check for duplicated index in encryptStorage
	encryptStorageIndexMap := make(map[string]bool)

	for _, elem := range gs.EncryptStorageList {
		if _, ok := encryptStorageIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for encryptStorage")
		}
		encryptStorageIndexMap[elem.Index] = true
	}

	return nil
}
