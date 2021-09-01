package keeper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/liubaninc/m0/x/mibc/types"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
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
	msgOffset := int32(ctx.Context().Value(baseapp.KeyMsgOffset).(int))
	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	if ctx.IsCheckTx() || ctx.IsReCheckTx() {
		return packetAck, nil
	}
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
				ToAddr:       output.Addr,
				Amount:       token,
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
				ToAddr:       output.Addr,
				Amount:       voucher,
				FrozenHeight: output.FrozenHeight,
			})
		}
	}

	var inputs []*utxotypes.Input
	if !totalNeeded.IsZero() {
		sInput, total, err := k.utxoKeeper.SelectUtxos(ctx, escrowAddress, totalNeeded, 10)
		if err != nil {
			return packetAck, err
		}
		inputs = append(inputs, sInput...)
		changeCoins := total.Sub(totalNeeded)
		for _, changeCoin := range changeCoins {
			outputs = append(outputs, &utxotypes.Output{
				ToAddr: escrowAddress.String(),
				Amount: changeCoin,
			})
		}
	}

	attrs, err := k.utxoKeeper.Transfer(ctx, hash, msgOffset, escrowAddress.String(), inputs, outputs)
	if err != nil {
		return packetAck, err
	}
	_ = attrs

	k.AppendItx(ctx, types.Itx{
		Creator:         strings.Join([]string{packet.SourceChannel, packet.SourcePort, data.Creator}, "-"),
		SourceHash:      data.Hash,
		DestinationHash: hash,
	})

	packetAck.Hash = hash
	return packetAck, nil
}

// OnAcknowledgementIbcUTXOPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIbcUTXOPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcUTXOPacketData, ack channeltypes.Acknowledgement) error {
	msgOffset := int32(ctx.Context().Value(baseapp.KeyMsgOffset).(int))
	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	if ctx.IsCheckTx() || ctx.IsReCheckTx() {
		return nil
	}
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		k.AppendItx(ctx, types.Itx{
			Creator:         data.Creator,
			SourceHash:      data.Hash,
			DestinationHash: hash,
			Log:             "ack: " + dispatchedAck.Error,
		})
		return k.refundPacketToken(ctx, packet, data, hash, msgOffset)
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IbcUTXOPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic
		k.AppendItx(ctx, types.Itx{
			Creator:         data.Creator,
			SourceHash:      data.Hash,
			DestinationHash: packetAck.Hash,
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
	msgOffset := int32(ctx.Context().Value(baseapp.KeyMsgOffset).(int))
	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	if ctx.IsCheckTx() || ctx.IsReCheckTx() {
		return nil
	}
	k.AppendItx(ctx, types.Itx{
		Creator:         data.Creator,
		SourceHash:      data.Hash,
		DestinationHash: hash,
		Log:             "ack: " + "timeout",
	})
	return k.refundPacketToken(ctx, packet, data, hash, msgOffset)
}

// refundPacketToken will unescrow and send back the tokens back to sender
// if the sending chain was the source chain. Otherwise, the sent tokens
// were burnt in the original send so new tokens are minted and sent to
// the sending address.
func (k Keeper) refundPacketToken(ctx sdk.Context, packet channeltypes.Packet, data types.IbcUTXOPacketData, hash string, msgOffset int32) error {
	var outputs []*utxotypes.Output
	escrowAddress := types.GetEscrowAddress(packet.SourcePort, packet.SourceChannel)
	totalNeeded := sdk.NewCoins()
	for _, output := range data.Outputs {
		// parse the denomination from the full denom path
		trace := types.ParseDenomTrace(output.Denom)

		token := sdk.NewCoin(trace.IBCDenom(), output.Amount)
		if types.SenderChainIsSource(packet.SourcePort, packet.SourceChannel, output.Denom) {
			// unescrow tokens back to sender
			outputs = append(outputs, &utxotypes.Output{
				ToAddr: data.Creator,
				Amount: token,
			})
			totalNeeded = totalNeeded.Add(token)
		} else {
			// mint vouchers back to sender
			outputs = append(outputs, &utxotypes.Output{
				ToAddr: data.Creator,
				Amount: token,
			})
		}
	}

	var inputs []*utxotypes.Input
	if !totalNeeded.IsZero() {
		sInputs, total, err := k.utxoKeeper.SelectUtxos(ctx, escrowAddress, totalNeeded, 10)
		if err != nil {
			return err
		}
		inputs = append(inputs, sInputs...)
		changeCoins := total.Sub(totalNeeded)
		for _, changeCoin := range changeCoins {
			outputs = append(outputs, &utxotypes.Output{
				ToAddr: escrowAddress.String(),
				Amount: changeCoin,
			})
		}
	}
	if _, err := k.utxoKeeper.Transfer(ctx, hash, msgOffset, escrowAddress.String(), inputs, outputs); err != nil {
		return err
	}
	return nil
}
