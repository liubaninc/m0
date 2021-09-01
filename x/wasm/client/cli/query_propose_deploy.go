package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/spf13/cobra"
)

func CmdListProposeDeploy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-propose-deploy",
		Short: "list all proposeDeploy",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllProposeDeployRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProposeDeployAll(context.Background(), params)
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

func CmdShowProposeDeploy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-propose-deploy [index]",
		Short: "shows a proposeDeploy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetProposeDeployRequest{
				Index: args[0],
			}

			res, err := queryClient.ProposeDeploy(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
