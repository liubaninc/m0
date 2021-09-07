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

func (k Keeper) ShareEncryptStorageAll(c context.Context, req *types.QueryAllShareEncryptStorageRequest) (*types.QueryAllShareEncryptStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var shareEncryptStorages []*types.ShareEncryptStorage
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	shareEncryptStorageStore := prefix.NewStore(store, types.KeyPrefix(types.ShareEncryptStorageKey))

	pageRes, err := query.Paginate(shareEncryptStorageStore, req.Pagination, func(key []byte, value []byte) error {
		var shareEncryptStorage types.ShareEncryptStorage
		if err := k.cdc.UnmarshalBinaryBare(value, &shareEncryptStorage); err != nil {
			return err
		}

		shareEncryptStorages = append(shareEncryptStorages, &shareEncryptStorage)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllShareEncryptStorageResponse{ShareEncryptStorage: shareEncryptStorages, Pagination: pageRes}, nil
}

func (k Keeper) ShareEncryptStorage(c context.Context, req *types.QueryGetShareEncryptStorageRequest) (*types.QueryGetShareEncryptStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetShareEncryptStorage(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetShareEncryptStorageResponse{ShareEncryptStorage: &val}, nil
}
