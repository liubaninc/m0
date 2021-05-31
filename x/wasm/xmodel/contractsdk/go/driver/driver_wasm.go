// +build wasm

package driver

import (
	"github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/go/code"
	"github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/go/driver/wasm"
)

// Serve run contract in wasm environment
func Serve(contract code.Contract) {
	driver := wasm.New()
	driver.Serve(contract)
}
