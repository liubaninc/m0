package cli

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/wasm/types"
)

var _ = strconv.Itoa(0)

func CmdProposeDeployContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-deploy-contract [contractName] [contractCode] [initArgs] [approval]",
		Short: "propose a deploy of contract",
		Args:  cobra.MinimumNArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsContractName := string(args[0])

			argsContractCode, err := ioutil.ReadFile(args[1])
			if err != nil {
				return fmt.Errorf("contract code file %v, error %v", args[2], err)
			}

			initArgs, err := convertToArgs(args[2])
			if err != nil {
				return fmt.Errorf("invoke init args, error %v", err)
			}
			argsInitArgs, _ := json.Marshal(initArgs)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProposeDeployContract(clientCtx.GetFromAddress().String(), string(argsContractName), []byte(argsContractCode), string(argsInitArgs), args[3:])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
