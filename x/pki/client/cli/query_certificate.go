package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/liubaninc/m0/x/pki/types"
	"github.com/spf13/cobra"
)

func CmdListCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-certificate",
		Short: "list all certificate",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllCertificateRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CertificateAll(context.Background(), params)
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

func CmdShowCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-certificate [issuer] [serialNumber]",
		Short: "shows a certificate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetCertificateRequest{
				Issuer:       args[0],
				SerialNumber: args[1],
			}

			res, err := queryClient.Certificate(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
