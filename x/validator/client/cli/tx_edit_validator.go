package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/validator/types"
)

var _ = strconv.Itoa(0)

func CmdEditValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-validator",
		Short: "edit an existing validator node",
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditValidator(clientCtx.GetFromAddress().String(), &types.Description{
				Moniker:  viper.GetString(FlagMoniker),
				Identity: viper.GetString(FlagIdentity),
				Website:  viper.GetString(FlagWebsite),
				Details:  viper.GetString(FlagDetails),
			})
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagMoniker, types.DoNotModifyDesc, "The validator's name")
	cmd.Flags().String(FlagWebsite, types.DoNotModifyDesc, "The validator's (optional) website")
	cmd.Flags().String(FlagDetails, types.DoNotModifyDesc, "The validator's (optional) details")
	cmd.Flags().String(FlagIdentity, types.DoNotModifyDesc, "The (optional) identity signature")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
