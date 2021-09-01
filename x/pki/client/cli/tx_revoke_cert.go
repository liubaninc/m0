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

func CmdRevokeCert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-cert [subject] [subjectKeyID]",
		Short: "revoke the given intermediate or leaf certificate.All the certificates in the subtree signed by the revoked certificate will be revoked as well.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			subject := string(args[0])
			subjectKeyID := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeCert(clientCtx.GetFromAddress().String(), subject, subjectKeyID)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
