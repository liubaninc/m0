package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctransfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
)

var _ sdk.Msg = &MsgSendIbcUTXO{}

func NewMsgSendIbcUTXO(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	inputs []*utxotypes.Input,
	outputs []*utxotypes.Output,
	desc string,
) *MsgSendIbcUTXO {
	return &MsgSendIbcUTXO{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Inputs:           inputs,
		Outputs:          outputs,
		Desc:             desc,
	}
}

func (msg *MsgSendIbcUTXO) Route() string {
	return RouterKey
}

func (msg *MsgSendIbcUTXO) Type() string {
	return "SendIbcUTXO"
}

func (msg *MsgSendIbcUTXO) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	var signers []sdk.AccAddress
	seen := map[string]bool{}
	signers = append(signers, sender)
	seen[msg.Sender] = true

	for _, input := range msg.Inputs {
		if addrStr := input.FromAddr; !seen[addrStr] {
			addr, err := sdk.AccAddressFromBech32(input.FromAddr)
			if err != nil {
				panic(err)
			}
			signers = append(signers, addr)
			seen[addrStr] = true
		}
	}

	return []sdk.AccAddress{sender}
}

func (msg *MsgSendIbcUTXO) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendIbcUTXO) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	if err := host.PortIdentifierValidator(msg.Port); err != nil {
		return sdkerrors.Wrap(err, "invalid source port ID")
	}
	if err := host.ChannelIdentifierValidator(msg.ChannelID); err != nil {
		return sdkerrors.Wrap(err, "invalid source channel ID")
	}

	totalIn := sdk.NewCoins()
	totalOut := sdk.NewCoins()

	inputsMap := map[string]bool{}
	for _, input := range msg.Inputs {
		if err := input.ValidateBasic(); err != nil {
			return err
		}
		if err := ibctransfertypes.ValidateIBCDenom(input.Amount.Denom); err != nil {
			return err
		}
		key := fmt.Sprintf("%s_%d_%d", input.RefTx, input.RefMsg, input.RefOffset)
		if _, ok := inputsMap[key]; ok {
			return sdkerrors.Wrapf(utxotypes.ErrUTXODuplicated, "invalid spend utxo %s (%s)", key, err)
		}
		inputsMap[key] = true
		totalIn = totalIn.Add(input.Amount)
	}

	for _, output := range msg.Outputs {
		if err := output.ValidateBasic(); err != nil {
			return err
		}
		if err := ibctransfertypes.ValidateIBCDenom(output.Amount.Denom); err != nil {
			return err
		}
		totalOut = totalOut.Add(output.Amount)
	}

	if !totalIn.IsEqual(totalOut) {
		return sdkerrors.Wrapf(utxotypes.ErrInputOutputNotEqual, "inputs amount %v should equal outputs amount %v", totalIn, totalOut)
	}
	return nil
}
