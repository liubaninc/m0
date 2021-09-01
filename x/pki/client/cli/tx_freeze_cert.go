package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/pki/types"
)

var _ = strconv.Itoa(0)

func CmdFreezeCert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freeze-cert [subject] [subjectKeyID]",
		Short: "freeze an intermediate or leaf certificate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSubject := string(args[0])
			argsSubjectKeyID := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgFreezeCert(clientCtx.GetFromAddress().String(), string(argsSubject), string(argsSubjectKeyID))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
