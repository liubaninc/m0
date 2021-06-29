package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendBlob{}

func NewMsgSendBlob(creator string, title string, content string) *MsgSendBlob {
	return &MsgSendBlob{
		Creator: creator,
		Title:   title,
		Content: content,
	}
}

func (msg *MsgSendBlob) Route() string {
	return RouterKey
}

func (msg *MsgSendBlob) Type() string {
	return "SendBlob"
}

func (msg *MsgSendBlob) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendBlob) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendBlob) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
