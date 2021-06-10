package keeper

import (
	"encoding/json"
	"time"

	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// preExec the contract model uses previous execution to generate RWSets
func PreExec(ctx sdk.Context, k Keeper, req *types.InvokeRPCRequest) (*types.InvokeRPCResponse, error) {
	t := time.Now()
	defer func() {
		k.Logger(ctx).Error("PreExec", "elapsed", time.Now().Sub(t).String())
	}()
	xmreader := NewXMReader(ctx, k)
	modelCache, err := NewXModelCache(xmreader, NewUtxoReader(ctx, k))
	if err != nil {
		return nil, err
	}

	contextConfig := &contract.ContextConfig{
		Logger:         k.Logger(ctx),
		XMCache:        modelCache,
		XMReader:       xmreader,
		Initiator:      req.Creator,
		ContractName:   "",
		ResourceLimits: contract.MaxLimits,
		Core:           newContractChainCore(),
	}

	gasPrice := &xmodel.GasPrice{}
	var requests []*types.InvokeRequest
	var responses []*xmodel.ContractResponse
	var response [][]byte
	var gasUsed = int64(0)
	for _, tmpReq := range req.Requests {
		if tmpReq == nil {
			continue
		}
		if tmpReq.ModuleName == "" && tmpReq.ContractName == "" && tmpReq.MethodName == "" {
			continue
		}

		vm, err := k.vmMgr.GetVM(tmpReq.ModuleName)
		if err != nil {
			return nil, err
		}

		contextConfig.ContractName = tmpReq.ContractName
		if !tmpReq.Amount.Empty() {
			contextConfig.TransferAmount = tmpReq.Amount.String()
		} else {
			contextConfig.TransferAmount = ""
		}
		vmCtx, err := vm.NewContext(contextConfig)
		if err != nil {
			k.Logger(ctx).Error("PreExec NewContext error", "error", err,
				"contractName", tmpReq.ContractName)
			return nil, err
		}

		var args map[string][]byte
		if err := json.Unmarshal([]byte(tmpReq.Args), &args); err != nil {
			return nil, err
		}
		res, err := vmCtx.Invoke(tmpReq.MethodName, args)
		if err != nil {
			k.Logger(ctx).Error("PreExec Invoke error", "error", err, "contractName", tmpReq.ContractName)
			vmCtx.Release()
			return nil, err
		}
		//if res.Status >= 400 {
		//	k.Logger(ctx).Error("PreExec Invoke error", "status", res.Status, "contractName", tmpReq.ContractName)
		//	vmCtx.Release()
		//	return nil, errors.New(res.Message)
		//}
		response = append(response, res.Body)
		responses = append(responses, contract.ToPBContractResponse(res))

		resourceUsed := vmCtx.ResourceUsed()
		gasUsed += resourceUsed.TotalGas(gasPrice)

		request := *tmpReq
		request.ResourceLimits = contract.ToPbLimits(resourceUsed)
		requests = append(requests, &request)
		vmCtx.Release()
	}

	inputs, outputs := modelCache.GetUtxoRWSets()

	err = modelCache.WriteTransientBucket()
	if err != nil {
		return nil, err
	}

	inputsExt, outputsExt, err := modelCache.GetRWSets()
	if err != nil {
		return nil, err
	}
	rsps := &types.InvokeRPCResponse{
		InputsExt:  types.GetInputsExt(inputsExt),
		OutputsExt: types.GetOutputsExt(outputsExt),
		Response:   response,
		Requests:   requests,
		GasUsed:    gasUsed,
		Responses:  responses,
		Inputs:     inputs,
		Outputs:    outputs,
	}
	return rsps, nil
}
