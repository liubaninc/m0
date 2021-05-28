package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
)

var _ sdk.Msg = &MsgUpgrade{}

func NewMsgUpgrade(creator string, inputs []*utxotypes.Input, outputs []*utxotypes.Output, inputsExt []*InputExt, outputsExt []*OutputExt, contractRequests []*InvokeRequest, desc string) *MsgUpgrade {
	return &MsgUpgrade{
		Creator:          creator,
		Desc:             desc,
		Inputs:           inputs,
		Outputs:          outputs,
		InputsExt:        inputsExt,
		OutputsExt:       outputsExt,
		ContractRequests: contractRequests,
	}
}

func (m *MsgUpgrade) Route() string {
	return RouterKey
}

func (m *MsgUpgrade) Type() string {
	return "Upgrade"
}

func (m *MsgUpgrade) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (m *MsgUpgrade) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgUpgrade) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(m.ContractRequests) == 0 {

	}

	totalIn := sdk.NewCoins()
	for _, input := range m.Inputs {
		if err := input.ValidateBasic(); err != nil {
			return err
		}
		totalIn = totalIn.Add(input.Amount)
	}

	totalOut := sdk.NewCoins()
	for _, output := range m.Outputs {
		if err := output.ValidateBasic(); err != nil {
			return err
		}
		totalOut = totalOut.Add(output.Amount)
	}

	if !totalIn.IsEqual(totalOut) {
		return sdkerrors.Wrapf(utxotypes.ErrInputOutputNotEqual, "inputs amount %v outputs amount %v", totalIn, totalOut)
	}
	return nil
}
