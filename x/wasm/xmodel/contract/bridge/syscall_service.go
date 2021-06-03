package bridge

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/liubaninc/m0/x/wasm/xmodel"

	"github.com/liubaninc/m0/x/wasm/xmodel/contract"
	"github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/go/pb"
)

var (
	// ErrOutOfDiskLimit define OutOfDiskLimit Error
	ErrOutOfDiskLimit = errors.New("out of disk limit")
)

const (
	// DefaultCap define default cap of NewIterator
	DefaultCap = 1000
	// MaxContractCallDepth define max contract call depth
	MaxContractCallDepth = 10
)

// VmManager define the virtual machine interface
type VmManager interface {
	GetVirtualMachine(name string) (contract.VirtualMachine, bool)
}

// SyscallService is the handler of contract syscalls
type SyscallService struct {
	ctxmgr *ContextManager
	vmm    VmManager
}

// NewSyscallService instances a new SyscallService
func NewSyscallService(ctxmgr *ContextManager, vmm VmManager) *SyscallService {
	return &SyscallService{
		ctxmgr: ctxmgr,
		vmm:    vmm,
	}
}

// Ping implements Syscall interface
func (c *SyscallService) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	return new(pb.PingResponse), nil
}

// QueryBlock implements Syscall interface
func (c *SyscallService) QueryBlock(ctx context.Context, in *pb.QueryBlockRequest) (*pb.QueryBlockResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().Ctxid)
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}

	blocksdk, err := nctx.Core.QueryBlock(in.Height)
	if err != nil {
		nctx.Logger.Error("QueryBlock", "contract", nctx.ContractName, "height", in.Height, "error", err)
		return nil, err
	}

	return &pb.QueryBlockResponse{
		Block: blocksdk,
	}, nil
}

// QueryTx implements Syscall interface
func (c *SyscallService) QueryTx(ctx context.Context, in *pb.QueryTxRequest) (*pb.QueryTxResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().Ctxid)
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}

	rawTxid, err := hex.DecodeString(in.Txid)
	if err != nil {
		return nil, err
	}

	txsdk, err := nctx.Core.QueryTransaction(rawTxid)
	if err != nil {
		nctx.Logger.Error("QueryTx", "contract", nctx.ContractName, "hash", in.Txid, "error", err)
		return nil, err
	}

	return &pb.QueryTxResponse{
		Tx: txsdk,
	}, nil
}

// Transfer implements Syscall interface
func (c *SyscallService) Transfer(ctx context.Context, in *pb.TransferRequest) (*pb.TransferResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().Ctxid)
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}
	from := authtypes.NewModuleAddress(nctx.ContractName).String()
	err := nctx.Cache.Transfer(from, in.GetTo(), in.GetAmount())
	if err != nil {
		nctx.Logger.Error("Transfer", "contract", nctx.ContractName, "from", from, "to", in.To, "amount", in.Amount, "error", err)
		return nil, err
	}
	resp := &pb.TransferResponse{}
	return resp, nil
}

// ContractCall implements Syscall interface
func (c *SyscallService) ContractCall(ctx context.Context, in *pb.ContractCallRequest) (*pb.ContractCallResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().Ctxid)
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}
	if nctx.ContractSet[in.GetContract()] {
		nctx.Logger.Error("ContractCall", "contract", nctx.ContractName, "error", "recursive contract call not permitted")
		return nil, errors.New("recursive contract call not permitted")
	}

	if len(nctx.ContractSet) >= MaxContractCallDepth {
		nctx.Logger.Error("ContractCall", "contract", nctx.ContractName, "error", "max contract call depth exceeds")
		return nil, errors.New("max contract call depth exceeds")
	}

	vm, ok := c.vmm.GetVirtualMachine(in.GetModule())
	if !ok {
		nctx.Logger.Error("ContractCall", "contract", nctx.ContractName, "module", in.Module, "error", "module not found")
		return nil, errors.New("module not found")
	}
	currentUsed := nctx.ResourceUsed()
	limits := new(contract.Limits).Add(nctx.ResourceLimits).Sub(currentUsed)
	// disk usage is shared between all context
	limits.Disk = nctx.ResourceLimits.Disk

	args := make(map[string][]byte)
	for _, arg := range in.GetArgs() {
		args[arg.GetKey()] = arg.GetValue()
	}

	nctx.ContractSet[in.GetContract()] = true
	cfg := &contract.ContextConfig{
		ContractName:   in.GetContract(),
		XMCache:        nctx.Cache,
		CanInitialize:  false,
		AuthRequire:    nctx.AuthRequire,
		Initiator:      nctx.Initiator,
		Core:           nctx.Core,
		ResourceLimits: *limits,
		ContractSet:    nctx.ContractSet,
	}
	vctx, err := vm.NewContext(cfg)
	if err != nil {
		nctx.Logger.Error("ContractCall", "contract", nctx.ContractName, "module", in.Module, "error", err)
		return nil, err
	}
	defer func() {
		vctx.Release()
		delete(nctx.ContractSet, in.GetContract())
	}()

	vresp, err := vctx.Invoke(in.GetMethod(), args)
	if err != nil {
		nctx.Logger.Error("ContractCall", "contract", nctx.ContractName, "module", in.Module, "error", err)
		return nil, err
	}
	nctx.SubResourceUsed.Add(vctx.ResourceUsed())

	return &pb.ContractCallResponse{
		Response: &pb.Response{
			Status:  int32(vresp.Status),
			Message: vresp.Message,
			Body:    vresp.Body,
		}}, nil
}

