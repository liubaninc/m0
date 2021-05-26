package keeper

import (
	"context"
	"fmt"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/utxo/types"
)

func (k msgServer) Destroy(goCtx context.Context, msg *types.MsgDestroy) (*types.MsgDestroyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	totalIn := sdk.NewCoins()
	totalOut := sdk.NewCoins()
	var attrs []sdk.Attribute
	for index, input := range msg.Inputs {
		attrs = append(attrs, sdk.NewAttribute(types.AttributeKeySender, input.FromAddr))
		tinput, found := k.GetInput(ctx, input.Index())
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrUTXONotFound, "index %d in inputs", index)
		}
		if !tinput.Amount.Equal(input.Amount) {
			return nil, sdkerrors.Wrapf(types.ErrUTXONotMismatch, "index %d in inputs, amount expect %s get %s", index, input.Amount, tinput.Amount)
		}
		if tinput.FrozenHeight != tinput.FrozenHeight {
			return nil, sdkerrors.Wrapf(types.ErrUTXONotMismatch, "index %d in inputs, frozen height expect %s get %s", index, input.FrozenHeight, tinput.FrozenHeight)
		}
		if input.FrozenHeight == -1 || input.FrozenHeight > ctx.BlockHeight() {
			return nil, sdkerrors.Wrapf(types.ErrUTXOFrozen, "index %d in inputs, frozen height expect %s get %s", index, input.FrozenHeight, ctx.BlockHeight())
		}

		k.RemoveInput(ctx, input.Index())
		totalIn = totalIn.Add(input.Amount)

		// bank
		addr, _ := sdk.AccAddressFromBech32(input.FromAddr)
		if err := k.bank.SubtractCoins(ctx, addr, sdk.NewCoins(input.Amount)); err != nil {
			return nil, err
		}
	}

	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	for index, output := range msg.Outputs {
		attrs = append(attrs, sdk.NewAttribute(types.AttributeKeyRecipient, output.ToAddr))
		input := types.Input{
			RefTx:        hash,
			RefMsg:       0, // TODO chaogaofeng
			RefOffset:    int32(index),
			FromAddr:     output.ToAddr,
			Amount:       output.Amount,
			FrozenHeight: output.FrozenHeight,
		}
		k.SetInput(ctx, input)

		totalOut = totalOut.Add(output.Amount)

		// bank
		addr, _ := sdk.AccAddressFromBech32(output.ToAddr)
		if !addr.Equals(authtypes.NewModuleAddress(authtypes.FeeCollectorName)) {
			// TODO chaogaofeng
		} else if err := k.bank.AddCoins(ctx, addr, sdk.NewCoins(output.Amount)); err != nil {
			return nil, err
		}
	}

	changeCoins := totalIn.Sub(totalOut)
	for _, coin := range changeCoins {
		token, found := k.GetToken(ctx, coin.Denom)
		if !found {
			panic("token not found")
		}
		circulating, _ := sdk.NewIntFromString(token.Circulating)
		token.Circulating = circulating.Sub(coin.Amount).String()
		k.SetToken(ctx, token)
	}
	supply := k.bank.GetSupply(ctx)
	supply.Deflate(changeCoins)
	k.bank.SetSupply(ctx, supply)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, msg.Type()),
		),
		sdk.NewEvent(types.EventTypeTransfer, attrs...),
	})

	return &types.MsgDestroyResponse{}, nil
}
