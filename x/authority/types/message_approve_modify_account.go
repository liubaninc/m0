package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgApproveModifyAccountRequest{}

func NewMsgApproveModifyAccountRequest(creator string, address string) *MsgApproveModifyAccountRequest {
	return &MsgApproveModifyAccountRequest{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgApproveModifyAccountRequest) Route() string {
	return RouterKey
}

func (msg *MsgApproveModifyAccountRequest) Type() string {
	return "ApproveModifyAccount"
}

func (msg *MsgApproveModifyAccountRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveModifyAccountRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveModifyAccountRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
