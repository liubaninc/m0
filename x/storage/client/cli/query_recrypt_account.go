package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/liubaninc/m0/x/storage/types"
	"github.com/spf13/cobra"
)

func CmdListRecryptAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-recrypt-account",
		Short: "list all recryptAccount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRecryptAccountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RecryptAccountAll(context.Background(), params)
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

func CmdShowRecryptAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-recrypt-account [index]",
		Short: "shows a recryptAccount",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetRecryptAccountRequest{
				Index: args[0],
			}

			res, err := queryClient.RecryptAccount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
