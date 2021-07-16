package cli

import (
	"github.com/spf13/cobra"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/authority/types"
)

var _ = strconv.Itoa(0)

func CmdProposeModifyAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-modify-account [address] [roles]",
		Short: "Propose modify of the account with the given address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAddress := args[0]
			argsRoles := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			var roles []string
			if len(argsRoles) > 0 {
				for _, role := range strings.Split(argsRoles, ",") {
					roles = append(roles, role)
				}
			}

			msg := types.NewMsgProposeModifyAccountRequest(clientCtx.GetFromAddress().String(), argsAddress, roles)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String("desc", "", "description of msg")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
