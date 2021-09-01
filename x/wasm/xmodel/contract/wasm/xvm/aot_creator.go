package xvm

import (
	"fmt"
	"io/ioutil"
	"os"
	osexec "os/exec"
	"path/filepath"

	tmlog "github.com/tendermint/tendermint/libs/log"

	"github.com/liubaninc/m0/x/wasm/xmodel/contract/bridge"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/teevm"
	"github.com/liubaninc/m0/x/wasm/xmodel/xvm/compile"
	"github.com/liubaninc/m0/x/wasm/xmodel/xvm/exec"
	"github.com/liubaninc/m0/x/wasm/xmodel/xvm/runtime/emscripten"
	gowasm "github.com/liubaninc/m0/x/wasm/xmodel/xvm/runtime/go"
)

var (
	log = tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout))
)

// XVMConfig contains the xvm configuration
type XVMConfig struct {
	// From 0 to 3
	// The higher the number, the faster the program runs,
	// but the compilation speed will be slower
	OptLevel int `yaml:"optlevel"`
}

// WasmConfig wasm config
type WasmConfig struct {
	Driver        string
	External      bool
	XVM           XVMConfig
	EnableUpgrade bool
	TEEConfig     teevm.TEEConfig `yaml:"teeConfig,omitempty"`
}

func (w *WasmConfig) DriverName() string {
	return w.Driver
}

func (w *WasmConfig) IsEnable() bool {
	return true
}

type xvmCreator struct {
	cm       *codeManager
	config   bridge.InstanceCreatorConfig
	vmconfig *WasmConfig

	wasm2cPath string
}

// 优先查找跟mchain同级目录的二进制，再在PATH里面找
func lookupWasm2c() (string, error) {
	// 首先查找跟mchain同级的目录
	wasm2cPath := filepath.Join(filepath.Dir(os.Args[0]), "wasm2c")
	stat, err := os.Stat(wasm2cPath)
	if err == nil {
		if m := stat.Mode(); !m.IsDir() && m&0111 != 0 {
			return filepath.Abs(wasm2cPath)
		}
	}
	// 再查找系统PATH目录
	return osexec.LookPath("wasm2c")
}

func newXVMCreator(creatorConfig *bridge.InstanceCreatorConfig) (bridge.InstanceCreator, error) {
	wasm2cPath, err := lookupWasm2c()
	if err != nil {
		return nil, err
	}
	creator := &xvmCreator{
		wasm2cPath: wasm2cPath,
		config:     *creatorConfig,
	}
	if creatorConfig.VMConfig != nil {
		creator.vmconfig = creatorConfig.VMConfig.(*WasmConfig)
		optlevel := creator.vmconfig.XVM.OptLevel
		if optlevel < 0 || optlevel > 3 {
			return nil, fmt.Errorf("bad xvm optlevel:%d", optlevel)
		}
	}
	creator.cm, err = newCodeManager(creator.config.Basedir,
		creator.CompileCode, creator.MakeExecCode)
	if err != nil {
		return nil, err
	}
	return creator, nil
}

func cpfile(dest, src string) error {
	buf, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dest, buf, 0700)
}

func (x *xvmCreator) CompileCode(buf []byte, outputPath string) error {
	tmpdir, err := ioutil.TempDir("", "xvm-compile")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpdir)
	wasmpath := filepath.Join(tmpdir, "code.wasm")
	err = ioutil.WriteFile(wasmpath, buf, 0600)
	if err != nil {
		return err
	}

	libpath := filepath.Join(tmpdir, "code.so")

	cfg := &compile.Config{
		Wasm2cPath: x.wasm2cPath,
		OptLevel:   x.vmconfig.XVM.OptLevel,
	}
	err = compile.CompileNativeLibrary(cfg, libpath, wasmpath)
	if err != nil {
		return err
	}
	return cpfile(outputPath, libpath)
}

func (x *xvmCreator) getContractCodeCache(name string, cp bridge.ContractCodeProvider) (*contractCode, error) {
	return x.cm.GetExecCode(name, cp)
}

func (x *xvmCreator) MakeExecCode(libpath string) (exec.Code, error) {
	resolvers := []exec.Resolver{
		gowasm.NewResolver(),
		emscripten.NewResolver(),
		newSyscallResolver(x.config.SyscallService),
		builtinResolver,
	}
	//AOT only for experiment;
	if x.vmconfig.TEEConfig.Enable {
		teeResolver, err := teevm.NewTrustFunctionResolver(x.vmconfig.TEEConfig)
		if err != nil {
			return nil, err
		}
		resolvers = append(resolvers, teeResolver)
	}
	resolver := exec.NewMultiResolver(
		resolvers...,
	)
	return exec.NewAOTCode(libpath, resolver)
}

func (x *xvmCreator) CreateInstance(ctx *bridge.Context, cp bridge.ContractCodeProvider) (bridge.Instance, error) {
	code, err := x.getContractCodeCache(ctx.ContractName, cp)
	if err != nil {
		//log.Error("get contract cache error", "error", err, "contract", ctx.ContractName)
		return nil, err
	}

	return createInstance(ctx, code, x.config.SyscallService)
}

func (x *xvmCreator) RemoveCache(contractName string) {
	x.cm.RemoveCode(contractName)
}

func init() {
	bridge.Register(bridge.TypeWasm, "xvm", newXVMCreator)
}
