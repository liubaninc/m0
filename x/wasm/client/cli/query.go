package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/liubaninc/m0/x/wasm/xmodel/contract"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"
	"github.com/spf13/viper"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/liubaninc/m0/x/wasm/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group wasm queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListContract())
	cmd.AddCommand(CmdShowContract())
	cmd.AddCommand(CmdQueryContract())
	return cmd
}

func CmdQueryContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query [contract] [method] [args]",
		Short: "query contract method",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			name := args[0]
			if err := kernel.ValidContractName(name); err != nil {
				return err
			}
			method := args[1]
			if len(method) == 0 {
				return fmt.Errorf("contract method empty")
			}
			methodArgs, err := convertToArgs(args[2])
			if err != nil {
				return fmt.Errorf("invoke init args, error %v", err)
			}
			mArgsStr, _ := json.Marshal(methodArgs)

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.PreExec(context.Background(), &types.InvokeRPCRequest{
				Creator: clientCtx.GetFromAddress().String(),
				Lock:    viper.GetInt64(flagLock),
				Requests: []*types.InvokeRequest{
					{
						ModuleName:   viper.GetString(flagModule),
						ContractName: name,
						MethodName:   method,
						Args:         string(mArgsStr),
					},
				},
			})
			if err != nil {
				return err
			}

			for _, res := range resp.Responses {
				if res.Status >= contract.StatusErrorThreshold {
					return fmt.Errorf("contract error status:%d message:%s", res.Status, res.Message)
				}
				fmt.Printf("contract response: %s\n", string(res.Body))
			}
			return nil
		},
	}

	cmd.Flags().String(flagModule, "wasm", "contract code moudle, wasm")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func convertToArgs(args string) (map[string][]byte, error) {
	args1 := make(map[string]string)
	if err := json.Unmarshal([]byte(args), &args1); err != nil {
		return nil, err
	}
	args2 := make(map[string][]byte)
	for k, v := range args1 {
		args2[k] = []byte(v)
	}
	return args2, nil
}
