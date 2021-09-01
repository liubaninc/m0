package cli

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/pki/types"
)

var _ = strconv.Itoa(0)

func CmdAddCert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-cert [certificate]",
		Short: "add an intermediate or leaf certificate signed by a chain of certificates which must be already present",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsCertificate := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bts, err := ioutil.ReadFile(argsCertificate)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddCert(clientCtx.GetFromAddress().String(), string(bts))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
