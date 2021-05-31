package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"sort"
)

// 快速对写集合排序
type pdSlice []*xmodel.PureData

// newPdSlice new a slice instance for PureData
func newPdSlice(vpd []*xmodel.PureData) pdSlice {
	s := make([]*xmodel.PureData, len(vpd))
	copy(s, vpd)
	return s
}

// Len length of slice of PureData
func (pds pdSlice) Len() int {
	return len(pds)
}

// Swap swap two pureData elements in a slice
func (pds pdSlice) Swap(i, j int) {
	pds[i], pds[j] = pds[j], pds[i]
}

// Less compare two pureData elements with pureData's key in a slice
func (pds pdSlice) Less(i, j int) bool {
	rawKeyI := MakeRawKey(pds[i].Bucket, pds[i].Key)
	rawKeyJ := MakeRawKey(pds[j].Bucket, pds[j].Key)
	ret := bytes.Compare(rawKeyI, rawKeyJ)
	if ret == 0 {
		// 注: 正常应该无法走到这个逻辑，因为写集合中的key一定是唯一的
		return bytes.Compare(pds[i].Value, pds[j].Value) < 0
	}
	return ret < 0
}

func equal(pd, vpd *xmodel.PureData) bool {
	rawKeyI := MakeRawKey(pd.Bucket, pd.Key)
	rawKeyJ := MakeRawKey(vpd.Bucket, vpd.Key)
	ret := bytes.Compare(rawKeyI, rawKeyJ)
	if ret != 0 {
		return false
	}
	return bytes.Equal(pd.Value, vpd.Value)
}

// Equal check if two PureData object equal
func Equal(pd, vpd []*xmodel.PureData) bool {
	if len(pd) != len(vpd) {
		return false
	}
	pds := newPdSlice(pd)
	vpds := newPdSlice(vpd)
	sort.Sort(pds)
	sort.Sort(vpds)
	for i, v := range pds {
		if equal(v, vpds[i]) {
			continue
		}
		return false
	}
	return true
}

func ConvertToMethodArgs(argsStr string) (map[string][]byte, error) {
	if len(argsStr) == 0 {
		return nil, nil
	}
	var args map[string]interface{}
	if err := json.Unmarshal([]byte(argsStr), &args); err != nil {
		return nil, err
	}

	argmap := make(map[string][]byte)
	for k, v := range args {
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("bad key %s, expect string value, got %v", k, v)
		}
		argmap[k] = []byte(s)
	}
	return argmap, nil
}
