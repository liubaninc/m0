package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (m *Input) Index() string {
	addr, _ := sdk.AccAddressFromBech32(m.FromAddr)
	return fmt.Sprintf("%X_%s_%s_%d_%d", addr.Bytes(), m.Amount.Denom, strings.ToUpper(m.RefTx), m.RefMsg, m.RefOffset)
}

func (m *Input) ValidateBasic() error {
	if len(m.FromAddr) == 0 {
		return sdkerrors.Wrap(ErrUnexpected, "missing address in input")
	}
	if len(m.RefTx) == 0 {
		return sdkerrors.Wrap(ErrUnexpected, "missing ref_tx in input")
	}
	if m.RefMsg < 0 {
		return sdkerrors.Wrapf(ErrUnexpected, "negative ref_msg in input")
	}
	if m.RefOffset < 0 {
		return sdkerrors.Wrapf(ErrUnexpected, "negative ref_offset in input")
	}
	_, err := sdk.AccAddressFromBech32(m.FromAddr)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid input address %s (%s)", m.FromAddr, err)
	}
	if !m.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}
	if !m.Amount.IsPositive() {
		return sdkerrors.Wrapf(ErrNegativeAmount, "invalid input amount %s (%s)", m.Amount, err)
	}
	return nil
}

func (m *Output) ValidateBasic() error {
	if len(m.ToAddr) == 0 {
		return sdkerrors.Wrap(ErrUnexpected, "missing address in output")
	}
	_, err := sdk.AccAddressFromBech32(m.ToAddr)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid output address %s (%s)", m.ToAddr, err)
	}
	if !m.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}
	if !m.Amount.IsPositive() {
		return sdkerrors.Wrapf(ErrNegativeAmount, "invalid output amount %s (%s)", m.Amount, err)
	}
	return nil
}
