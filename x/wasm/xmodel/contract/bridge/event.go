package bridge

import (
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract"
)

func eventsResourceUsed(events []*xmodel.ContractEvent) contract.Limits {
	var size int64
	for _, event := range events {
		size += int64(len(event.Contract) + len(event.Name) + len(event.Body))
	}
	return contract.Limits{
		Disk: size,
	}
}
