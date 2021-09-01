package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/liubaninc/m0/x/wasm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ApproveDeployAll(c context.Context, req *types.QueryAllApproveDeployRequest) (*types.QueryAllApproveDeployResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var approveDeploys []*types.ApproveDeploy
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	approveDeployStore := prefix.NewStore(store, types.KeyPrefix(types.ApproveDeployKey))

	pageRes, err := query.Paginate(approveDeployStore, req.Pagination, func(key []byte, value []byte) error {
		var approveDeploy types.ApproveDeploy
		if err := k.cdc.UnmarshalBinaryBare(value, &approveDeploy); err != nil {
			return err
		}

		approveDeploys = append(approveDeploys, &approveDeploy)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllApproveDeployResponse{ApproveDeploy: approveDeploys, Pagination: pageRes}, nil
}

func (k Keeper) ApproveDeploy(c context.Context, req *types.QueryGetApproveDeployRequest) (*types.QueryGetApproveDeployResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetApproveDeploy(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetApproveDeployResponse{ApproveDeploy: &val}, nil
}
