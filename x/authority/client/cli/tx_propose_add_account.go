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
		Short: "Broadcast message ProposeAddAccount",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAddress := string(args[0])
			argsPublickey := string(args[1])
			argsRoles := string(args[2])

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

			msg := types.NewMsgProposeAddAccount(clientCtx.GetFromAddress().String(), string(argsAddress), string(argsPublickey), roles)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
