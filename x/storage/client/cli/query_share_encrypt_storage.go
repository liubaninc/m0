package cli

import (
	"context"
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/liubaninc/m0/crypto/recrypt"
	"github.com/liubaninc/m0/x/storage/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func CmdListShareEncryptStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-share-encrypt-storage",
		Short: "list all shareEncryptStorage",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllShareEncryptStorageRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ShareEncryptStorageAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowShareEncryptStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-share-encrypt-storage key",
		Short: "shows a shareEncryptStorage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetShareEncryptStorageRequest{
				Index: args[0],
			}

			res, err := queryClient.ShareEncryptStorage(context.Background(), params)
			if err != nil {
				return err
			}

			if len(viper.GetString(flags.FlagFrom)) != 0 {
				res1, err := queryClient.EncryptStorage(context.Background(), &types.QueryGetEncryptStorageRequest{
					Index: res.ShareEncryptStorage.ShareIndex,
				})
				if err != nil {
					return err
				}

				var pubkey []byte
				if res, err := queryClient.RecryptAccount(context.Background(), &types.QueryGetRecryptAccountRequest{
					Index: res.ShareEncryptStorage.Creator,
				}); err == nil {
					pubkey, _ = hex.DecodeString(res.RecryptAccount.Pubkey)
				}

				priv, err := clientCtx.Keyring.ExportPrivateKeyObject(clientCtx.GetFromName())
				if err != nil {
					return err
				}
				content, err := recrypt.Dencrypt(priv, pubkey, res.ShareEncryptStorage.Envelope, res1.EncryptStorage.Encrypted)
				if err != nil {
					return err
				}
				res.ShareEncryptStorage.Envelope = content
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String(flags.FlagFrom, "", "specify user to decrypt the encrypted content")
	cmd.Flags().String(flags.FlagKeyringDir, "", "The client Keyring directory; if omitted, the default 'home' directory will be used")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|kwallet|pass|test|memory)")

	return cmd
}
