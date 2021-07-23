package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRevokeCert{}

func NewMsgRevokeCert(creator string, subject string, subjectKeyID string) *MsgRevokeCert {
	return &MsgRevokeCert{
		Creator:      creator,
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}
}

func (msg *MsgRevokeCert) Route() string {
	return RouterKey
}

func (msg *MsgRevokeCert) Type() string {
	return "RevokeCert"
}

func (msg *MsgRevokeCert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Subject) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Invalid Subject (it cannot be empty)")
	}

	if len(msg.SubjectKeyID) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Invalid SubjectKeyID (it cannot be empty)")
	}
	return nil
}
