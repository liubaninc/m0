package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateShareEncryptStorage{}

func NewMsgCreateShareEncryptStorage(creator string, key string, shareKey string, envelope string, sharer string) *MsgCreateShareEncryptStorage {
	return &MsgCreateShareEncryptStorage{
		Creator:  creator,
		Key:      key,
		ShareKey: shareKey,
		Envelope: envelope,
		Sharer:   sharer,
	}
}

func (msg *MsgCreateShareEncryptStorage) Route() string {
	return RouterKey
}

func (msg *MsgCreateShareEncryptStorage) Type() string {
	return "CreateShareEncryptStorage"
}

func (msg *MsgCreateShareEncryptStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateShareEncryptStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateShareEncryptStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Key) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid key (empty)")
	}

	if len(msg.ShareKey) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid share key (empty)")
	}

	if len(msg.Envelope) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid envelope (empty)")
	}

	_, err = sdk.AccAddressFromBech32(msg.Sharer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sharer address (%s)", err)
	}
	return nil
}
