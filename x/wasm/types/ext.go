package types

import (
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
)

var (
	// TransientBucket is the name of bucket that only appears in tx output set but does't persists in xmodel
	TransientBucket       = "$transient"
	ContractUtxoInputKey  = "ContractUtxo.Inputs"
	ContractUtxoOutputKey = "ContractUtxo.Outputs"
	ContractEventKey      = "ContractEvent"
)

// ParseInputsFromExt parse inputs from outputsExt
func ParseInputsFromExt(outputsExt []*OutputExt) ([]*utxotypes.Input, error) {
	var inputs []*utxotypes.Input
	for _, outputExt := range outputsExt {
		if outputExt.Bucket != TransientBucket || outputExt.Key != ContractUtxoInputKey {
			continue
		}
		// TODO chaogaofeng
		break
	}
	return inputs, nil
}

// ParseOutputsFromExt parse outputs from outputsExt
func ParseOutputsFromExt(outputsExt []*OutputExt) ([]*utxotypes.Output, error) {
	var outputs []*utxotypes.Output
	for _, outputExt := range outputsExt {
		if outputExt.Bucket != TransientBucket || outputExt.Key != ContractUtxoOutputKey {
			continue
		}
		// TODO chaogaofeng
		break
	}
	return outputs, nil
}

// ParseEventsFromExt parse events from outputsExt
func ParseEventsFromExt(outputsExt []*OutputExt) ([]*ContractEvent, error) {
	var events []*ContractEvent
	for _, outputExt := range outputsExt {
		if outputExt.Bucket != TransientBucket || outputExt.Key != ContractEventKey {
			continue
		}
		// TODO chaogaofeng
		break
	}
	return events, nil
}
