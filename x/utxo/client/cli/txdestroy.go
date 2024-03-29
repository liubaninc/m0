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

func CmdDestroy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "destroy [amount][,[amount]]",
		Short: "destroy an utxo token (denom)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return fmt.Errorf("invalid amount %s (%s)", args[0], err)
			}

			neededTotal := sdk.NewCoins()
			var inputs []*types.Input
			var outputs []*types.Output

			neededTotal = neededTotal.Add(coins...)
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

			msg := types.NewMsgDestroy(clientCtx.GetFromAddress().String(), inputs, outputs, flagDesc)
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
