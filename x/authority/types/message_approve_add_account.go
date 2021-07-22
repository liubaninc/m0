package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgApproveAddAccountRequest{}

func NewMsgApproveAddAccountRequest(creator string, address string) *MsgApproveAddAccountRequest {
	return &MsgApproveAddAccountRequest{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgApproveAddAccountRequest) Route() string {
	return RouterKey
}

func (msg *MsgApproveAddAccountRequest) Type() string {
	return "ApproveAddAccount"
}

func (msg *MsgApproveAddAccountRequest) GetSigners() []sdk.AccAddress {
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

func (msg *MsgApproveAddAccountRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveAddAccountRequest) ValidateBasic() error {
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
