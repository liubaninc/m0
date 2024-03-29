package bridge

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

type contractManager struct {
	xbridge *XBridge
}

// DeployContract deploy contract and initialize contract
func (c *contractManager) DeployContract(contextConfig *contract.ContextConfig, args map[string][]byte, number uint64) (*contract.Response, contract.Limits, error) {
	store := contextConfig.XMCache
	name := args["contract_name"]
	if name == nil {
		return nil, contract.Limits{}, errors.New("bad contract name")
	}
	contractName := string(name)
	_, err := newCodeProvider(contextConfig.XMReader).GetContractCodeDesc(contractName)
	if err == nil {
		return nil, contract.Limits{}, fmt.Errorf("contract %s already exists", contractName)
	}

	code := args["contract_code"]
	if code == nil {
		return nil, contract.Limits{}, errors.New("missing contract code")
	}
	initArgsBuf := args["init_args"]
	if initArgsBuf == nil {
		return nil, contract.Limits{}, errors.New("missing args field in args")
	}
	var initArgs map[string][]byte
	err = json.Unmarshal(initArgsBuf, &initArgs)
	if err != nil {
		return nil, contract.Limits{}, err
	}

	descbuf := args["contract_desc"]
	var desc xmodel.WasmCodeDesc
	err = proto.Unmarshal(descbuf, &desc)
	if err != nil {
		return nil, contract.Limits{}, err
	}
	desc.Digest = tmhash.Sum(code)
	descbuf, _ = proto.Marshal(&desc)

	store.Put("contract", ContractCodeDescKey(contractName), descbuf)
	store.Put("contract", ContractCodeKey(contractName), code)
	store.Put("contract", ContractNumberKey(contractName), []byte(strconv.FormatUint(number, 10)))

	contractType, err := getContractType(&desc)
	if err != nil {
		return nil, contract.Limits{}, err
	}
	creator := c.xbridge.getCreator(contractType)
	if creator == nil {
		return nil, contract.Limits{}, fmt.Errorf("contract type %s not found", contractType)
	}
	cp := newCodeProvider(store)
	instance, err := creator.CreateInstance(&Context{
		ContractName:   contractName,
		ResourceLimits: contextConfig.ResourceLimits,
	}, cp)
	if err != nil {
		creator.RemoveCache(contractName)
		return nil, contract.Limits{}, fmt.Errorf("create contract instance error when deploy contract %v, error %v", contractName, err)
	}
	instance.Release()

	initConfig := *contextConfig
	initConfig.ContractName = contractName
	initConfig.CanInitialize = true
	initConfig.ContractCodeFromCache = true
	out, resourceUsed, err := c.initContract(contractType, &initConfig, initArgs)
	if err != nil {
		if _, ok := err.(*ContractError); !ok {
			creator.RemoveCache(contractName)
		}
		return nil, contract.Limits{}, fmt.Errorf("call contract %v initialize method, error %v", contractName, err)
	}
	return out, resourceUsed, nil
}

func (v *contractManager) initContract(tp ContractType, contextConfig *contract.ContextConfig, args map[string][]byte) (*contract.Response, contract.Limits, error) {
	vm, ok := v.xbridge.GetVirtualMachine(string(tp))
	if !ok {
		return nil, contract.Limits{}, errors.New("wasm vm not registered")
	}

	ctx, err := vm.NewContext(contextConfig)
	if err != nil {
		return nil, contract.Limits{}, err
	}
	out, err := ctx.Invoke("initialize", args)
	if err != nil {
		return nil, contract.Limits{}, err
	}
	return out, ctx.ResourceUsed(), nil
}

// UpgradeContract deploy contract and initialize contract
func (c *contractManager) UpgradeContract(contextConfig *contract.ContextConfig, args map[string][]byte) (*contract.Response, contract.Limits, error) {
	if !c.xbridge.config.EnableUpgrade {
		return nil, contract.Limits{}, errors.New("contract upgrade disabled")
	}

	name := args["contract_name"]
	if name == nil {
		return nil, contract.Limits{}, errors.New("bad contract name")
	}
	contractName := string(name)
	desc, err := newCodeProvider(contextConfig.XMReader).GetContractCodeDesc(contractName)
	if err != nil {
		return nil, contract.Limits{}, fmt.Errorf("contract %s not exists", contractName)
	}

	code := args["contract_code"]
	if code == nil {
		return nil, contract.Limits{}, errors.New("missing contract code")
	}
	desc.Digest = tmhash.Sum(code)
	descbuf, _ := proto.Marshal(desc)

	store := contextConfig.XMCache
	store.Put("contract", ContractCodeDescKey(contractName), descbuf)
	store.Put("contract", ContractCodeKey(contractName), code)

	cp := newCodeProvider(store)

	contractType, err := getContractType(desc)
	if err != nil {
		return nil, contract.Limits{}, err
	}
	creator := c.xbridge.getCreator(contractType)
	if creator == nil {
		return nil, contract.Limits{}, fmt.Errorf("contract type %s not found", contractType)
	}
	instance, err := creator.CreateInstance(&Context{
		ContractName:   contractName,
		ResourceLimits: contract.MaxLimits,
	}, cp)
	if err != nil {
		fmt.Errorf("create contract instance error when upgrade contract %v, error %v", contractName, err)
		return nil, contract.Limits{}, err
	}
	instance.Release()

	return &contract.Response{
			Status: 200,
			Body:   []byte("upgrade success"),
		}, contract.Limits{
			Disk: modelCacheDiskUsed(store),
		}, nil
}

func modelCacheDiskUsed(cache xmodel.XModel) int64 {
	size := int64(0)
	_, wset, _ := cache.GetRWSets()
	for _, w := range wset {
		size += int64(len(w.GetKey()))
		size += int64(len(w.GetValue()))
	}
	return size
}

func ContractCodeDescKey(contractName string) []byte {
	return []byte(contractName + "." + "desc")
}

func ContractCodeKey(contractName string) []byte {
	return []byte(contractName + "." + "code")
}

func ContractNumberKey(contractName string) []byte {
	return []byte(contractName + "." + "num")
}

func getContractType(desc *xmodel.WasmCodeDesc) (ContractType, error) {
	switch desc.ContractType {
	case "", "wasm":
		return TypeWasm, nil
	case "native":
		return TypeNative, nil
	default:
		return "", fmt.Errorf("unknown contract type:%s", desc.ContractType)
	}
}
