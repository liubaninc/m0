package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/liubaninc/m0/x/pki/types"
	"github.com/spf13/cobra"
)

func CmdListCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-certificates",
		Short: "list all certificates",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllCertificatesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.CertificatesAll(context.Background(), params)
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

func CmdShowCertificates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-certificates [subject] [subjectKeyID]",
		Short: "shows a certificates",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetCertificatesRequest{
				Subject:      args[0],
				SubjectKeyID: args[1],
			}

			res, err := queryClient.Certificates(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
