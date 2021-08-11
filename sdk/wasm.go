package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	wasmtypes "github.com/liubaninc/m0/x/wasm/types"

	"github.com/cosmos/cosmos-sdk/types/query"
)

func (c Client) GetContract(name string) (*wasmtypes.QueryGetContractResponse, error) {
	if err := kernel.ValidContractName(name); err != nil {
		return nil, fmt.Errorf("invalid contract name %v (%v)", name, err)
	}

	queryClient := wasmtypes.NewQueryClient(c)
	res, err := queryClient.Contract(context.Background(), &wasmtypes.QueryGetContractRequest{
		Name: name,
	})
	return res, err
}

func (c Client) GetContracts(key []byte, offset uint64, limit uint64, countTotal bool) (*wasmtypes.QueryAllContractResponse, error) {
	queryClient := wasmtypes.NewQueryClient(c)
	res, err := queryClient.ContractAll(context.Background(), &wasmtypes.QueryAllContractRequest{
		Pagination: &query.PageRequest{
			Key:        key,
			Offset:     offset,
			Limit:      limit,
			CountTotal: countTotal,
		},
	})
	return res, err
}

func (c Client) GetAccountContracts(address string, key []byte, offset uint64, limit uint64, countTotal bool) (*wasmtypes.QueryGetAccountAllContractResponse, error) {
	if _, err := sdk.AccAddressFromBech32(address); err != nil {
		return nil, fmt.Errorf("invalid address %s (%s)", address, err)
	}

	queryClient := wasmtypes.NewQueryClient(c)
	res, err := queryClient.AccountContractAll(context.Background(), &wasmtypes.QueryGetAccountAllContractRequest{
		Name: address,
		Pagination: &query.PageRequest{
			Key:        key,
			Offset:     offset,
			Limit:      limit,
			CountTotal: countTotal,
		},
	})
	return res, err
}

