package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgApproveDeployContract{}

func NewMsgApproveDeployContract(creator string, index string) *MsgApproveDeployContract {
	return &MsgApproveDeployContract{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgApproveDeployContract) Route() string {
	return RouterKey
}

func (msg *MsgApproveDeployContract) Type() string {
	return "ApproveDeployContract"
}

func (msg *MsgApproveDeployContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveDeployContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveDeployContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
