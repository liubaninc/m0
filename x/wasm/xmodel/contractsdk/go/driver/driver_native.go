// +build !wasm

package driver

import (
	"github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/go/code"
	"github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/go/driver/native"
)

// Serve run contract in native environment
func Serve(contract code.Contract) {
	driver := native.New()
	driver.Serve(contract)
}
