package cli

import (
	"encoding/hex"
	"fmt"
	"github.com/liubaninc/m0/crypto/recrypt"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/storage/types"
)

var _ = strconv.Itoa(0)

func CmdCreateEncryptStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-encrypt-storage name content",
		Short: "Broadcast message createEncryptStorage",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			content := args[1]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			priv, err := clientCtx.Keyring.ExportPrivateKeyObject(clientCtx.GetFromName())
			if err != nil {
				return err
			}
			pubkey := recrypt.GetPubKey(priv)

			md5 := hex.EncodeToString(tmhash.Sum([]byte(content)))
			encrypted, envelope, err := recrypt.Encrypt(priv, pubkey, content)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateEncryptStorage(clientCtx.GetFromAddress().String(), key, md5, encrypted, envelope)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func MD5(fileName string) (string, int) {
	bts, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", tmhash.Sum(bts)), len(bts)
}
