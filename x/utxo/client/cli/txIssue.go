package cli

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/liubaninc/m0/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

const (
	flagDesc   = "desc"
	flagLock   = "lock"
	flagAesKey = "aes-key"
)

func CmdIssue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue [to_address] [amount] [[to_address] [amount]]",
		Short: "issue an utxo token (denom)",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if len(args)%2 != 0 {
				return fmt.Errorf("args should be even")
			}

			neededTotal := sdk.NewCoins()
			var inputs []*types.Input
			var outputs []*types.Output

			for index, arg := range args {
				if index%2 == 0 {
					if _, err := sdk.AccAddressFromBech32(arg); err != nil {
						return fmt.Errorf("invalid address %s (%s)", arg, err)
					}
					amount, err := sdk.ParseCoinNormalized(args[index+1])
					if err != nil {
						return fmt.Errorf("invalid amount %s (%s)", args[index+1], err)
					}
					outputs = append(outputs, &types.Output{
						ToAddr: arg,
						Amount: amount,
					})
				}
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			fees, err := fees(txf)
			if err != nil {
				return err
			}
			for _, fee := range fees {
				outputs = append(outputs, &types.Output{
					Amount: fee,
					ToAddr: authtypes.NewModuleAddress(authtypes.FeeCollectorName).String(),
				})
				neededTotal = neededTotal.Add(fee)
			}
			if !neededTotal.IsZero() {
				queryClient := types.NewQueryClient(clientCtx)
				params := &types.QueryInputRequest{
					Address: clientCtx.GetFromAddress().String(),
					Amounts: neededTotal.String(),
					Lock:    viper.GetInt64(flagLock),
				}
				res, err := queryClient.Input(context.Background(), params)
				if err != nil {
					return err
				}
				inputs = append(inputs, res.Inputs...)
				changeCoins := res.Amount.Sub(neededTotal)
				for _, changeCoin := range changeCoins {
					outputs = append(outputs, &types.Output{
						ToAddr: clientCtx.GetFromAddress().String(),
						Amount: changeCoin,
						Change: true,
					})
				}
			}

			flagAesKey := viper.GetString(flagAesKey)
			flagDesc := viper.GetString(flagDesc)

			if len(flagAesKey) > 0 {
				flagDesc = Encode(flagAesKey, flagDesc)
			}

			msg := types.NewMsgIssue(clientCtx.GetFromAddress().String(), inputs, outputs, flagDesc)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagDesc, "", "description of msg")
	cmd.Flags().String(flagAesKey, "", "aeskey to encode description of msg")
	cmd.Flags().Int64(flagLock, 60, "will lock inputs for a while. eg. 60s")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func fees(txf tx.Factory) (sdk.Coins, error) {
	fees := txf.Fees()
	gasPrices := txf.GasPrices()
	if !gasPrices.IsZero() {
		if !fees.IsZero() {
			return nil, errors.New("cannot provide both fees and gas prices")
		}

		glDec := sdk.NewDec(int64(txf.Gas()))

		// Derive the fees based on the provided gas prices, where
		// fee = ceil(gasPrice * gasLimit).
		fees = make(sdk.Coins, len(gasPrices))

		for i, gp := range gasPrices {
			fee := gp.Amount.Mul(glDec)
			fees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
		}
	}
	return fees, nil
}
