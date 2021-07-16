package cli

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/liubaninc/m0/x/authority/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group authority queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(GetCmdAccount())
	cmd.AddCommand(GetCmdPendingAccount())
	cmd.AddCommand(GetCmdPendingAccountRevocation())

	cmd.AddCommand(CmdListAccount())
	cmd.AddCommand(CmdListPendingAccount())
	cmd.AddCommand(CmdListPendingAccountRevocation())
	return cmd
}

func GetCmdAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-account [address]",
		Short: "show account associated with the address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			if err := sdk.ValidateDenom(args[0]); err != nil {
				return fmt.Errorf("invalid address %s (%s)", args[0], err)
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAccountRequest{
				Address: args[0],
			}

			res, err := queryClient.Account(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
func GetCmdPendingAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-proposed-account [address]",
		Short: "show proposed account by the address but not approved",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			if err := sdk.ValidateDenom(args[0]); err != nil {
				return fmt.Errorf("invalid address %s (%s)", args[0], err)
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryPendingAccountRequest{
				Address: args[0],
			}

			res, err := queryClient.PendingAccount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdPendingAccountRevocation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-proposed-account-to-revoke [address]",
		Short: "show proposed to be revoked account by the address but not approved",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			if err := sdk.ValidateDenom(args[0]); err != nil {
				return fmt.Errorf("invalid address %s (%s)", args[0], err)
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryPendingAccountRevocationRequest{
				Address: args[0],
			}

			res, err := queryClient.PendingAccountRevocation(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
