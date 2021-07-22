package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgApproveRevokeAccountRequest{}

func NewMsgApproveRevokeAccountRequest(creator string, address string) *MsgApproveRevokeAccountRequest {
	return &MsgApproveRevokeAccountRequest{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgApproveRevokeAccountRequest) Route() string {
	return RouterKey
}

func (msg *MsgApproveRevokeAccountRequest) Type() string {
	return "ApproveRevokeAccount"
}

func (msg *MsgApproveRevokeAccountRequest) GetSigners() []sdk.AccAddress {
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

func (msg *MsgApproveRevokeAccountRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveRevokeAccountRequest) ValidateBasic() error {
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
