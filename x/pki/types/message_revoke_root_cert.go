package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRevokeRootCert{}

func NewMsgRevokeRootCert(creator string, subject string, subjectKeyID string) *MsgRevokeRootCert {
	return &MsgRevokeRootCert{
		Creator:      creator,
		Subject:      subject,
		SubjectKeyID: subjectKeyID,
	}
}

func (msg *MsgRevokeRootCert) Route() string {
	return RouterKey
}

func (msg *MsgRevokeRootCert) Type() string {
	return "RevokeRootCert"
}

func (msg *MsgRevokeRootCert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeRootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeRootCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
