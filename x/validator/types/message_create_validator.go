package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateValidator{}

func NewMsgCreateValidator(creator string, pubKey string, description *Description) *MsgCreateValidator {
	return &MsgCreateValidator{
		Creator:     creator,
		PubKey:      pubKey,
		Description: description,
	}
}

func (msg *MsgCreateValidator) Route() string {
	return RouterKey
}

func (msg *MsgCreateValidator) Type() string {
	return "CreateValidator"
}

func (msg *MsgCreateValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, msg.PubKey); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "invalid validator public key: %v", err)
	}

	if err := msg.Description.ValidateBasic(); err != nil {
		return err
	}
	return nil
}
