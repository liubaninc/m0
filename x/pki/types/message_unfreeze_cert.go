package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUnfreezeCert{}

func NewMsgUnfreezeCert(creator string, subject string, subjectKeyID string) *MsgUnfreezeCert {
	return &MsgUnfreezeCert{
		Creator:      creator,
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}
}

func (msg *MsgUnfreezeCert) Route() string {
	return RouterKey
}

func (msg *MsgUnfreezeCert) Type() string {
	return "UnfreezeCert"
}

func (msg *MsgUnfreezeCert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnfreezeCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnfreezeCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
