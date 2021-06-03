package keeper

import (
	"context"
	"github.com/liubaninc/m0/x/wasm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ContractAll(c context.Context, req *types.QueryAllContractRequest) (*types.QueryAllContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	return &types.QueryAllContractResponse{}, nil
}

func (k Keeper) AccountContractAll(c context.Context, req *types.QueryGetAccountAllContractRequest) (*types.QueryGetAccountAllContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	return &types.QueryGetAccountAllContractResponse{}, nil
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
