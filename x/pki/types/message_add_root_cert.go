package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/pki/x509"
)

var _ sdk.Msg = &MsgAddRootCert{}

func NewMsgAddRootCert(creator string, certificate string) *MsgAddRootCert {
	return &MsgAddRootCert{
		Creator:     creator,
		Certificate: certificate,
	}
}

func (msg *MsgAddRootCert) Route() string {
	return RouterKey
}

func (msg *MsgAddRootCert) Type() string {
	return "AddRootCert"
}

func (msg *MsgAddRootCert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddRootCert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRootCert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Certificate) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid x509Cert (it cannot be empty)")
	}

	x509Certificate, err := x509.DecodeX509Certificate(msg.Certificate)
	if err != nil {
		return err
	}

	if !x509Certificate.IsSelfSigned() {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"certificate is not self-signed, so it cannot be used as a root certificate.")
	}

	return nil
}
