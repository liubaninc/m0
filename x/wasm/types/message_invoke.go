package types

import (
	"fmt"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"

	utxotypes "github.com/liubaninc/m0/x/utxo/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
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

	var signers []sdk.AccAddress
	seen := map[string]bool{}
	signers = append(signers, creator)
	seen[m.Creator] = true

	contractInputsMap := map[string]bool{}
	contractInputs, err := ParseInputsFromExt(m.OutputsExt)
	if err != nil {
		panic(err)
	}
	for _, contractInput := range contractInputs {
		contractInputsMap[contractInput.Index()] = true
	}

	for _, input := range m.Inputs {
		if _, ok := contractInputsMap[input.Index()]; ok {
			continue // 忽略合约内部转账的地址
		}
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

func (m *MsgInvoke) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgInvoke) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	totalIn := sdk.NewCoins()
	inputsMap := map[string]bool{}
	for _, input := range m.Inputs {
		if err := input.ValidateBasic(); err != nil {
			return err
		}
		key := fmt.Sprintf("%s_%d_%d", input.RefTx, input.RefMsg, input.RefOffset)
		if _, ok := inputsMap[input.Index()]; ok {
			return sdkerrors.Wrapf(utxotypes.ErrUTXODuplicated, "invalid spend utxo %s (%s)", key, err)
		}
		inputsMap[input.Index()] = true
		totalIn = totalIn.Add(input.Amount)
	}

	totalOut := sdk.NewCoins()
	outputsMap := map[string]int{}
	for _, output := range m.Outputs {
		if err := output.ValidateBasic(); err != nil {
			return err
		}
		key := fmt.Sprintf("%s_%s", output.ToAddr, output.Amount)
		outputsMap[key]++
		totalOut = totalOut.Add(output.Amount)
	}

	if !totalIn.IsAllLTE(totalOut) {
		return sdkerrors.Wrapf(utxotypes.ErrInputOutputNotLitter, "inputs amount %v outputs amount %v", totalIn, totalOut)
	}

	if len(m.ContractRequests) == 0 {
		return sdkerrors.Wrapf(ErrUnexpected, "contract request is empty")
	}

	for _, cr := range m.ContractRequests {
		if err := kernel.ValidContractName(cr.ContractName); err != nil {
			return fmt.Errorf("contract name %v, error %v", cr.ContractName, err)
		}
	}

	if len(m.OutputsExt) != 0 {
		// 读集合 key列表
		inputExtKeys := map[string]bool{}
		for _, inputExt := range m.InputsExt {
			key := fmt.Sprintf("%s_%s", inputExt.Bucket, inputExt.Key)
			if _, ok := inputExtKeys[key]; ok {
				return sdkerrors.Wrapf(ErrUnexpected, "duplicate bucket key in input ext")
			}
			inputExtKeys[key] = true
		}
		// 写集合key 必须 存在读集合key list
		for _, outputExt := range m.OutputsExt {
			if outputExt.Bucket == TransientBucket {
				continue
			}
			key := fmt.Sprintf("%s_%s", outputExt.Bucket, outputExt.Key)
			if _, ok := inputExtKeys[key]; !ok {
				return sdkerrors.Wrapf(ErrUnexpected, "output ext not such key %s in input ext", key)
			}
			if outputExt.Value == nil {
				return sdkerrors.Wrapf(ErrUnexpected, "output ext value is null")
			}
		}

		contractInputs, err := ParseInputsFromExt(m.OutputsExt)
		if err != nil {
			return sdkerrors.Wrapf(ErrUnexpected, "parse contract inputs error - %v", err)
		}
		for _, contractInput := range contractInputs {
			key := fmt.Sprintf("%s_%d_%d", contractInput.RefTx, contractInput.RefMsg, contractInput.RefOffset)
			if _, ok := inputsMap[key]; !ok {
				return sdkerrors.Wrapf(ErrUnexpected, "missing contract input in inputs")
			}
		}

		contractOutputs, err := ParseOutputsFromExt(m.OutputsExt)
		if err != nil {
			return sdkerrors.Wrapf(ErrUnexpected, "parse contract outputs error - %v", err)
		}
		for _, request := range m.ContractRequests {
			for _, amount := range request.Amount {
				contractOutputs = append(contractOutputs, &utxotypes.Output{
					ToAddr: authtypes.NewModuleAddress(request.ContractName).String(),
					Amount: amount,
				})
			}
		}
		for _, contractOutput := range contractOutputs {
			key := fmt.Sprintf("%s_%s", contractOutput.ToAddr, contractOutput.Amount)
			if val, ok := outputsMap[key]; !ok || val < 1 {
				return sdkerrors.Wrapf(ErrUnexpected, "missing contract output in outputs")
			} else {
				outputsMap[key] = val - 1
			}
		}
	}
	return nil
}
