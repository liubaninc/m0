package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSend{}

func NewMsgSend(creator string, inputs []*Input, outputs []*Output, desc string) *MsgSend {
	return &MsgSend{
		Creator: creator,
		Inputs:  inputs,
		Outputs: outputs,
		Desc:    desc,
	}
}

func (m *MsgSend) Route() string {
	return RouterKey
}

func (m *MsgSend) Type() string {
	return "Send"
}

func (m *MsgSend) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		panic(err)
	}

	var signers []sdk.AccAddress
	seen := map[string]bool{}
	signers = append(signers, creator)
	seen[m.Creator] = true

	for _, input := range m.Inputs {
		if addrStr := input.FromAddr; !seen[addrStr] {
			addr, err := sdk.AccAddressFromBech32(input.FromAddr)
			if err != nil {
				panic(err)
			}
			signers = append(signers, addr)
			seen[addrStr] = true
		}
	}

	return signers
}

func (m *MsgSend) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgSend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address %s (%s)", m.Creator, err)
	}

	totalIn := sdk.NewCoins()
	totalOut := sdk.NewCoins()

	inputsMap := map[string]bool{}
	for _, input := range m.Inputs {
		if err := input.ValidateBasic(); err != nil {
			return err
		}
		key := fmt.Sprintf("%s_%d_%d", input.RefTx, input.RefMsg, input.RefOffset)
		if _, ok := inputsMap[key]; ok {
			return sdkerrors.Wrapf(ErrUTXODuplicated, "invalid spend utxo %s (%s)", key, err)
		}
		inputsMap[key] = true
		totalIn = totalIn.Add(input.Amount)
	}

	for _, output := range m.Outputs {
		if err := output.ValidateBasic(); err != nil {
			return err
		}
		totalOut = totalOut.Add(output.Amount)
	}

	if !totalIn.IsEqual(totalOut) {
		return sdkerrors.Wrapf(ErrInputOutputNotEqual, "inputs amount %v should equal outputs amount %v", totalIn, totalOut)
	}
	return nil
}
