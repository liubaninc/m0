package keeper

import (
	"context"
	"fmt"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	"github.com/liubaninc/m0/x/mibc/types"
)

func (k msgServer) SendIbcUTXO(goCtx context.Context, msg *types.MsgSendIbcUTXO) (*types.MsgSendIbcUTXOResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet
	msgOffset := int32(ctx.Context().Value("msg-index").(int))
	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	t := time.Now()
	defer func() {
		k.Logger(ctx).Debug("handler", "route", msg.Route(), "msg", msg.Type(), "hash", hash, "index", msgOffset, "elapsed", time.Now().Sub(t).String())
	}()

	feeAddr := authtypes.NewModuleAddress(authtypes.FeeCollectorName).String()

	var outputs []*utxotypes.Output
	var iOutputs []*types.Output
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
			if strings.Compare(msg.Sender, output.ToAddr) == 0 || strings.Compare(feeAddr, output.ToAddr) == 0 {
				// send
				outputs = append(outputs, output)
			} else {
				// ibc send
				outputs = append(outputs, &utxotypes.Output{
					ToAddr: escrowAddress.String(),
					Amount: output.Amount,
				})
				iOutputs = append(iOutputs, &types.Output{
					Addr: output.ToAddr,
					Denom: fullDenomPath,
					Amount: output.Amount.Amount,
					FrozenHeight: output.FrozenHeight,
				})
			}
		} else {
			if strings.Compare(msg.Sender, output.ToAddr) == 0 || strings.Compare(feeAddr, output.ToAddr) == 0 {
				// send
				outputs = append(outputs, output)
			} else {
				// ibc burn
				iOutputs = append(iOutputs, &types.Output{
					Addr: output.ToAddr,
					Denom: fullDenomPath,
					Amount: output.Amount.Amount,
					FrozenHeight: output.FrozenHeight,
				})
			}
		}
	}
	if err := k.utxoKeeper.Transfer(ctx, hash, msgOffset, msg.Sender, msg.Inputs, outputs); err != nil{
		return nil, err
	}

	// Construct the packet
	var packet types.IbcUTXOPacketData

	packet.Creator = msg.Sender
	packet.Outputs = iOutputs

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

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgSendIbcUTXOResponse{}, nil
}
