package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendIbcUTXO{}

func NewMsgSendIbcUTXO(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	receiver string,
	amount string,
) *MsgSendIbcUTXO {
	return &MsgSendIbcUTXO{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Receiver:         receiver,
		Amount:           amount,
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
	return nil
}
