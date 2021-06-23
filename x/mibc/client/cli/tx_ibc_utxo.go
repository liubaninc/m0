package cli

import (
	"context"
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/client/utils"
	"github.com/liubaninc/m0/x/mibc/types"
)

var _ = strconv.Itoa(0)

const (
	flagDesc = "desc"
	flagLock = "lock"
)

func CmdSendIbcUTXO() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-ibc-utxo [src-port] [src-channel] [receiver] [amount]",
		Short: "Send a ibcUTXO over IBC",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsReceiver := string(args[2])
			argsAmount := string(args[3])
			lock, _ := cmd.Flags().GetInt64(flagLock)
			desc, _ := cmd.Flags().GetString(flagDesc)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			if _, err := sdk.AccAddressFromBech32(argsReceiver); err != nil {
				return fmt.Errorf("invalid address %s (%s)", argsReceiver, err)
			}
			amount, err := sdk.ParseCoinNormalized(argsAmount)
			if err != nil {
				return fmt.Errorf("invalid amount %s (%s)", argsAmount, err)
			}

			neededTotal := sdk.NewCoins()
			var inputs []*utxotypes.Input
			var outputs []*utxotypes.Output

			outputs = append(outputs, &utxotypes.Output{
				ToAddr: argsReceiver,
				Amount: amount,
			})
			neededTotal = neededTotal.Add(amount)

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags())
			fees, err := fees(txf)
			if err != nil {
				return err
			}
			for _, fee := range fees {
				outputs = append(outputs, &utxotypes.Output{
					Amount: fee,
					ToAddr: authtypes.NewModuleAddress(authtypes.FeeCollectorName).String(),
				})
				neededTotal = neededTotal.Add(fee)
			}

			if !neededTotal.IsZero() {
				queryClient := utxotypes.NewQueryClient(clientCtx)
				res, err := queryClient.Input(context.Background(), &utxotypes.QueryInputRequest{
					Address: clientCtx.GetFromAddress().String(),
					Amounts: neededTotal.String(),
					Lock:    lock,
				})
				if err != nil {
					return err
				}
				inputs = append(inputs, res.Inputs...)
				changeCoins := res.Amount.Sub(neededTotal)
				for _, changeCoin := range changeCoins {
					outputs = append(outputs, &utxotypes.Output{
						ToAddr: clientCtx.GetFromAddress().String(),
						Amount: changeCoin,
					})
				}
			}

			msg := types.NewMsgSendIbcUTXO(sender, srcPort, srcChannel, timeoutTimestamp, inputs, outputs, desc)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagDesc, "", "description of msg")
	cmd.Flags().Bool(flagLock, true, "will lock inputs for a while. eg. 60s")
	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func fees(txf tx.Factory) (sdk.Coins, error) {
	fees := txf.Fees()
	gasPrices := txf.GasPrices()
	if !gasPrices.IsZero() {
		if !fees.IsZero() {
			return nil, errors.New("cannot provide both fees and gas prices")
		}

		glDec := sdk.NewDec(int64(txf.Gas()))

		// Derive the fees based on the provided gas prices, where
		// fee = ceil(gasPrice * gasLimit).
		fees = make(sdk.Coins, len(gasPrices))

		for i, gp := range gasPrices {
			fee := gp.Amount.Mul(glDec)
			fees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
		}
	}
	return fees, nil
}
