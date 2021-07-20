package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/x/validator/types"
)

var _ = strconv.Itoa(0)

const (
	FlagMoniker  = "moniker"
	FlagIdentity = "identity"
	FlagWebsite  = "website"
	FlagDetails  = "details"
)

func CmdCreateValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-validator [validator-pubkey]",
		Short: "Adds a new validator node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			valPubKeyString := args[0]
			valPubKey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, valPubKeyString)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateValidator(clientCtx.GetFromAddress().String(), sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, valPubKey), &types.Description{
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

	cmd.Flags().String(FlagMoniker, "", "The validator's name")
	cmd.Flags().String(FlagWebsite, "", "The validator's (optional) website")
	cmd.Flags().String(FlagDetails, "", "The validator's (optional) details")
	cmd.Flags().String(FlagIdentity, "", "The (optional) identity signature")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
