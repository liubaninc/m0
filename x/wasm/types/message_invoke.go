package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
)

var _ sdk.Msg = &MsgInvoke{}

func NewMsgInvoke(creator string, inputs []*utxotypes.Input, outputs []*utxotypes.Output, inputsExt []*InputExt, outputsExt []*OutputExt, contractRequests []*InvokeRequest, desc string) *MsgInvoke {
	return &MsgInvoke{
		Creator:          creator,
		Desc:             desc,
		Inputs:           inputs,
		Outputs:          outputs,
		InputsExt:        inputsExt,
		OutputsExt:       outputsExt,
		ContractRequests: contractRequests,
	}
}

func (m *MsgInvoke) Route() string {
	return RouterKey
}

func (m *MsgInvoke) Type() string {
	return "Invoke"
}

func (m *MsgInvoke) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgInvoke) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgInvoke) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
