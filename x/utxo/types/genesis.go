package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:    DefaultParams(),
		TokenList: []*Token{},
		InputList: []*Input{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	// Check for duplicated index in token
	tokenIndexMap := make(map[string]bool)
	for _, elem := range gs.TokenList {
		if err := sdk.ValidateDenom(elem.Name); err != nil {
			return fmt.Errorf("invalid denom %s: %w", elem.Name, err)
		}
		if _, ok := tokenIndexMap[elem.Name]; ok {
			return fmt.Errorf("duplicated name %s for token", elem.Name)
		}
		tokenIndexMap[elem.Name] = true
	}

	// Check for duplicated index in input
	inputIndexMap := make(map[string]bool)
	for _, elem := range gs.InputList {
		if _, ok := tokenIndexMap[elem.Index()]; ok {
			return fmt.Errorf("duplicated name %s for input", elem.Index())
		}
		if _, ok := tokenIndexMap[elem.Amount.Denom]; !ok {
			return fmt.Errorf("token %s for input not exist", elem.Amount.Denom)
		}
		inputIndexMap[elem.Index()] = true
	}

	return nil
}
