package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePeerID{}

func NewMsgCreatePeerID(creator string, index string, certIssuer string, certSerialNum string) *MsgCreatePeerID {
	return &MsgCreatePeerID{
		Creator:       creator,
		Index:         index,
		CertIssuer:    certIssuer,
		CertSerialNum: certSerialNum,
	}
}

func (msg *MsgCreatePeerID) Route() string {
	return RouterKey
}

func (msg *MsgCreatePeerID) Type() string {
	return "CreatePeerID"
}

func (msg *MsgCreatePeerID) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePeerID) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePeerID) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.CertIssuer) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid cert issuer: it cannot be empty")
	}

	if len(msg.CertSerialNum) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid cert serialNum: it cannot be empty")
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePeerID{}

func NewMsgUpdatePeerID(creator string, index string, certIssuer string, certSerialNum string) *MsgUpdatePeerID {
	return &MsgUpdatePeerID{
		Creator:       creator,
		Index:         index,
		CertIssuer:    certIssuer,
		CertSerialNum: certSerialNum,
	}
}

func (msg *MsgUpdatePeerID) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePeerID) Type() string {
	return "UpdatePeerID"
}

func (msg *MsgUpdatePeerID) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePeerID) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePeerID) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.CertIssuer) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid cert issuer: it cannot be empty")
	}

	if len(msg.CertSerialNum) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid cert serialNum: it cannot be empty")
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePeerID{}

func NewMsgDeletePeerID(creator string, index string) *MsgDeletePeerID {
	return &MsgDeletePeerID{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeletePeerID) Route() string {
	return RouterKey
}

func (msg *MsgDeletePeerID) Type() string {
	return "DeletePeerID"
}

func (msg *MsgDeletePeerID) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePeerID) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePeerID) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
