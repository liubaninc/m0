package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/validator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// Params queries the staking parameters
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// HistoricalInfo queries the historical info for given height
func (k Keeper) HistoricalInfo(c context.Context, req *types.QueryHistoricalInfoRequest) (*types.QueryHistoricalInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.Height < 0 {
		return nil, status.Error(codes.InvalidArgument, "height cannot be negative")
	}
	ctx := sdk.UnwrapSDKContext(c)
	hi, found := k.GetHistoricalInfo(ctx, req.Height)
	if !found {
		return nil, status.Errorf(codes.NotFound, "historical info for height %d not found", req.Height)
	}

	return &types.QueryHistoricalInfoResponse{Hist: &hi}, nil
}
