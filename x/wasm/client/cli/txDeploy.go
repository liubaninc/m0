package cli

import (
	"context"
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/wasm/types"
)

var _ = strconv.Itoa(0)

const (
	flagDesc = "desc"
	flagLock = "lock"
	flagModule  = "moudle"
	flagRuntime = "runtime"
	flagAmount  = "amount"
)

func CmdDeploy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy [name] [code-file] [init-args]",
		Short: "deploy an wasm contract",
		Args:  cobra.RangeArgs(2, 3),
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
				return fmt.Errorf("contract code file %v, error %v", args[2], err)
			}
			var initArgs map[string][]byte
			if err := json.Unmarshal([]byte(args[2]), &initArgs); err != nil {
				return fmt.Errorf("contract init args, error %v", err)
			}
			desc := &xmodel.WasmCodeDesc{
				Runtime:      viper.GetString(flagRuntime),
				ContractType: viper.GetString(flagModule),
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.PreExec(context.Background(), &types.InvokeRPCRequest{
				Creator: clientCtx.GetFromAddress().String(),
				Lock: viper.GetInt64(flagLock),
				Requests: []*types.InvokeRequest{
					types.NewMsgDeploy(clientCtx.GetFromAddress().String(), name, code, desc, initArgs, nil, nil,nil,nil, nil, viper.GetString(flagDesc)).ConvertInvokeRequest(),
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
					Amount:  neededTotal.String(),
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

			msg := types.NewMsgDeploy(clientCtx.GetFromAddress().String(), name, code, desc, initArgs, resp.Requests[0].ResourceLimits, append(inputs, resp.Inputs...), append(outputs, resp.Outputs...), resp.InputsExt, resp.OutputsExt, viper.GetString(flagDesc))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagDesc, "", "description of msg")
	cmd.Flags().Int64(flagLock, 60, "will lock inputs for a while. eg. 60s")
	cmd.Flags().String(flagModule, "wasm", "contract code module, wasm")
	cmd.Flags().String(flagRuntime, "c", "if contract code use golang, then go or if use c lang, then c")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
