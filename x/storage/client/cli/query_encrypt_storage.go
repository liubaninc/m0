package cli

import (
	"context"
	"github.com/liubaninc/m0/crypto/recrypt"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/liubaninc/m0/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdListEncryptStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-encrypt-storage",
		Short: "list all encryptStorage",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllEncryptStorageRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.EncryptStorageAll(context.Background(), params)
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

func CmdShowEncryptStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-encrypt-storage key",
		Short: "shows a encryptStorage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetEncryptStorageRequest{
				Index: args[0],
			}

			res, err := queryClient.EncryptStorage(context.Background(), params)
			if err != nil {
				return err
			}

			if len(viper.GetString(flags.FlagFrom)) != 0 {
				priv, err := clientCtx.Keyring.ExportPrivateKeyObject(clientCtx.GetFromName())
				if err != nil {
					return err
				}

				pubkey := recrypt.GetPubKey(priv)
				content, err := recrypt.Dencrypt(priv, pubkey, res.EncryptStorage.Envelope, res.EncryptStorage.Encrypted)
				if err != nil {
					return err
				}
				res.EncryptStorage.Encrypted = content
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
