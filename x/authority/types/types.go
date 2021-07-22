package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	UTXO      string = "utxo"
	WASM      string = "wasm"
	MIBC      string = "mibc"
	AUTHORITY string = "authority"
)

var Roles = []string{UTXO, WASM, MIBC, AUTHORITY}

func Validate(role string) error {
	for _, r := range Roles {
		if role == r {
			return nil
		}
	}
	return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Invalid Account Role: %v. Supported roles: [%v]", role, Roles)

}

// Validate checks for errors on the account roles.
func Validates(roles []string) error {
	for _, role := range roles {
		if err := Validate(role); err != nil {
			return err
		}
	}
	return nil
}
