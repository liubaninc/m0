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

func (k Keeper) EncryptStorageAll(c context.Context, req *types.QueryAllEncryptStorageRequest) (*types.QueryAllEncryptStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var encryptStorages []*types.EncryptStorage
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	encryptStorageStore := prefix.NewStore(store, types.KeyPrefix(types.EncryptStorageKey))

	pageRes, err := query.Paginate(encryptStorageStore, req.Pagination, func(key []byte, value []byte) error {
		var encryptStorage types.EncryptStorage
		if err := k.cdc.UnmarshalBinaryBare(value, &encryptStorage); err != nil {
			return err
		}

		encryptStorages = append(encryptStorages, &encryptStorage)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEncryptStorageResponse{EncryptStorage: encryptStorages, Pagination: pageRes}, nil
}

func (k Keeper) EncryptStorage(c context.Context, req *types.QueryGetEncryptStorageRequest) (*types.QueryGetEncryptStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEncryptStorage(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetEncryptStorageResponse{EncryptStorage: &val}, nil
}
