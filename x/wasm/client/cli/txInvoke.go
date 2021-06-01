package cli

import (
	"encoding/json"
	"fmt"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/wasm/types"
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
			var initArgs map[string][]byte
			if err := json.Unmarshal([]byte(args[2]), &initArgs); err != nil {
				return fmt.Errorf("init args, error %v", err)
			}
			_ = method

			var inputs []*utxotypes.Input
			var outputs []*utxotypes.Output
			var inputsExt []*types.InputExt
			var outputsExt []*types.OutputExt
			var request []*types.InvokeRequest

			msg := types.NewMsgInvoke(clientCtx.GetFromAddress().String(), inputs, outputs, inputsExt, outputsExt, request, viper.GetString(flagDesc))
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
