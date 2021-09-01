package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFreezeCert{}

func NewMsgFreezeCert(creator string, subject string, subjectKeyID string) *MsgFreezeCert {
	return &MsgFreezeCert{
		Creator:      creator,
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}
}

func (msg *MsgFreezeCert) Route() string {
	return RouterKey
}

func (msg *MsgFreezeCert) Type() string {
	return "FreezeCert"
}

func (msg *MsgFreezeCert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFreezeCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFreezeCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
