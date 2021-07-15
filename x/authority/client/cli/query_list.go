package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/liubaninc/m0/x/authority/types"
	"github.com/spf13/cobra"
)

func CmdListAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-account",
		Short: "list all account",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAccountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllAccount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "list-account")
	return cmd
}

func CmdListPendingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-proposed-account",
		Short: "list all proposed but not approved accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPendingAccountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllPendingAccount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "list-proposed-account")
	return cmd
}

func CmdListPendingAccountRevocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-proposed-account-to-revoke",
		Short: "list all proposed but not approved accounts to be revoked",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPendingAccountRevocationRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllPendingAccountRevocation(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "list-proposed-account-to-revoke")
	return cmd
}
