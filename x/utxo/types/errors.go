package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/utxo module sentinel errors
var (
	// this line is used by starport scaffolding # ibc/errors
	ErrNoEnoughUTXO          = sdkerrors.Register(ModuleName, 1100, "no enough utxo to get")
	ErrUTXONotFound          = sdkerrors.Register(ModuleName, 1101, "utxo can not be found")
	ErrUTXONotMismatch       = sdkerrors.Register(ModuleName, 1102, "utxo is not mismatch")
	ErrUTXOFrozen            = sdkerrors.Register(ModuleName, 1103, "utxo is still frozen")
	ErrInputOutputNotEqual   = sdkerrors.Register(ModuleName, 1104, "input's amount is not equal to output's")
	ErrInputOutputNotGreater = sdkerrors.Register(ModuleName, 1105, "input's amount should be greater to output's")
	ErrInputOutputNotLitter  = sdkerrors.Register(ModuleName, 1116, "input's amount should be litter to output's")
	ErrNegativeAmount        = sdkerrors.Register(ModuleName, 1107, "amount can not be negative number")
	ErrDescSizeLimitExceeded = sdkerrors.Register(ModuleName, 1108, "desc size limit exceeded")
	ErrUTXODuplicated        = sdkerrors.Register(ModuleName, 1109, "found duplicated utxo in input")

	ErrInvalidIssuer = sdkerrors.Register(ModuleName, 1200, "the issuer is invalid")

	ErrUnexpected = sdkerrors.Register(ModuleName, 1300, "a unexpected error")
)
