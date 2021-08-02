package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"
)

var _ sdk.Msg = &MsgProposeDeployContract{}

func NewMsgProposeDeployContract(creator string, contractName string, contractCode []byte, initArgs string, approval []string) *MsgProposeDeployContract {
	return &MsgProposeDeployContract{
		Creator:      creator,
		ContractName: contractName,
		ContractCode: contractCode,
		InitArgs:     initArgs,
		Approval:     approval,
	}
}

func (msg *MsgProposeDeployContract) Route() string {
	return RouterKey
}

func (msg *MsgProposeDeployContract) Type() string {
	return "ProposeDeployContract"
}

func (msg *MsgProposeDeployContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProposeDeployContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeDeployContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.ContractName) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid contract name: it cannot be empty")
	}

	if err := kernel.ValidContractName(msg.ContractName); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "contract name %v, error %v", msg.ContractName, err)
	}

	if len(msg.ContractCode) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid contract code: it cannot be empty")
	}

	approval := map[string]bool{}
	for _, proposer := range msg.Approval {
		if approval[proposer] {
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "duplicate proposer")
		}
		approval[proposer] = true
	}

	if len(approval) < 2 {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid contract approval: (num must greater 1)")
	}

	return nil
}
