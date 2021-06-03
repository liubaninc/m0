package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDestroy{}

func NewMsgDestroy(creator string, inputs []*Input, outputs []*Output, desc string) *MsgDestroy {
	return &MsgDestroy{
		Creator: creator,
		Inputs:  inputs,
		Outputs: outputs,
		Desc:    desc,
	}
}

func (m *MsgDestroy) Route() string {
	return RouterKey
}

func (m *MsgDestroy) Type() string {
	return "Destroy"
}

func (m *MsgDestroy) GetSigners() []sdk.AccAddress {
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

func (m *MsgDestroy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDestroy) ValidateBasic() error {
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
			return sdkerrors.Wrapf(ErrUTXODuplicated, "invalid utxo %s (%s)", key, err)
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

	if !totalIn.IsAllGTE(totalOut) {
		return sdkerrors.Wrapf(ErrInputOutputNotGreater, "inputs amount %v outputs amount %v", totalIn, totalOut)
	}

	return nil
}
