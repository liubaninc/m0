package cli

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/liubaninc/m0/crypto/recrypt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/storage/types"
)

var _ = strconv.Itoa(0)

func CmdCreateShareEncryptStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-share-encrypt-storage name share_name address",
		Short: "Broadcast message createShareEncryptStorage",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			shareKey := args[1]
			address := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			var newPubkey []byte
			if res, err := queryClient.RecryptAccount(context.Background(), &types.QueryGetRecryptAccountRequest{
				Index: address,
			}); err == nil {
				newPubkey, _ = hex.DecodeString(res.RecryptAccount.Pubkey)
			}

			priv, err := clientCtx.Keyring.ExportPrivateKeyObject(clientCtx.GetFromName())
			if err != nil {
				return err
			}

			envelope := ""
			pubkey := recrypt.GetPubKey(priv)
			if viper.GetBool("reshare") {
				params := &types.QueryGetShareEncryptStorageRequest{
					Index: shareKey,
				}
				res, err := queryClient.ShareEncryptStorage(context.Background(), params)
				if err != nil {
					return err
				}
				if clientCtx.GetFromAddress().String() != res.ShareEncryptStorage.Sharer {
					return fmt.Errorf("no permssion")
				}
				envelope = res.ShareEncryptStorage.Envelope
				shareKey = res.ShareEncryptStorage.ShareIndex

				if res, err := queryClient.RecryptAccount(context.Background(), &types.QueryGetRecryptAccountRequest{
					Index: res.ShareEncryptStorage.Creator,
				}); err == nil {
					pubkey, _ = hex.DecodeString(res.RecryptAccount.Pubkey)
				}
			} else {
				params := &types.QueryGetEncryptStorageRequest{
					Index: shareKey,
				}
				res, err := queryClient.EncryptStorage(context.Background(), params)
				if err != nil {
					return err
				}
				if clientCtx.GetFromAddress().String() != res.EncryptStorage.Creator {
					return fmt.Errorf("no permssion")
				}
				envelope = res.EncryptStorage.Envelope
			}

			reEnvelope, err := recrypt.ReEnvelope(priv, pubkey, newPubkey, envelope)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateShareEncryptStorage(clientCtx.GetFromAddress().String(), key, shareKey, reEnvelope, address)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	cmd.Flags().Bool("reshare", false, "none original")
	return cmd
}
