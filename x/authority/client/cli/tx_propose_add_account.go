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

func CmdProposeAddAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-add-account [address] [publickey] [roles]",
		Short: "Propose a new account with the given address, public key and roles",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAddress := args[0]
			argsPublickey := args[1]
			argsRoles := args[2]

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

			msg := types.NewMsgProposeAddAccount(clientCtx.GetFromAddress().String(), argsAddress, argsPublickey, roles)
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
