package cli

import (
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/wasm/types"
)

var _ = strconv.Itoa(0)

func CmdUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy [desc]",
		Short: "deploy an wasm contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var inputs []*utxotypes.Input
			var outputs []*utxotypes.Output
			var inputsExt []*types.InputExt
			var outputsExt []*types.OutputExt
			var request []*types.InvokeRequest

			msg := types.NewMsgUpgrade(clientCtx.GetFromAddress().String(), inputs, outputs, inputsExt, outputsExt, request, viper.GetString(flagDesc))
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
