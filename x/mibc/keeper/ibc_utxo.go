package keeper

import (
	"context"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/liubaninc/m0/x/mibc/types"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"strings"
)

// TransmitIbcUTXOPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIbcUTXOPacket(
	ctx sdk.Context,
	packetData types.IbcUTXOPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvIbcUTXOPacket processes packet reception
func (k Keeper) OnRecvIbcUTXOPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcUTXOPacketData) (packetAck types.IbcUTXOPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic
	msgOffset := int32(ctx.Context().Value("msg-index").(int))
	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	escrowAddress := types.GetEscrowAddress(packet.GetDestPort(), packet.GetDestChannel())
	var outputs []*utxotypes.Output
	totalNeeded := sdk.NewCoins()
	for _, output := range data.Outputs {
		if types.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), output.Denom) {
			// sender chain is not the source, unescrow tokens

			// remove prefix added by sender chain
			voucherPrefix := types.GetDenomPrefix(packet.GetSourcePort(), packet.GetSourceChannel())
			unprefixedDenom := output.Denom[len(voucherPrefix):]

			// coin denomination used in sending from the escrow address
			denom := unprefixedDenom

			// The denomination used to send the coins is either the native denom or the hash of the path
			// if the denomination is not native.
			denomTrace := types.ParseDenomTrace(unprefixedDenom)
			if denomTrace.Path != "" {
				denom = denomTrace.IBCDenom()
			}
			token := sdk.NewCoin(denom, output.Amount)

			outputs = append(outputs, &utxotypes.Output{
				ToAddr: output.Addr,
				Amount: token,
				FrozenHeight: output.FrozenHeight,
			})
			totalNeeded = totalNeeded.Add(token)
		} else {
			// sender chain is the source, mint vouchers

			// since SendPacket did not prefix the denomination, we must prefix denomination here
			sourcePrefix := types.GetDenomPrefix(packet.GetDestPort(), packet.GetDestChannel())
			// NOTE: sourcePrefix contains the trailing "/"
			prefixedDenom := sourcePrefix + output.Denom
			// construct the denomination trace from the full raw denomination
			denomTrace := types.ParseDenomTrace(prefixedDenom)

			traceHash := denomTrace.Hash()
			if !k.HasDenomTrace(ctx, traceHash) {
				k.SetDenomTrace(ctx, denomTrace)
			}

			voucherDenom := denomTrace.IBCDenom()
			voucher := sdk.NewCoin(voucherDenom, output.Amount)

			outputs = append(outputs, &utxotypes.Output{
				ToAddr: output.Addr,
				Amount: voucher,
				FrozenHeight: output.FrozenHeight,
			})
		}
	}

	var inputs []*utxotypes.Input
	if !totalNeeded.IsZero() {
		res, err := k.utxoKeeper.Input(context.Background(), &utxotypes.QueryInputRequest{
			Address: escrowAddress.String(),
			Amounts: totalNeeded.String(),
			Lock:    10,
		})
		if err != nil {
			return packetAck, err
		}
		inputs = append(inputs, res.Inputs...)
		changeCoins := res.Amount.Sub(totalNeeded)
		for _, changeCoin := range changeCoins {
			outputs = append(outputs, &utxotypes.Output{
				ToAddr: escrowAddress.String(),
				Amount: changeCoin,
			})
		}
	}

	if err := k.utxoKeeper.Transfer(ctx, hash, msgOffset, escrowAddress.String(), inputs, outputs); err != nil{
		return packetAck, err
	}

	k.AppendItx(ctx, types.Itx{
		Creator: strings.Join([]string{packet.SourceChannel, packet.SourcePort, data.Creator}, "-"),
		SourceHash: data.Hash,
		DestinationHash: hash,
		Source: false,
	})

	packetAck.Hash = hash
	return packetAck, nil
}

// OnAcknowledgementIbcUTXOPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIbcUTXOPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcUTXOPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		k.AppendItx(ctx, types.Itx{
			Creator: strings.Join([]string{packet.SourceChannel, packet.SourcePort, data.Creator}, "-"),
			SourceHash: data.Hash,
			Source: false,
			Log: dispatchedAck.Error,
		})
		return k.refundPacketToken(ctx, packet, data)
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IbcUTXOPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic
		k.AppendItx(ctx, types.Itx{
			Creator: strings.Join([]string{packet.SourceChannel, packet.SourcePort, data.Creator}, "-"),
			SourceHash: data.Hash,
			DestinationHash: packetAck.Hash,
			Source: false,
		})

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIbcUTXOPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIbcUTXOPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcUTXOPacketData) error {
	// TODO: packet timeout logic
	k.AppendItx(ctx, types.Itx{
		Creator: strings.Join([]string{packet.SourceChannel, packet.SourcePort, data.Creator}, "-"),
		SourceHash: data.Hash,
		Source: false,
		Log: "timeout",
	})
	return k.refundPacketToken(ctx, packet, data)
}

// refundPacketToken will unescrow and send back the tokens back to sender
// if the sending chain was the source chain. Otherwise, the sent tokens
// were burnt in the original send so new tokens are minted and sent to
// the sending address.
func (k Keeper) refundPacketToken(ctx sdk.Context, packet channeltypes.Packet, data types.IbcUTXOPacketData) error {

	return nil
}