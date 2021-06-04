package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/gogo/protobuf/proto"
	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/bridge"
	"github.com/liubaninc/m0/x/wasm/xmodel/contract/kernel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ContractAll(c context.Context, req *types.QueryAllContractRequest) (*types.QueryAllContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var items []*types.Contract
	store := ctx.KVStore(k.storeKey)
	keyPrefix := append([]byte(types.ExtUtxoTablePrefix), types.MakeRawKey(kernel.Contract2AccountBucket, nil)...)
	tokenStore := prefix.NewStore(store, types.KeyPrefix(string(keyPrefix)))

	pageRes, err := query.Paginate(tokenStore, req.Pagination, func(key []byte, value []byte) error {
		verData, err := k.GetVersionedData(ctx, value)
		if err != nil {
			panic(err)
		}
		account := string(verData.PureData.Value)
		name := string(key[len(keyPrefix):])

		verData, err = k.GetLastVersionedData(ctx, "contract", bridge.ContractCodeDescKey(name))
		if err != nil {
			panic(err)
		}
		var desc xmodel.WasmCodeDesc
		if err := proto.Unmarshal(verData.PureData.Value, &desc); err != nil {
			panic(err)
		}
		verData, err = k.GetLastVersionedData(ctx, "contract", bridge.ContractNumberKey(name))
		if err != nil {
			panic(err)
		}
		number, err := strconv.ParseUint(string(verData.PureData.Value), 10, 64)
		_ = number
		items = append(items, &types.Contract{
			Name:         name,
			Initiator:    account,
			Desc: &desc,
		})
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractResponse{Contract: items, Pagination: pageRes}, nil
}

func (k Keeper) AccountContractAll(c context.Context, req *types.QueryGetAccountAllContractRequest) (*types.QueryGetAccountAllContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var items []*types.Contract
	store := ctx.KVStore(k.storeKey)
	keyPrefix :=append([]byte(types.ExtUtxoTablePrefix), types.MakeRawKey(kernel.Account2ContractBucket, []byte(req.Name+kernel.Account2ContractSeparator))...)
	tokenStore := prefix.NewStore(store, types.KeyPrefix(string(keyPrefix)))

	pageRes, err := query.Paginate(tokenStore, req.Pagination, func(key []byte, value []byte) error {
		verData, err := k.GetVersionedData(ctx, value)
		if err != nil {
			panic(err)
		}
		account := string(verData.PureData.Value)
		name := string(key[len(keyPrefix):])

		verData, err = k.GetLastVersionedData(ctx, "contract", bridge.ContractCodeDescKey(name))
		if err != nil {
			panic(err)
		}
		var desc xmodel.WasmCodeDesc
		if err := proto.Unmarshal(verData.PureData.Value, &desc); err != nil {
			panic(err)
		}
		verData, err = k.GetLastVersionedData(ctx, "contract", bridge.ContractNumberKey(name))
		if err != nil {
			panic(err)
		}
		number, err := strconv.ParseUint(string(verData.PureData.Value), 10, 64)
		_ = number
		items = append(items, &types.Contract{
			Name:         name,
			Initiator:    account,
			Desc: &desc,
		})
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetAccountAllContractResponse{Contract: items, Pagination: pageRes}, nil
}

func (k Keeper) Contract(c context.Context, req *types.QueryGetContractRequest) (*types.QueryGetContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, err := k.GetContract(ctx, req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if val == nil {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetContractResponse{Contract: val}, nil
}
