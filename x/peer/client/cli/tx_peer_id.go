package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/peer/types"
)

func CmdCreatePeerID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-peer-id [nodeID] [certIssuer] [certSerialNum]",
		Short: "Create a new peerID",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]
			argsCertIssuer, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsCertSerialNum, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePeerID(clientCtx.GetFromAddress().String(), index, argsCertIssuer, argsCertSerialNum)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdatePeerID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-peer-id [nodeID] [certIssuer] [certSerialNum]",
		Short: "Update a peerID",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			argsCertIssuer, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsCertSerialNum, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePeerID(clientCtx.GetFromAddress().String(), index, argsCertIssuer, argsCertSerialNum)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeletePeerID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-peer-id [nodeID]",
		Short: "Delete a peerID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePeerID(clientCtx.GetFromAddress().String(), index)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
