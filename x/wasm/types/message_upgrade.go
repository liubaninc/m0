package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpgrade{}

func NewMsgUpgrade(creator string, contractName string, contractCode []byte, inputsExt []*InputExt, outputsExt []*OutputExt, desc string) *MsgUpgrade {
	return &MsgUpgrade{
		Creator:      creator,
		Desc:         desc,
		InputsExt:    inputsExt,
		OutputsExt:   outputsExt,
		ContractName: contractName,
		ContractCode: contractCode,
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
	}
	return nil
}

func (m *MsgUpgrade) ConvertInvokeRequest() *InvokeRequest {
	return &InvokeRequest{
		ModuleName: "kernel",
		ContractName: "",
		MethodName: "Upgrade",
		Args: map[string][]byte{
			"contract_name": []byte(m.ContractName),
			"contract_code": m.ContractCode,
		},
	}
}