package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/tendermint/tendermint/crypto/tmhash"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/liubaninc/m0/x/mibc/types"
)

func (k msgServer) SendIbcUTXO(goCtx context.Context, msg *types.MsgSendIbcUTXO) (*types.MsgSendIbcUTXOResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet
	msgOffset := int32(ctx.Context().Value(baseapp.KeyMsgOffset).(int))
	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	var outputs []*utxotypes.Output
	var iOutputs []*types.Output
	feeAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName).String()
	escrowAddress := types.GetEscrowAddress(msg.Port, msg.ChannelID)
	for _, output := range msg.Outputs {
		fullDenomPath := output.Amount.Denom
		// deconstruct the token denomination into the denomination trace info
		// to determine if the sender is the source chain
		if strings.HasPrefix(fullDenomPath, types.DenomPrefix+"/") {
			var err error
			fullDenomPath, err = k.DenomPathFromHash(ctx, fullDenomPath)
			if err != nil {
				return nil, err
			}
		}

		if types.SenderChainIsSource(msg.Port, msg.ChannelID, fullDenomPath) {
			if output.Change || strings.Compare(feeAddr, output.ToAddr) == 0 {
				// send
				outputs = append(outputs, output)
			} else {
				// ibc send
				outputs = append(outputs, &utxotypes.Output{
					ToAddr: escrowAddress.String(),
					Amount: output.Amount,
				})
				iOutputs = append(iOutputs, &types.Output{
					Addr:         output.ToAddr,
					Denom:        fullDenomPath,
					Amount:       output.Amount.Amount,
					FrozenHeight: output.FrozenHeight,
				})
			}
		} else {
			if output.Change || strings.Compare(feeAddr, output.ToAddr) == 0 {
				// send
				outputs = append(outputs, output)
			} else {
				// ibc burn
				iOutputs = append(iOutputs, &types.Output{
					Addr:         output.ToAddr,
					Denom:        fullDenomPath,
					Amount:       output.Amount.Amount,
					FrozenHeight: output.FrozenHeight,
				})
			}
		}
	}

	attrs, err := k.utxoKeeper.Transfer(ctx, hash, msgOffset, msg.Sender, msg.Inputs, outputs)
	if err != nil {
		return nil, err
	}
	_ = attrs

	// Construct the packet
	var packet types.IbcUTXOPacketData

	packet.Creator = msg.Sender
	packet.Outputs = iOutputs
	packet.Hash = hash

	// Transmit the packet
	err = k.TransmitIbcUTXOPacket(
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

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgSendIbcUTXOResponse{}, nil
}
