package cli

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/liubaninc/m0/x/utxo/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func CmdListToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-token",
		Short: "list all token",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTokenRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TokenAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "list-tokens")

	return cmd
}

func CmdShowToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-token [name]",
		Short: "shows a token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			if  err := sdk.ValidateDenom(args[0]); err != nil {
				return fmt.Errorf("invalid token %s (%s)", args[0], err)
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetTokenRequest{
				Name: args[0],
			}

			res, err := queryClient.Token(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
