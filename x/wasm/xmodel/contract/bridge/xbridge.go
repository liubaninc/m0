package bridge

import (
	"path/filepath"

	"github.com/liubaninc/m0/x/wasm/xmodel/contract"
)

// ContractConfig define the config of XuperBridge
type ContractConfig struct {
	EnableUpgrade bool
}

type XBridgeConfig struct {
	Basedir   string
	VMConfigs map[ContractType]VMConfig
	Config    ContractConfig
}

// XBridge 用于注册用户虚拟机以及向其他注册可被识别的vm.VirtualMachine
type XBridge struct {
	ctxmgr         *ContextManager
	syscallService *SyscallService
	basedir        string
	vmconfigs      map[ContractType]VMConfig
	creators       map[ContractType]InstanceCreator
	vms            map[string]contract.VirtualMachine
	config         ContractConfig

	*contractManager
}

// New instances a new XBridge
func New(cfg *XBridgeConfig) (*XBridge, error) {
	ctxmgr := NewContextManager()
	xbridge := &XBridge{
		ctxmgr:    ctxmgr,
		basedir:   cfg.Basedir,
		vmconfigs: cfg.VMConfigs,
		creators:  make(map[ContractType]InstanceCreator),
		vms:       make(map[string]contract.VirtualMachine),
		config:    cfg.Config,
	}
	xbridge.contractManager = &contractManager{
		xbridge: xbridge,
	}

	syscallService := NewSyscallService(ctxmgr, xbridge)
	xbridge.syscallService = syscallService
	err := xbridge.initVM()
	if err != nil {
		return nil, err
	}

	return xbridge, nil
}

func (v *XBridge) initVM() error {
	types := []ContractType{TypeWasm, TypeNative}
	for _, tp := range types {
		vmconfig, ok := v.vmconfigs[tp]
		if !ok {
			continue
		}
		if !vmconfig.IsEnable() {
			continue
		}
		creatorConfig := &InstanceCreatorConfig{
			Basedir:        filepath.Join(v.basedir, vmconfig.DriverName()),
			SyscallService: v.syscallService,
			VMConfig:       vmconfig,
		}
		creator, err := Open(tp, vmconfig.DriverName(), creatorConfig)
		if err != nil {
			return err
		}
		vm := &vmImpl{
			ctxmgr:  v.ctxmgr,
			xbridge: v,
			name:    string(tp),
		}
		v.creators[tp] = creator
		v.vms[string(tp)] = vm
	}
	return nil
}

func (v *XBridge) getCreator(tp ContractType) InstanceCreator {
	return v.creators[tp]
}

// GetVirtualMachine returns a contract.VirtualMachine from the given name
func (v *XBridge) GetVirtualMachine(name string) (contract.VirtualMachine, bool) {
	vm, ok := v.vms[name]
	return vm, ok
}

// RegisterToXCore register VirtualMachines to xchain core
func (v *XBridge) RegisterToXCore(regfunc func(name string, vm contract.VirtualMachine) error) {
	for _, vm := range v.vms {
		err := regfunc(vm.GetName(), vm)
		if err != nil {
			panic(err)
		}
	}
}
