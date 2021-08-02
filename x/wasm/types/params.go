package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultEnabled = true
)

var (
	KeyEnabled = []byte("KeyEnabled")
)

// ParamKeyTable for bank module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new parameter configuration for the module
func NewParams(enabled bool) Params {
	return Params{
		Enabled: enabled,
	}
}

// DefaultParams is the default parameter configuration for the module
func DefaultParams() Params {
	return Params{
		Enabled: DefaultEnabled,
	}
}

// Validate all bank module parameters
func (p Params) Validate() error {
	return validateIsBool(p.Enabled)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyEnabled, &p.Enabled, validateIsBool),
	}
}

func validateIsBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
