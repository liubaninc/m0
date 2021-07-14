package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUnfreeze{}

func NewMsgUnfreeze(creator string, contractName string) *MsgUnfreeze {
	return &MsgUnfreeze{
		Creator:      creator,
		ContractName: contractName,
	}
}

func (msg *MsgUnfreeze) Route() string {
	return RouterKey
}

func (msg *MsgUnfreeze) Type() string {
	return "Unfreeze"
}

func (msg *MsgUnfreeze) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnfreeze) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnfreeze) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
