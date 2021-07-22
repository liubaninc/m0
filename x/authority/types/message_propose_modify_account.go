package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProposeModifyAccountRequest{}

func NewMsgProposeModifyAccountRequest(creator string, address string, roles []string) *MsgProposeModifyAccountRequest {
	return &MsgProposeModifyAccountRequest{
		Creator: creator,
		Address: address,
		Roles:   roles,
	}
}

func (msg *MsgProposeModifyAccountRequest) Route() string {
	return RouterKey
}

func (msg *MsgProposeModifyAccountRequest) Type() string {
	return "ProposeModifyAccount"
}

func (msg *MsgProposeModifyAccountRequest) GetSigners() []sdk.AccAddress {
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

func (msg *MsgProposeModifyAccountRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeModifyAccountRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address %s (%s)", msg.Creator, err)
	}

	if len(msg.Address) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address: it cannot be empty")
	}
	if err := Validates(msg.GetRoles()); err != nil {
		return err
	}

	return nil
}
