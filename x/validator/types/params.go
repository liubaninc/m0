package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Default parameter namespace
const (
	// Default max number of validators
	DefaultMaxValidators uint16 = 100
	// Default min number of validators
	DefaultMinValidators uint16 = 1
)

// Parameter store keys
var (
	KeyMaxValidators = []byte("MaxValidators")
	KeyMinValidators = []byte("MinValidators")
)

// ParamKeyTable for poa module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Params - used for initializing default parameter for poa at genesis
type Params struct {
	MaxValidators uint16 `json:"max_validators"`
	MinValidators uint16 `json:"min_validators"`
}

// NewParams creates a new Params object
func NewParams(maxValidators uint16, minValidators uint16) Params {
	return Params{
		MaxValidators: maxValidators,
		MinValidators: minValidators,
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	return fmt.Sprintf("Max validators: %d, Min validators: %d percents", p.MaxValidators, p.MinValidators)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxValidators, &p.MaxValidators, validateMaxValidators),
		paramtypes.NewParamSetPair(KeyMinValidators, &p.MinValidators, validateMinValidators),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(DefaultMaxValidators, DefaultMinValidators)
}

// Validate a set of params
func (p Params) Validate() error {
	if err := validateMaxValidators(p.MaxValidators); err != nil {
		return err
	}
	if err := validateMinValidators(p.MinValidators); err != nil {
		return err
	}
	return nil
}

// Validate maxValidators param
func validateMaxValidators(i interface{}) error {
	v, ok := i.(uint16)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}

	return nil
}

// Quorum must be a percentage
func validateMinValidators(i interface{}) error {
	v, ok := i.(uint16)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("min validators must be positive: %d", v)
	}

	return nil
}
