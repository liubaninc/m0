package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProposeAddAccountRequest{}

func NewMsgProposeAddAccount(creator string, address string, publickey string, roles []string) *MsgProposeAddAccountRequest {
	return &MsgProposeAddAccountRequest{
		Creator:   creator,
		Address:   address,
		PublicKey: publickey,
		Roles:     roles,
	}
}

func (msg *MsgProposeAddAccountRequest) Route() string {
	return RouterKey
}

func (msg *MsgProposeAddAccountRequest) Type() string {
	return "ProposeAddAccount"
}

func (msg *MsgProposeAddAccountRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	var signers []sdk.AccAddress
	seen := map[string]bool{}
	signers = append(signers, creator)
	seen[msg.Creator] = true
	return signers
}

func (msg *MsgProposeAddAccountRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeAddAccountRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address %s (%s)", msg.Creator, err)
	}
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}
	return nil

	if len(address) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid account address: it cannot be empty")

	}

	if len(msg.PublicKey) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "Invalid publicKey: it cannot be empty")

	}

	if err := Validates(msg.GetRoles()); err != nil {
		return err
	}

	return nil
}
