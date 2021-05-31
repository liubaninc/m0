package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/liubaninc/m0/x/wasm/xmodel"
)

var _ sdk.Msg = &MsgDeploy{}

func NewMsgDeploy(creator string, contractName string, contractCode []byte, contractDesc *xmodel.WasmCodeDesc, method string, args map[string][]byte, inputsExt []*InputExt, outputsExt []*OutputExt, desc string) *MsgDeploy {
	return &MsgDeploy{
		Creator:      creator,
		Desc:         desc,
		InputsExt:    inputsExt,
		OutputsExt:   outputsExt,
		ContractName: contractName,
		ContractCode: contractCode,
		ContractDesc: contractDesc,
		MethodName:   method,
		Args:         args,
	}
}

func (m *MsgDeploy) Route() string {
	return RouterKey
}

func (m *MsgDeploy) Type() string {
	return "Deploy"
}

func (m *MsgDeploy) GetSigners() []sdk.AccAddress {
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

func (m *MsgDeploy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDeploy) ValidateBasic() error {
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

func (m *MsgDeploy) ConvertInvokeRequest() *InvokeRequest {
	desc, err := proto.Marshal(m.ContractDesc)
	if err != nil {
		panic(err)
	}
	args, err := json.Marshal(m.Args)
	if err != nil {
		panic(err)
	}
	initArgs, err := json.Marshal(map[string][]byte{
		m.MethodName: args,
	})
	if err != nil {
		panic(err)
	}
	return &InvokeRequest{
		ModuleName: "kernel",
		ContractName: "",
		MethodName: "Deploy",
		Args: map[string][]byte{
			"contract_name": []byte(m.ContractName),
			"contract_code": m.ContractCode,
			"contract_desc": desc,
			"init_args": initArgs,
		},
	}
}