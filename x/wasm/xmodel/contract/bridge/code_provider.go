package bridge

import (
	"errors"
	"fmt"

	"github.com/liubaninc/m0/x/wasm/xmodel"

	"github.com/golang/protobuf/proto"
)

type xmodelStore interface {
	Get(bucket string, key []byte) (*xmodel.VersionedData, error)
}

type codeProvider struct {
	xstore xmodelStore
}

func newCodeProvider(xstore xmodelStore) ContractCodeProvider {
	return &codeProvider{
		xstore: xstore,
	}
}

func (c *codeProvider) GetContractCode(name string) ([]byte, error) {
	value, err := c.xstore.Get("contract", ContractCodeKey(name))
	if err != nil {
		return nil, fmt.Errorf("get contract code for '%s' error:%s", name, err)
	}
	codebuf := value.GetPureData().GetValue()
	if len(codebuf) == 0 {
		return nil, errors.New("empty wasm code")
	}
	return codebuf, nil
}

func (c *codeProvider) GetContractCodeDesc(name string) (*xmodel.WasmCodeDesc, error) {
	value, err := c.xstore.Get("contract", ContractCodeDescKey(name))
	if err != nil {
		return nil, fmt.Errorf("get contract desc for '%s' error:%s", name, err)
	}
	descbuf := value.GetPureData().GetValue()
	// FIXME: 如果key不存在ModuleCache不应该返回零长度的value
	if len(descbuf) == 0 {
		return nil, errors.New("empty wasm code desc")
	}
	var desc xmodel.WasmCodeDesc
	err = proto.Unmarshal(descbuf, &desc)
	if err != nil {
		return nil, err
	}
	return &desc, nil
}

type descProvider struct {
	ContractCodeProvider
	desc *xmodel.WasmCodeDesc
}

func newDescProvider(cp ContractCodeProvider, desc *xmodel.WasmCodeDesc) ContractCodeProvider {
	return &descProvider{
		ContractCodeProvider: cp,
		desc:                 desc,
	}
}

func (d *descProvider) GetContractCodeDesc(name string) (*xmodel.WasmCodeDesc, error) {
	return d.desc, nil
}
