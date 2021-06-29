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
		CommentList: []*Comment{},
		BlobList:    []*Blob{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in comment
	commentIdMap := make(map[uint64]bool)

	for _, elem := range gs.CommentList {
		if _, ok := commentIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for comment")
		}
		commentIdMap[elem.Id] = true
	}
	// Check for duplicated ID in blob
	blobIdMap := make(map[uint64]bool)

	for _, elem := range gs.BlobList {
		if _, ok := blobIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for blob")
		}
		blobIdMap[elem.Id] = true
	}

	return nil
}
