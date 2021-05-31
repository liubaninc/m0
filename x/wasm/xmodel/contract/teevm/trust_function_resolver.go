package teevm

import (
	"errors"
	"plugin"

	"github.com/liubaninc/m0/x/wasm/xmodel/xvm/exec"
	"github.com/liubaninc/m0/x/wasm/xmodel/xvm/runtime/emscripten"
)

// TEEConfig sets up the private ledger
type TEEConfig struct {
	Enable     bool   `yaml:"enable"`     // enable: on or off to enable private ledger
	PluginPath string `yaml:"pluginPath"` // path to dynamic library
	ConfigPath string `yaml:"configPath"` // config path for the dynamic
}

// TrustFunctionResolver
type TrustFunctionResolver struct {
	handler *plugin.Plugin
	runFunc func([]byte) ([]byte, error)
}

var _ exec.Resolver = (*TrustFunctionResolver)(nil)

//
//
func NewTrustFunctionResolver(conf TEEConfig) (*TrustFunctionResolver, error) {
	if conf.Enable == false {
		return nil, errors.New("private ledger is not enabled")
	}
	p, err := plugin.Open(conf.PluginPath)
	if err != nil {
		return nil, err
	}
	initFuncRaw, err := p.Lookup("Init")
	if err != nil {
		return nil, err
	}
	initFunc, ok := initFuncRaw.(func(string) error)
	if !ok {
		return nil, errors.New(conf.PluginPath + " doesn't implement Init(string) error")
	}
	if err := initFunc(conf.ConfigPath); err != nil {
		return nil, err
	}
	runFuncRaw, err := p.Lookup("Run")
	if err != nil {
		return nil, err
	}
	runFunc, ok := runFuncRaw.(func([]byte) ([]byte, error))
	if !ok {
		return nil, errors.New(conf.PluginPath + " doesn't implement Run([]byte) ([]byte, error)")
	}
	return &TrustFunctionResolver{handler: p, runFunc: runFunc}, nil
}

func (tf *TrustFunctionResolver) ResolveGlobal(module, name string) (int64, bool) {
	return 0, false
}

func (tf *TrustFunctionResolver) ResolveFunc(module, name string) (interface{}, bool) {
	fullname := module + "." + name
	switch fullname {
	case "env._xvm_tfcall":
		return tf.tfcall, true
	default:
		return nil, false
	}
}

func (tf *TrustFunctionResolver) tfcall(ctx exec.Context, inptr, inlen, outpptr, outlenptr uint32) uint32 {
	var (
		retCode uint32 = 0
	)
	codec := exec.NewCodec(ctx)
	requestBuf := codec.Bytes(inptr, inlen)
	responseBuf, err := tf.runFunc(requestBuf)

	if err != nil {
		copy(responseBuf, []byte(err.Error()))
		retCode = 1
	}
	codec.SetUint32(outpptr, bytesdup(ctx, responseBuf))
	codec.SetUint32(outlenptr, uint32(len(responseBuf)))
	return retCode
}

//copied from https://github.com/xuperchain/xuperchain/blob/master/core/contract/wasm/vm/xvm/builtin_resolver.go#L180, TODO refer not copy
func bytesdup(ctx exec.Context, b []byte) uint32 {
	codec := exec.NewCodec(ctx)
	memptr, err := emscripten.Malloc(ctx, len(b))
	if err != nil {
		exec.ThrowError(err)
	}
	mem := codec.Bytes(memptr, uint32(len(b)))
	copy(mem, b)
	return memptr
}
