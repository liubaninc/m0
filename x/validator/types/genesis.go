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
		Params:                 DefaultParams(),
		LastValidatorPowerList: []*LastValidatorPower{},
		ValidatorList:          []*Validator{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in lastValidatorPower
	lastValidatorPowerIndexMap := make(map[string]bool)

	for _, elem := range gs.LastValidatorPowerList {
		if _, ok := lastValidatorPowerIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for lastValidatorPower")
		}
		lastValidatorPowerIndexMap[elem.Index] = true
	}
	// Check for duplicated ID in validator
	validatorIdMap := make(map[string]bool)

	for _, elem := range gs.ValidatorList {
		if _, ok := validatorIdMap[elem.Creator]; ok {
			return fmt.Errorf("duplicated id for validato")
		}
		validatorIdMap[elem.Creator] = true
	}

	return nil
}
