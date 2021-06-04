package cli

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/viper"

	"github.com/liubaninc/m0/x/utxo/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

func CmdListInput() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-input [address] [denom]",
		Short: "query all available unspent outputs list of a specific address and denom.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
				return fmt.Errorf("invalid address %s (%s)", args[0], err)
			}
			if  err := sdk.ValidateDenom(args[1]); err != nil {
				return fmt.Errorf("invalid denom %s (%s)", args[1], err)
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllInputRequest{
				Pagination: pageReq,
				Address:    args[0],
				Denom:      args[1],
			}

			res, err := queryClient.InputAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "all-inputs")

	return cmd
}

func CmdShowInput() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-input [address] [amount][,[amount]]",
		Short: "query for the available unspent outputs of a specific address and specific amount of coins",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			if _, err := sdk.AccAddressFromBech32(args[0]); err != nil {
				return fmt.Errorf("invalid address %s (%s)", args[0], err)
			}
			if _, err := sdk.ParseCoinsNormalized(args[1]); err != nil {
				return fmt.Errorf("invalid amount %s (%s)", args[1], err)
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryInputRequest{
				Address: args[0],
				Amount:  args[1],
				Lock:    viper.GetInt64(flagLock),
			}

			res, err := queryClient.Input(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	cmd.Flags().Int64(flagLock, 0, "will lock inputs for a while. eg. 60s")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
