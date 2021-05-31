package cli

import (
	"fmt"
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
)

func CmdDeploy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy [name] [code] [method] [args]",
		Short: "deploy an wasm contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			code, err := ioutil.ReadFile(args[1])
			if err != nil {
				return fmt.Errorf("read code file %v, error %v", args[2], err)
			}
			method := ""
			var method_args map[string][]byte

			var inputsExt []*types.InputExt
			var outputsExt []*types.OutputExt

			msg := types.NewMsgDeploy(clientCtx.GetFromAddress().String(), name, code, nil, method, method_args, inputsExt, outputsExt, viper.GetString(flagDesc))
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
