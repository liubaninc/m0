package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	wasmtypes "github.com/liubaninc/m0/x/wasm/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/spf13/cobra"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"io/ioutil"
	"math"
	"strconv"
	"time"
)

var (
	flagFrom  = "from"
	flagDenom = "denom"

	flagDuration   = "duration"
	flagModuleUTXO = "utxo"
	flagModuleWASM = "wasm"

	flagWASMName  = "wasm_code"
	flagWASMArg   = "wasm_init_arg"
	flagMethod    = "wasm_method"
	flagMethodArg = "wasm_method_arg"
)

func automaticCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "automatic [max_num]",
		Short: "automatic send tx",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			maxNum := uint64(math.MaxUint64)
			if len(args) > 0 {
				num, err := strconv.ParseUint(args[0], 10, 64)
				if err != nil {
					return fmt.Errorf("invalid max_num %s (%s)", args[0], err)
				}
				maxNum = num
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			_ = config

			denom, _ := cmd.Flags().GetString(flagDenom)
			from, _ := cmd.Flags().GetString(flagFrom)
			duration, _ := cmd.Flags().GetInt(flagDuration)
			var contractCode []byte
			var contractInitArgs string
			var contractMethod string
			var contractMethodArgs string
			if ok, _ := cmd.Flags().GetBool(flagModuleWASM); ok {
				codeFile, err := cmd.Flags().GetString(flagWASMName)
				if err != nil {
					return err
				}
				bts, err := ioutil.ReadFile(codeFile)
				if err != nil {
					return err
				}
				contractCode = bts

				initArg, err := cmd.Flags().GetString(flagWASMArg)
				if err != nil {
					return err
				}
				initArg, err = convertToArgs(initArg)
				if err != nil {
					return err
				}
				contractInitArgs = initArg

				contractMethod, err = cmd.Flags().GetString(flagMethod)
				if err != nil {
					return err
				}

				methodArg, err := cmd.Flags().GetString(flagMethodArg)
				if err != nil {
					return err
				}
				methodArg, err = convertToArgs(methodArg)
				if err != nil {
					return err
				}
				contractMethodArgs = methodArg
			}

			kr := keyring.NewInMemory()
			// 导入账户
			rid := "r" + tmrand.NewRand().Str(6)
			info, err := kr.NewAccount(rid, from, keyring.DefaultBIP39Passphrase, sdk.GetConfig().GetFullFundraiserPath(), hd.Secp256k1)
			if err != nil {
				return err
			}

			clientCtx = clientCtx.WithKeyring(kr).
				WithBroadcastMode("block").
				WithSkipConfirmation(true).
				WithFromAddress(info.GetAddress()).
				WithFromName(rid)


			utxoQueryClient := utxotypes.NewQueryClient(clientCtx)
			wasmQueryClient := wasmtypes.NewQueryClient(clientCtx)
			for i := 0; i < 10; i++ {
				res, err := utxoQueryClient.Input(context.Background(), &utxotypes.QueryInputRequest{
					Address: clientCtx.GetFromAddress().String(),
					Amounts: "10"+denom,
					Lock: 0,
				})
				if err == nil && len(res.Inputs) > 0 {
					break
				}
				fmt.Println("query test token", clientCtx.ChainID, clientCtx.NodeURI, err)
				time.Sleep(5 * time.Second)
			}

			// 新建测试账户
			uid := "auto" + tmrand.NewRand().Str(6)
			info, _, err = clientCtx.Keyring.NewMnemonic(uid, keyring.English, sdk.GetConfig().GetFullFundraiserPath(), hd.Secp256k1)
			if err != nil {
				return err
			}

			sendMsg := func(msg sdk.Msg) error {
				if err := msg.ValidateBasic(); err != nil {
					return err
				}
				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			}

			// 1、注册测试账户地址
			registerMsg, err := send(utxoQueryClient, clientCtx.GetFromAddress().String(), info.GetAddress().String(), "1"+denom, "auto register")
			if err != nil {
				fmt.Errorf("register error %s\n", err)
				return err
			}
			if err := sendMsg(registerMsg); err != nil {
				fmt.Errorf("register error %s\n", err)
				return err
			}
			clientCtx = clientCtx.WithFromName(uid).WithFromAddress(info.GetAddress())
			// 2、 初始测试账户代币及合约
			if ok, _ := cmd.Flags().GetBool(flagModuleUTXO); ok {
				// 2.1、初始发行资产
				issueMsg, err := issue(clientCtx.GetFromAddress().String(), clientCtx.GetFromAddress().String(), "1000000"+uid, "auto issue")
				if err != nil {
					fmt.Errorf("issue error %s\n", err)
					return err
				}
				if err := sendMsg(issueMsg); err != nil {
					fmt.Errorf("issue error %s\n", err)
					return err
				}
			}
			if ok, _ := cmd.Flags().GetBool(flagModuleWASM); ok {
				// 2.2、初始部署合约
				deployMsg, err := deploy(wasmQueryClient, clientCtx.GetFromAddress().String(), uid, contractCode, contractInitArgs, "auto deploy")
				if err != nil {
					fmt.Errorf("deploy error %s\n", err)
					return err
				}
				if err := sendMsg(deployMsg); err != nil {
					fmt.Errorf("deploy error %s\n", err)
					return err
				}
			}

			// 3、 测试账户随机发送交易
			for i := uint64(0); i < maxNum; i++ {
				if ok, _ := cmd.Flags().GetBool(flagModuleUTXO); ok {
					// 转移资产 8
					// 销毁资产 1
					// 增发资产 1
					var msg sdk.Msg
					var err error
					if n := tmrand.Intn(10); n < 8 {
						msg, err = send(utxoQueryClient, clientCtx.GetFromAddress().String(), clientCtx.GetFromAddress().String(), strconv.FormatInt(int64(n+1), 10)+uid, "auto send")
					} else if n < 9 {
						msg, err = destroy(utxoQueryClient, clientCtx.GetFromAddress().String(), strconv.FormatInt(int64(n), 10)+uid, "auto destroy")
					} else {
						msg, err = issue(clientCtx.GetFromAddress().String(), clientCtx.GetFromAddress().String(), strconv.FormatInt(int64(n*100), 10)+uid, "auto reissue")
					}
					if err != nil {
						fmt.Errorf("auto error %s\n", err)
						return err
					}
					if err := sendMsg(msg); err != nil {
						fmt.Errorf("auto error %s\n", err)
						return err
					}
				}

				if duration > 0 {
					time.Sleep(time.Duration(int64(tmrand.Intn(duration)+1)) * time.Second)
				}

				if ok, _ := cmd.Flags().GetBool(flagModuleWASM); ok {
					// 调用合约 8
					invokeMsg, err := invoke(wasmQueryClient, clientCtx.GetFromAddress().String(), uid, contractMethod, contractMethodArgs, "auto invoke")
					if err != nil {
						fmt.Errorf("invoke error %s\n", err)
						return err
					}
					if err := sendMsg(invokeMsg); err != nil {
						fmt.Errorf("invoke error %s\n", err)
					}
				}

				if duration > 0 {
					time.Sleep(time.Duration(int64(tmrand.Intn(duration)+1)) * time.Second)
				}
			}
			return nil
		},
	}
	cmd.Flags().Int(flagDuration, 0, "sleep duration, default 0s")
	cmd.Flags().Bool(flagModuleUTXO, true, "utxo module tx")
	cmd.Flags().Bool(flagModuleWASM, true, "wasm module tx")

	cmd.Flags().String(flagWASMName, "./wasm/counter.wasm", "wasm file for deploy")
	cmd.Flags().String(flagWASMArg, "{\"creator\":\"aaa\"}", "init arg for deploy")
	cmd.Flags().String(flagMethod, "increase", "method for invoke")
	cmd.Flags().String(flagMethodArg, "{\"key\":\"aaa\"}", "method arg for invoke")

	cmd.Flags().String(flagFrom, "key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear", "will to register other address")
	cmd.Flags().String(flagDenom, "m0token", "used denom for registers address")

	cmd.PersistentFlags().String(flags.FlagNode, "tcp://localhost:26657", "<host>:<port> to Tendermint RPC interface for this chain")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")
	cmd.Flags().Float64(flags.FlagGasAdjustment, 1.5, "adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored ")
	cmd.Flags().String(flags.FlagGas, flags.GasFlagAuto, fmt.Sprintf("gas limit to set per-transaction; set to %q to calculate sufficient gas automatically (default %d)", flags.GasFlagAuto, flags.DefaultGasLimit))

	return cmd
}

