package cli

import (
	"context"
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

func CmdSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [to_address] [amount] [[to_address] [amount]]",
		Short: "send an utxo token (denom)",
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
					neededTotal = neededTotal.Add(amount)
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
					})
				}
			}

			msg := types.NewMsgSend(clientCtx.GetFromAddress().String(), inputs, outputs, viper.GetString(flagDesc))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagDesc, "", "description of msg")
	cmd.Flags().Uint64(flagLock, 60, "will lock inputs for a while. eg. 60s")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
