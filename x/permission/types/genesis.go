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
		AccountList: []*Account{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in account
	accountIndexMap := make(map[string]bool)

	for _, elem := range gs.AccountList {
		if _, ok := accountIndexMap[elem.Address]; ok {
			return fmt.Errorf("duplicated index for account")
		}
		accountIndexMap[elem.Address] = true
	}

	return nil
}
