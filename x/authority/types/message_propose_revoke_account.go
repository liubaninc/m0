package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgProposeRevokeAccountRequest{}

func NewMsgProposeRevokeAccountRequest(creator string, address string) *MsgProposeRevokeAccountRequest {
	return &MsgProposeRevokeAccountRequest{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgProposeRevokeAccountRequest) Route() string {
	return RouterKey
}

func (msg *MsgProposeRevokeAccountRequest) Type() string {
	return "ProposeRevokeAccount"
}

func (msg *MsgProposeRevokeAccountRequest) GetSigners() []sdk.AccAddress {
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

func (msg *MsgProposeRevokeAccountRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeRevokeAccountRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid account address (%s)", err)
	}
	return nil
}
