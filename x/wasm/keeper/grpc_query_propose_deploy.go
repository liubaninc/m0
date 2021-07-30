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

func (k Keeper) ProposeDeployAll(c context.Context, req *types.QueryAllProposeDeployRequest) (*types.QueryAllProposeDeployResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proposeDeploys []*types.ProposeDeploy
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	proposeDeployStore := prefix.NewStore(store, types.KeyPrefix(types.ProposeDeployKey))

	pageRes, err := query.Paginate(proposeDeployStore, req.Pagination, func(key []byte, value []byte) error {
		var proposeDeploy types.ProposeDeploy
		if err := k.cdc.UnmarshalBinaryBare(value, &proposeDeploy); err != nil {
			return err
		}

		proposeDeploys = append(proposeDeploys, &proposeDeploy)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProposeDeployResponse{ProposeDeploy: proposeDeploys, Pagination: pageRes}, nil
}

func (k Keeper) ProposeDeploy(c context.Context, req *types.QueryGetProposeDeployRequest) (*types.QueryGetProposeDeployResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProposeDeploy(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetProposeDeployResponse{ProposeDeploy: &val}, nil
}