func issue(from, to string, amounts string, desc string) (sdk.Msg, error) {
	var inputs []*utxotypes.Input
	var outputs []*utxotypes.Output
	amountCoins, err := sdk.ParseCoinsNormalized(amounts)
	if err != nil {
		return nil, err
	}
	for _, amountCoin := range amountCoins {
		outputs = append(outputs, &utxotypes.Output{
			ToAddr: to,
			Amount: amountCoin,
		})
	}
	return utxotypes.NewMsgIssue(from, inputs, outputs, desc), nil
}

func send(queryClient utxotypes.QueryClient, from string, to string, amounts string, desc string) (sdk.Msg, error) {
	var inputs []*utxotypes.Input
	var outputs []*utxotypes.Output
	amountCoins, err := sdk.ParseCoinsNormalized(amounts)
	if err != nil {
		return nil, err
	}
	for _, amountCoin := range amountCoins {
		outputs = append(outputs, &utxotypes.Output{
			ToAddr: to,
			Amount: amountCoin,
		})
	}

	res, err := queryClient.Input(context.Background(), &utxotypes.QueryInputRequest{
		Address: from,
		Amounts: amounts,
		Lock:    60,
	})
	if err != nil {
		return nil, err
	}
	inputs = append(inputs, res.Inputs...)
	changeCoins := res.Amount.Sub(amountCoins)
	for _, changeCoin := range changeCoins {
		outputs = append(outputs, &utxotypes.Output{
			ToAddr: from,
			Amount: changeCoin,
		})
	}
	return utxotypes.NewMsgSend(from, inputs, outputs, desc), nil
}

