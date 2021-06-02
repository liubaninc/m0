package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
)

var _ sdk.Msg = &MsgUpgrade{}

func NewMsgUpgrade(creator string, contractName string, contractCode []byte, limits []*xmodel.ResourceLimit, inputs []*utxotypes.Input, outputs []*utxotypes.Output, inputsExt []*InputExt, outputsExt []*OutputExt, desc string) *MsgUpgrade {
	return &MsgUpgrade{
		Creator:      creator,
		Desc:         desc,
		Inputs: inputs,
		Outputs: outputs,
		InputsExt:    inputsExt,
		OutputsExt:   outputsExt,
		ContractName: contractName,
		ContractCode: contractCode,
		ResourceLimits: limits,
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

	var signers []sdk.AccAddress
	seen := map[string]bool{}
	signers = append(signers, creator)
	seen[m.Creator] = true
	return signers
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

func (m *MsgUpgrade) ConvertInvokeRequest() *InvokeRequest {
	args, _ := json.Marshal(map[string][]byte{
		"contract_name": []byte(m.ContractName),
		"contract_code": m.ContractCode,
	})
	return &InvokeRequest{
		ModuleName: "kernel",
		ContractName: "",
		MethodName: "Upgrade",
		Args: string(args),
		ResourceLimits: m.ResourceLimits,
	}
}