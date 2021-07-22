package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetPermission{}

func NewMsgSetPermission(creator string, address string, perms []string) *MsgSetPermission {
	return &MsgSetPermission{
		Creator: creator,
		Address: address,
		Perms:   perms,
	}
}

func (msg *MsgSetPermission) Route() string {
	return RouterKey
}

func (msg *MsgSetPermission) Type() string {
	return "SetPermission"
}

func (msg *MsgSetPermission) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetPermission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPermission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
