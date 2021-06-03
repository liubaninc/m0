package types

import (
	"encoding/hex"
	"encoding/json"

	"github.com/liubaninc/m0/x/wasm/xmodel"
	"gopkg.in/yaml.v2"
)

// Contract 合约
type Contract struct {
	Name                 string `json:"name"`      // 合约名
	Initiator            string `json:"initiator"` // 合约发行人
	Number               uint64 `json:"number"`    // 数量
	Address              string `json:"address"`   // 地址
	*xmodel.WasmCodeDesc        // 合约描述
}

type msgContract struct {
	Name         string `json:"name" yaml:"name"`           // 合约名
	Initiator    string `json:"initiator" yaml:"initiator"` // 合约发行人
	Number       uint64 `json:"number" yaml:"number"`       // 数量
	Address      string `json:"address" yaml:"address"`     // 地址
	Runtime      string `json:"runtime" yaml:"runtime"`
	Compiler     string `json:"compiler" yaml:"compiler"`
	Digest       string `json:"digest" yaml:"digest"`
	VmCompiler   string `json:"vm_compiler" yaml:"vm_compiler"`
	ContractType string `json:"contract_type" yaml:"contract_type"`
}

func (cc *Contract) MarshalYAML() (interface{}, error) {
	bs, err := yaml.Marshal(msgContract{
		Name:         cc.Name,
		Initiator:    cc.Initiator,
		Number:       cc.Number,
		Address:      cc.Address,
		Runtime:      cc.Runtime,
		Compiler:     cc.Compiler,
		Digest:       hex.EncodeToString(cc.Digest),
		VmCompiler:   cc.VmCompiler,
		ContractType: cc.ContractType,
	})

	if err != nil {
		return nil, err
	}

	return string(bs), nil
}

func (cc *Contract) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgContract{
		Name:         cc.Name,
		Initiator:    cc.Initiator,
		Number:       cc.Number,
		Address:      cc.Address,
		Runtime:      cc.Runtime,
		Compiler:     cc.Compiler,
		Digest:       hex.EncodeToString(cc.Digest),
		VmCompiler:   cc.VmCompiler,
		ContractType: cc.ContractType,
	})
}

func (cc *Contract) UnmarshalJSON(bz []byte) error {
	var alias msgContract
	if err := json.Unmarshal(bz, &alias); err != nil {
		return err
	}
	cc.Name = alias.Name
	cc.Initiator = alias.Initiator
	cc.Number = alias.Number
	cc.Address = alias.Address
	cc.WasmCodeDesc = &xmodel.WasmCodeDesc{}
	cc.Runtime = alias.Runtime
	cc.Compiler = alias.Compiler
	cc.Digest, _ = hex.DecodeString(alias.Digest)
	cc.VmCompiler = alias.VmCompiler
	cc.ContractType = alias.ContractType
	return nil
}
