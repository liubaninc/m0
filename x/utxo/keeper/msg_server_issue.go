package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/telemetry"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/liubaninc/m0/x/utxo/types"
	"github.com/tendermint/tendermint/crypto/tmhash"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Issue(goCtx context.Context, msg *types.MsgIssue) (*types.MsgIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	msgIndex := int32(ctx.Context().Value("msg-index").(int))

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
		if err := k.bankKeeper.SubtractCoins(ctx, addr, sdk.NewCoins(input.Amount)); err != nil {
			return nil, err
		}
	}

	hash := fmt.Sprintf("%X", tmhash.Sum(ctx.TxBytes()))
	for index, output := range msg.Outputs {
		attrs = append(attrs, sdk.NewAttribute(types.AttributeKeyRecipient, output.ToAddr))
		input := types.Input{
			RefTx:        hash,
			RefMsg:       msgIndex,
			RefOffset:    int32(index),
			FromAddr:     output.ToAddr,
			Amount:       output.Amount,
			FrozenHeight: output.FrozenHeight,
		}
		k.SetInput(ctx, input)
		totalOut = totalOut.Add(output.Amount)

		// bank
		addr, _ := sdk.AccAddressFromBech32(input.FromAddr)
		if addr.Equals(authtypes.NewModuleAddress(authtypes.FeeCollectorName)) {
			// TODO chaogaofeng
		} else if err := k.bankKeeper.AddCoins(ctx, addr, sdk.NewCoins(output.Amount)); err != nil {
			return nil, err
		}
		acc := k.accountKeeper.GetAccount(ctx, addr)
		if acc == nil {
			defer telemetry.IncrCounter(1, "new", "account")
			k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, addr))
		}
	}
	supply := k.bankKeeper.GetSupply(ctx)
	supply.Inflate(totalOut)
	k.bankKeeper.SetSupply(ctx, supply)

	changeCoins := totalOut.Sub(totalIn)
	for _, coin := range changeCoins {
		token, found := k.GetToken(ctx, coin.Denom)
		if found {
			if strings.Compare(token.Issuer, msg.Creator) != 0 {
				return nil, sdkerrors.Wrapf(types.ErrInvalidIssuer, "except %s get %s", token.Issuer, msg.Creator)
			}
			supply, _ := sdk.NewIntFromString(token.Supply)
			token.Supply = supply.Add(coin.Amount).String()
			circulating, _ := sdk.NewIntFromString(token.Circulating)
			token.Circulating = circulating.Add(coin.Amount).String()
		} else {
			token = types.Token{
				Name:        coin.Denom,
				Issuer:      msg.Creator,
				IssueTx:     hash,
				Supply:      coin.Amount.String(),
				Circulating: coin.Amount.String(),
			}
			defer telemetry.IncrCounter(1, "new", "token")
		}
		k.SetToken(ctx, token)
	}
	attrs = append(attrs, sdk.NewAttribute(sdk.AttributeKeyAmount, totalOut.String()))
	attrs = append(attrs, sdk.NewAttribute(types.AttributeKeyAmountChanged, changeCoins.String()))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
		),
		sdk.NewEvent(types.EventTypeTransfer, attrs...),
	})
	return &types.MsgIssueResponse{}, nil
}
