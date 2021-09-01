package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

func (m *InputExt) ValidateBasic() error {
	if len(m.RefTx) == 0 {
		return sdkerrors.Wrap(ErrUnexpected, "missing ref_tx in inputExt")
	}
	if m.RefMsg < 0 {
		return sdkerrors.Wrapf(ErrUnexpected, "negative ref_msg in inputExt")
	}
	if m.RefOffset < 0 {
		return sdkerrors.Wrapf(ErrUnexpected, "negative ref_offset in inputExt")
	}
	if len(m.Bucket) == 0 {
		return sdkerrors.Wrapf(ErrUnexpected, "missing bucket in inputExt")
	}
	if len(m.Key) == 0 {
		return sdkerrors.Wrapf(ErrUnexpected, "missing key in inputExt")
	}
	return nil
}

func (m *OutputExt) ValidateBasic() error {
	if len(m.Bucket) == 0 {
		return sdkerrors.Wrapf(ErrUnexpected, "missing bucket in outputExt")
	}
	if len(m.Key) == 0 {
		return sdkerrors.Wrapf(ErrUnexpected, "missing key in outputExt")
	}
	return nil
}

const (
	Normarl uint16 = iota
	Freeze
)
