package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"
)

var _ sdk.Msg = &MsgUndeploy{}

func NewMsgUndeploy(creator string, contractName string) *MsgUndeploy {
	return &MsgUndeploy{
		Creator:      creator,
		ContractName: contractName,
	}
}

func (msg *MsgUndeploy) Route() string {
	return RouterKey
}

func (msg *MsgUndeploy) Type() string {
	return "Undeploy"
}

func (msg *MsgUndeploy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUndeploy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUndeploy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := kernel.ValidContractName(msg.ContractName); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "contract name %v, error %v", msg.ContractName, err)
	}
	return nil
}
