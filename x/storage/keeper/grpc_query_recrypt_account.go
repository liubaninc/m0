package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/liubaninc/m0/x/storage/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RecryptAccountAll(c context.Context, req *types.QueryAllRecryptAccountRequest) (*types.QueryAllRecryptAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var recryptAccounts []*types.RecryptAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	recryptAccountStore := prefix.NewStore(store, types.KeyPrefix(types.RecryptAccountKey))

	pageRes, err := query.Paginate(recryptAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var recryptAccount types.RecryptAccount
		if err := k.cdc.UnmarshalBinaryBare(value, &recryptAccount); err != nil {
			return err
		}

		recryptAccounts = append(recryptAccounts, &recryptAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRecryptAccountResponse{RecryptAccount: recryptAccounts, Pagination: pageRes}, nil
}

func (k Keeper) RecryptAccount(c context.Context, req *types.QueryGetRecryptAccountRequest) (*types.QueryGetRecryptAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRecryptAccount(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetRecryptAccountResponse{RecryptAccount: &val}, nil
}
