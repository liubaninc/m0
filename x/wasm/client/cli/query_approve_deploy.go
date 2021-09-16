package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/spf13/cobra"
)

func CmdListApproveDeploy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-approve-deploy",
		Short: "list all approveDeploy",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllApproveDeployRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ApproveDeployAll(context.Background(), params)
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

func CmdShowApproveDeploy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-approve-deploy [name]",
		Short: "shows a approveDeploy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetApproveDeployRequest{
				Index: args[0],
			}

			res, err := queryClient.ApproveDeploy(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
