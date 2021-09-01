package types

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"time"
)

// Default parameter namespace
const (
	// DefaultUnbondingTime reflects three weeks in seconds as the default unbonding time.
	DefaultUnbondingTime time.Duration = time.Hour * 24 * 7 * 3

	// Default max number of validators
	DefaultMaxValidators uint32 = 100
	// Default min number of validators
	DefaultMinValidators uint32 = 4

	// DefaultHistorical entries is 10000. Apps that don't use IBC can ignore this
	// value by not adding the staking module to the application module manager's
	// SetOrderBeginBlockers.
	DefaultHistoricalEntries uint32 = 10000
)

// Parameter store keys
var (
	KeyUnbondingTime     = []byte("UnbondingTime")
	KeyMaxValidators     = []byte("MaxValidators")
	KeyMinValidators     = []byte("MinValidators")
	KeyHistoricalEntries = []byte("HistoricalEntries")
)

// ParamKeyTable for poa module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params object
func NewParams(unbondingTime time.Duration, maxValidators uint32, minValidators uint32, historicalEntries uint32) Params {
	return Params{
		UnbondingTime:     unbondingTime,
		MaxValidators:     maxValidators,
		MinValidators:     minValidators,
		HistoricalEntries: historicalEntries,
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyUnbondingTime, &p.UnbondingTime, validateUnbondingTime),
		paramtypes.NewParamSetPair(KeyMaxValidators, &p.MaxValidators, validateMaxValidators),
		paramtypes.NewParamSetPair(KeyMinValidators, &p.MinValidators, validateMinValidators),
		paramtypes.NewParamSetPair(KeyHistoricalEntries, &p.HistoricalEntries, validateHistoricalEntries),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(DefaultUnbondingTime, DefaultMaxValidators, DefaultMinValidators, DefaultHistoricalEntries)
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

func validateUnbondingTime(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}

// Validate maxValidators param
func validateMaxValidators(i interface{}) error {
	v, ok := i.(uint32)
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
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("min validators must be positive: %d", v)
	}

	return nil
}

func validateHistoricalEntries(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
