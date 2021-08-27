package types

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRecryptAccount{}

func NewMsgCreateRecryptAccount(creator string, pubkey string) *MsgCreateRecryptAccount {
	return &MsgCreateRecryptAccount{
		Creator: creator,
		Pubkey:  pubkey,
	}
}

func (msg *MsgCreateRecryptAccount) Route() string {
	return RouterKey
}

func (msg *MsgCreateRecryptAccount) Type() string {
	return "CreateRecryptAccount"
}

func (msg *MsgCreateRecryptAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRecryptAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRecryptAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Pubkey) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid pubkey (empty)")
	}

	if _, err := hex.DecodeString(msg.Pubkey); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid pubkey (%s)", err)
	}

	return nil
}
