package keeper

import (
	"errors"
	"fmt"
	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Env data structure for read/write sets environment
type Env struct {
	outputs    []*xmodel.PureData
	modelCache *XMCache
}

func NewEnv(outputs []*xmodel.PureData, modelCache *XMCache) *Env {
	return &Env{
		outputs:    outputs,
		modelCache: modelCache,
	}
}

// GetModelCache get instance of model cache
func (e *Env) GetModelCache() *XMCache {
	if e != nil {
		return e.modelCache
	}
	return nil
}

// GetOutputs get outputs
func (e *Env) GetOutputs() []*xmodel.PureData {
	if e != nil {
		return e.outputs
	}
	return nil
}

// prepareEnv get read/write sets with transaction
func prepareEnv(ctx sdk.Context, k Keeper, msg *types.MsgInvoke) (*Env, error) {
	var inputsExt []*xmodel.VersionedData
	var outputsExt []*xmodel.PureData

	for _, inputExt := range msg.InputsExt {
		var verData *xmodel.VersionedData
		var err error
		verData, err = k.GetLastVersionedData(ctx, inputExt.Bucket, []byte(inputExt.Key))
		if err != nil {
			return nil, err
		}
		if verData == nil {
			verData = types.MakeEmptyVersionedData(inputExt.Bucket, []byte(inputExt.Key))
		}
		localVer := types.MakeVersion(verData.RefTxid, verData.RefMsgOffset, verData.RefOffset)
		remoteVer := types.MakeVersion([]byte(inputExt.RefTx), inputExt.RefMsg, inputExt.RefOffset)
		if localVer != remoteVer {
			return nil, sdkerrors.Wrapf(types.ErrUnexpected, "prepareEnv fail, key:%s, inputs version is not valid: %s != %s", string(verData.PureData.Key), localVer, remoteVer)
		}
		inputsExt = append(inputsExt, verData)
	}
	for _, outputExt := range msg.OutputsExt {
		outputsExt = append(outputsExt, &xmodel.PureData{
			Bucket: outputExt.Bucket,
			Key:    []byte(outputExt.Key),
			Value:  outputExt.Value,
		})
	}
	contractInputs, err := types.ParseInputsFromExt(msg.OutputsExt)
	if err != nil {
		return nil, err
	}

	modelCache := NewXModelCacheWithInputs(inputsExt, contractInputs)
	env := NewEnv(outputsExt, modelCache)
	k.Logger(ctx).Info("PrepareEnv done!", "env", env)
	return env, nil
}

func VerifyTxRWSets(ctx sdk.Context, k Keeper, msg *types.MsgInvoke) (bool, error) {
	if len(msg.ContractRequests) == 0 {
		return true, nil
	}

	env, err := prepareEnv(ctx, k, msg)
	if err != nil {
		return false, err
	}

	contextConfig := &contract.ContextConfig{
		Logger:       k.Logger(ctx),
		XMCache:      env.GetModelCache(),
		XMReader:     NewXMReader(ctx, k),
		Initiator:    msg.Creator,
		ContractName: "",
		Core:         newContractChainCore(),
	}

	defer func() {
		if r := recover(); r != nil {
			switch rType := r.(type) {
			case sdk.ErrorOutOfGas:
				err = sdkerrors.Wrap(sdkerrors.ErrOutOfGas, rType.Descriptor)
			default:
				panic(r)
			}
		}
	}()
	// get gas rate to utxo
	gasPrice := &xmodel.GasPrice{}
	for _, tmpReq := range msg.ContractRequests {
		vm, err := k.vmMgr.GetVM(tmpReq.ModuleName)
		if err != nil {
			return false, err
		}

		resourceLimits := contract.FromPbLimits(tmpReq.ResourceLimits)
		gas := resourceLimits.TotalGas(gasPrice)
		ctx.GasMeter().ConsumeGas(sdk.Gas(gas), fmt.Sprintf("out of gas contractName %s; gasWanted: %d, gasUsed: %d", tmpReq.ContractName, gas, ctx.GasMeter().GasConsumed()))

		contextConfig.ResourceLimits = resourceLimits
		contextConfig.ContractName = tmpReq.ContractName
		if !tmpReq.Amount.Empty() {
			contextConfig.TransferAmount = tmpReq.Amount.String()
		} else {
			contextConfig.TransferAmount = ""
		}

		vmCtx, err := vm.NewContext(contextConfig)
		if err != nil {
			k.Logger(ctx).Error("verifyTxRWSets NewContext error", "err", err, "contractName", tmpReq.ContractName)
			return false, err
		}

		ctxResponse, ctxErr := vmCtx.Invoke(tmpReq.MethodName, tmpReq.Args)
		if ctxErr != nil {
			vmCtx.Release()
			k.Logger(ctx).Error("verifyTxRWSets Invoke error", "error", ctxErr, "contractName", tmpReq.ContractName)
			return false, ctxErr
		}
		// 判断合约调用的返回码
		if ctxResponse.Status >= 400 {
			vmCtx.Release()
			k.Logger(ctx).Error("verifyTxRWSets Invoke error", "status", ctxResponse.Status, "contractName", tmpReq.ContractName)
			return false, errors.New(ctxResponse.Message)
		}

		vmCtx.Release()
	}

	err = env.GetModelCache().WriteTransientBucket()
	if err != nil {
		return false, err
	}

	_, writeSet, err := env.GetModelCache().GetRWSets()
	if err != nil {
		return false, err
	}
	ok := types.Equal(env.GetOutputs(), writeSet)
	if !ok {
		return false, fmt.Errorf("write set not equal")
	}
	return true, nil
}
