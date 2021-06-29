package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/blob/types"
)

var _ = strconv.Itoa(0)

func CmdSendBlob() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-blob [title] [content]",
		Short: "Post Blog",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsTitle := string(args[0])
			argsContent := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendBlob(clientCtx.GetFromAddress().String(), string(argsTitle), string(argsContent))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