func (c Client) Query(name string, method string, arg string) (*xmodel.ContractResponse, error) {
	if err := kernel.ValidContractName(name); err != nil {
		return nil, fmt.Errorf("invalid name %v (%v)", name, err)
	}
	if len(method) == 0 {
		return nil, fmt.Errorf("invalid method (empty)")
	}
	methodArgs, err := convertToArgs(arg)
	if err != nil {
		return nil, fmt.Errorf("invalid arg %v (%v)", arg, err)
	}
	mArgsStr, _ := json.Marshal(methodArgs)

	queryClient := wasmtypes.NewQueryClient(c)
	res, err := queryClient.PreExec(context.Background(), &wasmtypes.InvokeRPCRequest{
		Requests: []*wasmtypes.InvokeRequest{
			{
				ModuleName:   "wasm",
				ContractName: name,
				MethodName:   method,
				Args:         string(mArgsStr),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return res.Responses[0], nil
}

func (c Client) DeployMsg(from string, name string, code []byte, args string, desc string, fees string) (sdk.Msg, error) {
	if _, err := sdk.AccAddressFromBech32(from); err != nil {
		return nil, fmt.Errorf("invalid from %s (%s)", from, err)
	}
	if err := kernel.ValidContractName(name); err != nil {
		return nil, fmt.Errorf("invalid name %v (%v)", name, err)
	}
	//code, err := ioutil.ReadFile(codeFile)
	//if err != nil {
	//	return nil, fmt.Errorf("invalid code file %v (%v)", codeFile, err)
	//}
	initArgs, err := convertToArgs(args)
	if err != nil {
		return nil, fmt.Errorf("invalid args %v (%v)", args, err)
	}
	initArgsStr, _ := json.Marshal(initArgs)

	feeCoins, err := sdk.ParseCoinsNormalized(fees)
	if err != nil {
		return nil, fmt.Errorf("invalid fees %s (%s)", fees, err)
	}

	codeDesc := &xmodel.WasmCodeDesc{
		Runtime:      "c",
		ContractType: "wasm",
	}

	queryClient := wasmtypes.NewQueryClient(c)
	resp, err := queryClient.PreExec(context.Background(), &wasmtypes.InvokeRPCRequest{
		Creator: from,
		Lock:    c.locked,
		Requests: []*wasmtypes.InvokeRequest{
			wasmtypes.NewMsgDeploy(from, name, code, codeDesc, string(initArgsStr), nil, nil, nil, nil, nil, desc).ConvertInvokeRequest(),
		},
	})
	if err != nil {
		return nil, err
	}

	var inputs []*utxotypes.Input
	var outputs []*utxotypes.Output
	neededTotal := sdk.NewCoins()
	for _, fee := range feeCoins {
		outputs = append(outputs, &utxotypes.Output{
			Amount: fee,
			ToAddr: authtypes.NewModuleAddress(authtypes.FeeCollectorName).String(),
		})
		neededTotal = neededTotal.Add(fee)
	}
	if !neededTotal.IsZero() {
		res, err := c.GetInput(from, neededTotal.String(), c.locked)
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, res.Inputs...)
		changeCoins := res.Amount.Sub(neededTotal)
		for _, changeCoin := range changeCoins {
			outputs = append(outputs, &utxotypes.Output{
				ToAddr: from,
				Amount: changeCoin,
			})
		}
	}

	msg := wasmtypes.NewMsgDeploy(from, name, code, codeDesc, string(initArgsStr), resp.Requests[0].ResourceLimits, append(inputs, resp.Inputs...), append(outputs, resp.Outputs...), resp.InputsExt, resp.OutputsExt, desc)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c Client) UpgradeMsg(from string, name string, code []byte, desc string, fees string) (sdk.Msg, error) {
	if _, err := sdk.AccAddressFromBech32(from); err != nil {
		return nil, fmt.Errorf("invalid from %s (%s)", from, err)
	}
	if err := kernel.ValidContractName(name); err != nil {
		return nil, fmt.Errorf("invalid name %v (%v)", name, err)
	}
	//code, err := ioutil.ReadFile(codeFile)
	//if err != nil {
	//	return nil, fmt.Errorf("invalid code file %v (%v)", codeFile, err)
	//}

	feeCoins, err := sdk.ParseCoinsNormalized(fees)
	if err != nil {
		return nil, fmt.Errorf("invalid fees %s (%s)", fees, err)
	}

	queryClient := wasmtypes.NewQueryClient(c)
	resp, err := queryClient.PreExec(context.Background(), &wasmtypes.InvokeRPCRequest{
		Creator: from,
		Lock:    c.locked,
		Requests: []*wasmtypes.InvokeRequest{
			wasmtypes.NewMsgUpgrade(from, name, code, nil, nil, nil, nil, nil, desc).ConvertInvokeRequest(),
		},
	})
	if err != nil {
		return nil, err
	}

	var inputs []*utxotypes.Input
	var outputs []*utxotypes.Output
	neededTotal := sdk.NewCoins()
	for _, fee := range feeCoins {
		outputs = append(outputs, &utxotypes.Output{
			Amount: fee,
			ToAddr: authtypes.NewModuleAddress(authtypes.FeeCollectorName).String(),
		})
		neededTotal = neededTotal.Add(fee)
	}
	if !neededTotal.IsZero() {
		res, err := c.GetInput(from, neededTotal.String(), c.locked)
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, res.Inputs...)
		changeCoins := res.Amount.Sub(neededTotal)
		for _, changeCoin := range changeCoins {
			outputs = append(outputs, &utxotypes.Output{
				ToAddr: from,
				Amount: changeCoin,
			})
		}
	}

	msg := wasmtypes.NewMsgUpgrade(from, name, code, resp.Requests[0].ResourceLimits, append(inputs, resp.Inputs...), append(outputs, resp.Outputs...), resp.InputsExt, resp.OutputsExt, desc)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c Client) InvokeMsg(from string, name string, method string, args string, desc string, fees string, amounts string) (sdk.Msg, error) {
	if _, err := sdk.AccAddressFromBech32(from); err != nil {
		return nil, fmt.Errorf("invalid from %s (%s)", from, err)
	}
	if err := kernel.ValidContractName(name); err != nil {
		return nil, fmt.Errorf("invalid name %v (%v)", name, err)
	}
	if len(method) == 0 {
		return nil, fmt.Errorf("invalid method %v (%v)", method, "empty")
	}
	initArgs, err := convertToArgs(args)
	if err != nil {
		return nil, fmt.Errorf("invalid args %v (%v)", args, err)
	}
	initArgsStr, _ := json.Marshal(initArgs)

	feeCoins, err := sdk.ParseCoinsNormalized(fees)
	if err != nil {
		return nil, fmt.Errorf("invalid fees %s (%s)", fees, err)
	}

	amountCoins, err := sdk.ParseCoinsNormalized(amounts)
	if err != nil {
		return nil, fmt.Errorf("invalid amount %v", err)
	}

	queryClient := wasmtypes.NewQueryClient(c)
	resp, err := queryClient.PreExec(context.Background(), &wasmtypes.InvokeRPCRequest{
		Creator: from,
		Lock:    c.locked,
		Requests: []*wasmtypes.InvokeRequest{
			{
				Amount:       amountCoins,
				ModuleName:   "wasm",
				ContractName: name,
				MethodName:   method,
				Args:         string(initArgsStr),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var inputs []*utxotypes.Input
	var outputs []*utxotypes.Output
	neededTotal := sdk.NewCoins()
	for _, coin := range amountCoins {
		outputs = append(outputs, &utxotypes.Output{
			Amount: coin,
			ToAddr: authtypes.NewModuleAddress(name).String(),
		})
		neededTotal = neededTotal.Add(coin)
	}
	for _, fee := range feeCoins {
		outputs = append(outputs, &utxotypes.Output{
			Amount: fee,
			ToAddr: authtypes.NewModuleAddress(authtypes.FeeCollectorName).String(),
		})
		neededTotal = neededTotal.Add(fee)
	}
	if !neededTotal.IsZero() {
		res, err := c.GetInput(from, neededTotal.String(), c.locked)
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, res.Inputs...)
		changeCoins := res.Amount.Sub(neededTotal)
		for _, changeCoin := range changeCoins {
			outputs = append(outputs, &utxotypes.Output{
				ToAddr: from,
				Amount: changeCoin,
			})
		}
	}

	msg := wasmtypes.NewMsgInvoke(from, append(inputs, resp.Inputs...), append(outputs, resp.Outputs...), resp.InputsExt, resp.OutputsExt, resp.Requests, desc)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c Client) FreezeMsg(from string, name string) (sdk.Msg, error) {
	if _, err := sdk.AccAddressFromBech32(from); err != nil {
		return nil, fmt.Errorf("invalid from %s (%s)", from, err)
	}
	if err := kernel.ValidContractName(name); err != nil {
		return nil, fmt.Errorf("invalid name %v (%v)", name, err)
	}

	msg := wasmtypes.NewMsgFreeze(from, name)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c Client) UnfreezeMsg(from string, name string) (sdk.Msg, error) {
	if _, err := sdk.AccAddressFromBech32(from); err != nil {
		return nil, fmt.Errorf("invalid from %s (%s)", from, err)
	}
	if err := kernel.ValidContractName(name); err != nil {
		return nil, fmt.Errorf("invalid name %v (%v)", name, err)
	}

	msg := wasmtypes.NewMsgUnfreeze(from, name)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c Client) UndeployMsg(from string, name string) (sdk.Msg, error) {
	if _, err := sdk.AccAddressFromBech32(from); err != nil {
		return nil, fmt.Errorf("invalid from %s (%s)", from, err)
	}
	if err := kernel.ValidContractName(name); err != nil {
		return nil, fmt.Errorf("invalid name %v (%v)", name, err)
	}

	msg := wasmtypes.NewMsgUndeploy(from, name)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
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

func (c Client) BroadcastDeployTx(from string, name string, codeFile []byte, args string, desc string, fees string, memo string) (*sdk.TxResponse, error) {
	msg, err := c.DeployMsg(from, name, codeFile, args, desc, fees)
	if err != nil {
		return nil, err
	}
	return c.GenerateAndBroadcastTx(from, fees, memo, 0, msg)
}
func (c Client) BroadcastUpgradeTx(from string, name string, codeFile []byte, desc string, fees string, memo string) (*sdk.TxResponse, error) {
	msg, err := c.UpgradeMsg(from, name, codeFile, desc, fees)
	if err != nil {
		return nil, err
	}
	return c.GenerateAndBroadcastTx(from, fees, memo, 0, msg)
}
func (c Client) BroadcastInvokeTx(from string, name string, method string, args string, amounts string, desc string, fees string, memo string) (*sdk.TxResponse, error) {
	msg, err := c.InvokeMsg(from, name, method, args, desc, fees, amounts)
	if err != nil {
		return nil, err
	}
	return c.GenerateAndBroadcastTx(from, fees, memo, 0, msg)
}
func (c Client) BroadcastFreezeTx(from string, name string, fees string, memo string) (*sdk.TxResponse, error) {
	msg, err := c.FreezeMsg(from, name)
	if err != nil {
		return nil, err
	}
	return c.GenerateAndBroadcastTx(from, fees, memo, 0, msg)
}
func (c Client) BroadcastUnfreezeTx(from string, name string, fees string, memo string) (*sdk.TxResponse, error) {
	msg, err := c.UnfreezeMsg(from, name)
	if err != nil {
		return nil, err
	}
	return c.GenerateAndBroadcastTx(from, fees, memo, 0, msg)
}
func (c Client) BroadcastUndeployTx(from string, name string, fees string, memo string) (*sdk.TxResponse, error) {
	msg, err := c.UndeployMsg(from, name)
	if err != nil {
		return nil, err
	}
	return c.GenerateAndBroadcastTx(from, fees, memo, 0, msg)
}
