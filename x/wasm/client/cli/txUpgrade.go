package cli

import (
	"context"
	"fmt"
	"io/ioutil"
	"strconv"

	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/liubaninc/m0/x/wasm/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade [name] [code-file]",
		Short: "upgrade an wasm contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			if err := kernel.ValidContractName(name); err != nil {
				return fmt.Errorf("contract name %v, error %v", args[1], err)
			}
			code, err := ioutil.ReadFile(args[1])
			if err != nil {
				return fmt.Errorf("read code file %v, error %v", args[2], err)
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.PreExec(context.Background(), &types.InvokeRPCRequest{
				Creator: clientCtx.GetFromAddress().String(),
				Lock:    viper.GetInt64(flagLock),
				Requests: []*types.InvokeRequest{
					types.NewMsgUpgrade(clientCtx.GetFromAddress().String(), name, code, nil, nil, nil, nil, nil, viper.GetString(flagDesc)).ConvertInvokeRequest(),
				},
			})
			if err != nil {
				return err
			}

			var inputs []*utxotypes.Input
			var outputs []*utxotypes.Output
			neededTotal := sdk.NewCoins()
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			fees := txf.Fees()
			for _, fee := range fees {
				outputs = append(outputs, &utxotypes.Output{
					Amount: fee,
					ToAddr: authtypes.NewModuleAddress(authtypes.FeeCollectorName).String(),
				})
				neededTotal = neededTotal.Add(fee)
			}
			if !neededTotal.IsZero() {
				queryClient := utxotypes.NewQueryClient(clientCtx)
				params := &utxotypes.QueryInputRequest{
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
					outputs = append(outputs, &utxotypes.Output{
						ToAddr: clientCtx.GetFromAddress().String(),
						Amount: changeCoin,
						Change: true,
					})
				}
			}

			msg := types.NewMsgUpgrade(clientCtx.GetFromAddress().String(), name, code, resp.Requests[0].ResourceLimits, append(inputs, resp.Inputs...), append(outputs, resp.Outputs...), resp.InputsExt, resp.OutputsExt, viper.GetString(flagDesc))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagDesc, "", "description of msg")
	cmd.Flags().Int64(flagLock, 60, "will lock inputs for a while. eg. 60s")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
