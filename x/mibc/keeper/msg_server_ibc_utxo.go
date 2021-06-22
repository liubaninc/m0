package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/liubaninc/m0/x/mibc/types"
)

func (k msgServer) SendIbcUTXO(goCtx context.Context, msg *types.MsgSendIbcUTXO) (*types.MsgSendIbcUTXOResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcUTXOPacketData

	packet.Receiver = msg.Receiver
	packet.Amount = msg.Amount

	// Transmit the packet
	err := k.TransmitIbcUTXOPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcUTXOResponse{}, nil
}
