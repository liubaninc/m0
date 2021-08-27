package cli

import (
	"encoding/hex"
	"github.com/liubaninc/m0/crypto/recrypt"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/storage/types"
)

var _ = strconv.Itoa(0)

func CmdCreateRecryptAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-recrypt-account",
		Short: "Broadcast message createRecryptAccount",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			priv, err := clientCtx.Keyring.ExportPrivateKeyObject(clientCtx.GetFromName())
			if err != nil {
				return err
			}
			pubkey := recrypt.GetPubKey(priv)

			msg := types.NewMsgCreateRecryptAccount(clientCtx.GetFromAddress().String(), hex.EncodeToString(pubkey))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
