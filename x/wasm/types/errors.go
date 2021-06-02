package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/wasm module sentinel errors
var (
	ErrPreExec              = sdkerrors.Register(ModuleName, 1101, "a preexec error")
	ErrRWSet              = sdkerrors.Register(ModuleName, 1102, "a verify error")
	// this line is used by starport scaffolding # ibc/errors

	ErrUnexpected              = sdkerrors.Register(ModuleName, 1300, "a unexpected error")
)
