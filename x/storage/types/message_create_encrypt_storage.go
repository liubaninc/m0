package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateEncryptStorage{}

func NewMsgCreateEncryptStorage(creator string, key string, md5 string, encrypted string, envelope string) *MsgCreateEncryptStorage {
	return &MsgCreateEncryptStorage{
		Creator:   creator,
		Key:       key,
		Md5:       md5,
		Encrypted: encrypted,
		Envelope:  envelope,
	}
}

func (msg *MsgCreateEncryptStorage) Route() string {
	return RouterKey
}

func (msg *MsgCreateEncryptStorage) Type() string {
	return "CreateEncryptStorage"
}

func (msg *MsgCreateEncryptStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEncryptStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEncryptStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Key) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid key (empty)")
	}

	if len(msg.Md5) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid md5 (empty)")
	}

	if len(msg.Encrypted) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid encrypted (empty)")
	}

	if len(msg.Envelope) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid envelope (empty)")
	}

	return nil
}
