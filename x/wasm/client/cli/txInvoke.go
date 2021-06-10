package cli

import (
	"context"
	"encoding/json"
	"fmt"
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

func CmdInvoke() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke [name] [method] [args]",
		Short: "invoke an wasm contract's method",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			if err := kernel.ValidContractName(name); err != nil {
				return fmt.Errorf("contract name %v, error %v", args[1], err)
			}
			method := args[1]
			if len(method) == 0 {
				return fmt.Errorf("contract method empty")
			}
			methodArgs, err := convertToArgs(args[2])
			if err != nil {
				return fmt.Errorf("invoke init args, error %v", err)
			}
			mArgsStr, _ := json.Marshal(methodArgs)
			amount := sdk.NewCoins()
			if amountStr := viper.GetString(flagAmount); len(amountStr) != 0 {
				coins, err := sdk.ParseCoinsNormalized(amountStr)
				if err != nil {
					return fmt.Errorf("invalid amount %v", err)
				}
				amount = coins
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.PreExec(context.Background(), &types.InvokeRPCRequest{
				Creator: clientCtx.GetFromAddress().String(),
				Lock:    viper.GetInt64(flagLock),
				Requests: []*types.InvokeRequest{
					{
						Amount:       amount,
						ModuleName:   viper.GetString(flagModule),
						ContractName: name,
						MethodName:   method,
						Args:         string(mArgsStr),
					},
				},
			})
			if err != nil {
				return err
			}

			var inputs []*utxotypes.Input
			var outputs []*utxotypes.Output
			neededTotal := sdk.NewCoins()
			for _, coin := range amount {
				outputs = append(outputs, &utxotypes.Output{
					Amount: coin,
					ToAddr: authtypes.NewModuleAddress(name).String(),
				})
				neededTotal = neededTotal.Add(coin)
			}
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
					Amounts:  neededTotal.String(),
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
					})
				}
			}

			msg := types.NewMsgInvoke(clientCtx.GetFromAddress().String(), append(inputs, resp.Inputs...), append(outputs, resp.Outputs...), resp.InputsExt, resp.OutputsExt, resp.Requests, viper.GetString(flagDesc))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagDesc, "", "description of msg")
	cmd.Flags().Int64(flagLock, 60, "will lock inputs for a while. eg. 60s")
	cmd.Flags().String(flagModule, "wasm", "contract code moudle, wasm")
	cmd.Flags().String(flagAmount, "", "the amount transfer to contract")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
