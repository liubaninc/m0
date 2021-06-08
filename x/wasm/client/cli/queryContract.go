package cli

import (
	"context"
	"fmt"

	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CmdListContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-contract [[address]]",
		Short: "list all contract",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			if len(args) == 1 {
				if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
					return fmt.Errorf("invalid address %s (%s)", args[0], err)
				}

				params := &types.QueryGetAccountAllContractRequest{
					Name:       args[0],
					Pagination: pageReq,
				}

				res, err := queryClient.AccountContractAll(context.Background(), params)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			}

			params := &types.QueryAllContractRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ContractAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "list-contracts")

	return cmd
}

func CmdShowContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-contract [name]",
		Short: "shows a contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			if err := kernel.ValidContractName(args[0]); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetContractRequest{
				Name: args[0],
			}

			res, err := queryClient.Contract(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
