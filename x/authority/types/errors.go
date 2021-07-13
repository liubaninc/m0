package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/authority module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	// this line is used by starport scaffolding # ibc/errors
	ErrUnauthorized         = sdkerrors.Register(ModuleName, 1101, "authorization error")
	ErrAccountAlreadyExists = sdkerrors.Register(ModuleName, 1102, "account already exists on the ledger")
)