func destroy(queryClient utxotypes.QueryClient, from string, amounts string, desc string) (sdk.Msg, error) {
	var inputs []*utxotypes.Input
	var outputs []*utxotypes.Output

	amountCoins, err := sdk.ParseCoinsNormalized(amounts)
	if err != nil {
		return nil, err
	}

	res, err := queryClient.Input(context.Background(), &utxotypes.QueryInputRequest{
		Address: from,
		Amounts: amounts,
		Lock:    60,
	})
	if err != nil {
		return nil, err
	}
	inputs = append(inputs, res.Inputs...)
	changeCoins := res.Amount.Sub(amountCoins)
	for _, changeCoin := range changeCoins {
		outputs = append(outputs, &utxotypes.Output{
			ToAddr: from,
			Amount: changeCoin,
		})
	}
	return utxotypes.NewMsgDestroy(from, inputs, outputs, desc), nil
}

func deploy(queryClient wasmtypes.QueryClient, from string, name string, code []byte, args string, desc string) (sdk.Msg, error) {
	codeDesc := &xmodel.WasmCodeDesc{
		Runtime:      "c",
		ContractType: "wasm",
	}

	tm := wasmtypes.NewMsgDeploy(from, name, code, codeDesc, args, nil, nil, nil, nil, nil, desc)
	res, err := queryClient.PreExec(context.Background(), &wasmtypes.InvokeRPCRequest{
		Creator: from,
		Lock:    60,
		Requests: []*wasmtypes.InvokeRequest{
			tm.ConvertInvokeRequest(),
		},
	})
	if err != nil {
		return nil, err
	}
	return wasmtypes.NewMsgDeploy(from, name, code, codeDesc, args, res.Requests[0].ResourceLimits, res.Inputs, res.Outputs, res.InputsExt, res.OutputsExt, desc), nil
}

func upgrade(queryClient wasmtypes.QueryClient, from string, name string, code []byte, desc string) (sdk.Msg, error) {
	tm := wasmtypes.NewMsgUpgrade(from, name, code, nil, nil, nil, nil, nil, desc)
	res, err := queryClient.PreExec(context.Background(), &wasmtypes.InvokeRPCRequest{
		Creator: from,
		Lock:    60,
		Requests: []*wasmtypes.InvokeRequest{
			tm.ConvertInvokeRequest(),
		},
	})
	if err != nil {
		return nil, err
	}
	return wasmtypes.NewMsgUpgrade(from, name, code, res.Requests[0].ResourceLimits, res.Inputs, res.Outputs, res.InputsExt, res.OutputsExt, desc), nil
}

func invoke(queryClient wasmtypes.QueryClient, from string, name string, method string, args string, desc string) (sdk.Msg, error) {
	res, err := queryClient.PreExec(context.Background(), &wasmtypes.InvokeRPCRequest{
		Creator: from,
		Lock:    60,
		Requests: []*wasmtypes.InvokeRequest{
			{
				ModuleName:   "wasm",
				ContractName: name,
				MethodName:   method,
				Args:         args,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return wasmtypes.NewMsgInvoke(from, res.Inputs, res.Outputs, res.InputsExt, res.OutputsExt, res.Requests, desc), nil
}

func convertToArgs(args string) (string, error) {
	args1 := make(map[string]string)
	if err := json.Unmarshal([]byte(args), &args1); err != nil {
		return "", err
	}
	args2 := make(map[string][]byte)
	for k, v := range args1 {
		args2[k] = []byte(v)
	}
	bts, err := json.Marshal(args2)
	return string(bts), err
}