// PutObject implements Syscall interface
func (c *SyscallService) PutObject(ctx context.Context, in *pb.PutRequest) (*pb.PutResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().Ctxid)
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}
	if in.Value == nil {
		nctx.Logger.Error("PutObject", "contract", nctx.ContractName, "key", in.Key, "error", "put nil value")
		return nil, errors.New("put nil value")
	}

	err := nctx.Cache.Put(nctx.ContractName, in.Key, in.Value)
	if err != nil {
		nctx.Logger.Error("PutObject", "contract", nctx.ContractName, "key", in.Key, "value", in.Value, "error", err)
		return nil, err
	}

	return &pb.PutResponse{}, nil
}

// GetObject implements Syscall interface
func (c *SyscallService) GetObject(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().Ctxid)
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}

	value, err := nctx.Cache.Get(nctx.ContractName, in.Key)
	if err != nil {
		nctx.Logger.Error("GetObject", "contract", nctx.ContractName, "key", in.Key, "error", err)
		return nil, err
	}
	return &pb.GetResponse{
		Value: value.GetPureData().GetValue(),
	}, nil
}

// DeleteObject implements Syscall interface
func (c *SyscallService) DeleteObject(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().Ctxid)
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}

	err := nctx.Cache.Del(nctx.ContractName, in.Key)
	if err != nil {
		nctx.Logger.Error("DeleteObject", "contract", nctx.ContractName, "key", in.Key, "error", err)
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

// NewIterator implements Syscall interface
func (c *SyscallService) NewIterator(ctx context.Context, in *pb.IteratorRequest) (*pb.IteratorResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().Ctxid)
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}

	limit := in.Cap
	if limit <= 0 {
		limit = DefaultCap
	}
	iter, err := nctx.Cache.Select(nctx.ContractName, in.Start, in.Limit)
	if err != nil {
		return nil, err
	}
	out := new(pb.IteratorResponse)
	for iter.Next() && limit > 0 {
		out.Items = append(out.Items, &pb.IteratorItem{
			Key:   append([]byte(""), iter.Data().GetPureData().GetKey()...), //make a copy
			Value: append([]byte(""), iter.Data().GetPureData().GetValue()...),
		})
		limit -= 1
	}
	if iter.Error() != nil {
		nctx.Logger.Error("NewIterator", "contract", nctx.ContractName, "start", string(in.Start), "limit", string(in.Limit), "cap", limit, "error", iter.Error())
		return nil, iter.Error()
	}
	iter.Release()
	return out, nil
}

// GetCallArgs implements Syscall interface
func (c *SyscallService) GetCallArgs(ctx context.Context, in *pb.GetCallArgsRequest) (*pb.CallArgs, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().Ctxid)
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}
	var args []*pb.ArgPair
	for key, value := range nctx.Args {
		args = append(args, &pb.ArgPair{
			Key:   key,
			Value: value,
		})
	}
	sort.Slice(args, func(i, j int) bool {
		return args[i].Key < args[j].Key
	})
	return &pb.CallArgs{
		Method:         nctx.Method,
		Args:           args,
		Initiator:      nctx.Initiator,
		AuthRequire:    nctx.AuthRequire,
		TransferAmount: nctx.TransferAmount,
	}, nil
}

// SetOutput implements Syscall interface
func (c *SyscallService) SetOutput(ctx context.Context, in *pb.SetOutputRequest) (*pb.SetOutputResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.Header.Ctxid)
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}
	nctx.Output = in.GetResponse()
	return new(pb.SetOutputResponse), nil
}

// PostLog handle log entry from contract
func (c *SyscallService) PostLog(ctx context.Context, in *pb.PostLogRequest) (*pb.PostLogResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().GetCtxid())
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.Header.Ctxid)
	}
	nctx.Logger.Info(in.GetEntry())
	return &pb.PostLogResponse{}, nil
}

// PostLog handle log entry from contract
func (c *SyscallService) EmitEvent(ctx context.Context, in *pb.EmitEventRequest) (*pb.EmitEventResponse, error) {
	nctx, ok := c.ctxmgr.Context(in.GetHeader().GetCtxid())
	if !ok {
		return nil, fmt.Errorf("bad ctx id:%d", in.GetHeader().GetCtxid())
	}
	event := &xmodel.ContractEvent{
		Contract: nctx.ContractName,
		Name:     in.GetName(),
		Body:     in.GetBody(),
	}
	nctx.Events = append(nctx.Events, event)
	nctx.Cache.AddEvent(event)
	return &pb.EmitEventResponse{}, nil
}
