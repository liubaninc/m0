package types

import (
	"bytes"
	"encoding/json"
	"fmt"

	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
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
		err := json.Unmarshal(outputExt.Value, &inputs)
		if err != nil {
			return nil, err
		}
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
		err := json.Unmarshal(outputExt.Value, &outputs)
		if err != nil {
			return nil, err
		}
		break
	}
	return outputs, nil
}

// ParseEventsFromExt parse events from outputsExt
func ParseEventsFromExt(outputsExt []*OutputExt) ([]*xmodel.ContractEvent, error) {
	var events []*xmodel.ContractEvent
	for _, outputExt := range outputsExt {
		if outputExt.Bucket != TransientBucket || outputExt.Key != ContractEventKey {
			continue
		}
		err := json.Unmarshal(outputExt.Value, &events)
		if err != nil {
			return nil, err
		}
		break
	}
	return events, nil
}

const (
	// DelFlag delete flag
	DelFlag = "\x00"
	// BucketSeperator separator between bucket and raw key
	BucketSeperator = "/"
)

// IsDelFlag value是否是删除标志
func IsDelFlag(value []byte) bool {
	return bytes.Equal([]byte(DelFlag), value)
}

func MakeVersion(txID []byte, msgOffset int32, offset int32) string {
	return fmt.Sprintf("%x_%d_%d", txID, msgOffset, offset)
}

func ParseVersion(version string) ([]byte, int32, int32, error) {
	var txID []byte
	msgOffset := int32(0)
	offset := int32(0)
	okNum, err := fmt.Sscanf(version, "%x_%d_%d", &txID, &msgOffset, &offset)
	if okNum != 3 && err != nil {
		return nil, 0, 0, fmt.Errorf("parseVersion failed, invalid version: %s", version)
	}
	return txID, msgOffset, offset, nil
}

// MakeRawKey 组装Key
func MakeRawKey(bucket string, key []byte) []byte {
	k := append([]byte(bucket), []byte(BucketSeperator)...)
	return append(k, key...)
}

// ParseRawKey 解析Key
func ParseRawKey(rawKey []byte) (string, []byte, error) {
	idx := bytes.Index(rawKey, []byte(BucketSeperator))
	if idx < 0 {
		return "", nil, fmt.Errorf("parseRawKey failed, invalid raw key:%s", string(rawKey))
	}
	bucket := string(rawKey[:idx])
	key := rawKey[idx+len(BucketSeperator):]
	return bucket, key, nil
}

// IsEmptyVersionedData check if VersionedData is empty
func IsEmptyVersionedData(vd *xmodel.VersionedData) bool {
	return vd.RefTxid == nil && vd.RefOffset == 0
}

func MakeEmptyVersionedData(bucket string, key []byte) *xmodel.VersionedData {
	verData := &xmodel.VersionedData{PureData: &xmodel.PureData{}}
	verData.PureData.Bucket = bucket
	verData.PureData.Key = key
	return verData
}

// GetOutputsExt get transaction outputs_ext
func GetOutputsExt(pds []*xmodel.PureData) []*OutputExt {
	outputs := make([]*OutputExt, 0, len(pds))
	for _, pd := range pds {
		outputs = append(outputs, &OutputExt{
			Bucket: pd.Bucket,
			Key:    string(pd.Key),
			Value:  pd.Value,
		})
	}
	return outputs
}

// GetInputsExt get transaction inputs_ext
func GetInputsExt(vds []*xmodel.VersionedData) []*InputExt {
	inputs := make([]*InputExt, 0, len(vds))
	for _, vd := range vds {
		inputs = append(inputs, &InputExt{
			Bucket:    vd.PureData.Bucket,
			Key:       string(vd.PureData.Key),
			RefTx:     string(vd.RefTxid),
			RefMsg:    vd.RefMsgOffset,
			RefOffset: vd.RefOffset,
		})
	}
	return inputs
}
