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

func CmdRevokeRootCert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-root-cert [subject] [subjectKeyID]",
		Short: "revoke the given root certificate.All the certificates in the subtree signed by the revoked certificate will be revoked as well.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSubject := string(args[0])
			argsSubjectKeyID := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeRootCert(clientCtx.GetFromAddress().String(), string(argsSubject), string(argsSubjectKeyID))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
