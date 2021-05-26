package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultMaxDescCharacters uint64 = 256
)

var (
	KeyMaxDescCharacters = []byte("KeyMaxDescCharacters")
)

// ParamKeyTable for bank module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new parameter configuration for the bank module
func NewParams(maxDescCharacters uint64) Params {
	return Params{
		MaxDescCharacters: maxDescCharacters,
	}
}

// DefaultParams is the default parameter configuration for the bank module
func DefaultParams() Params {
	return Params{
		MaxDescCharacters: DefaultMaxDescCharacters,
	}
}

// Validate all bank module parameters
func (p Params) Validate() error {
	if err := validateMaxDescCharacters(p.MaxDescCharacters); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxDescCharacters, &p.MaxDescCharacters, validateMaxDescCharacters),
	}
}

func validateMaxDescCharacters(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid max desc characters: %d", v)
	}

	return nil
}
