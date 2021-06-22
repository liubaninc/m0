package types

import (
	"fmt"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId: PortID,
		// this line is used by starport scaffolding # genesis/types/default
		ItxList: []*Itx{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in itx
	itxIdMap := make(map[uint64]bool)

	for _, elem := range gs.ItxList {
		if _, ok := itxIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for itx")
		}
		itxIdMap[elem.Id] = true
	}

	return nil
}
