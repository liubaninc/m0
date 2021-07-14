package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/authority/types"
)

var _ = strconv.Itoa(0)

func CmdApproveAddAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-add-account [address]",
		Short: "Broadcast message ApproveAddAccount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAddress := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgApproveAddAccountRequest(clientCtx.GetFromAddress().String(), string(argsAddress))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
