package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/liubaninc/m0/x/blob/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BlobAll(c context.Context, req *types.QueryAllBlobRequest) (*types.QueryAllBlobResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var blobs []*types.Blob
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	blobStore := prefix.NewStore(store, types.KeyPrefix(types.BlobKey))

	pageRes, err := query.Paginate(blobStore, req.Pagination, func(key []byte, value []byte) error {
		var blob types.Blob
		if err := k.cdc.UnmarshalBinaryBare(value, &blob); err != nil {
			return err
		}

		blobs = append(blobs, &blob)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBlobResponse{Blob: blobs, Pagination: pageRes}, nil
}

func (k Keeper) Blob(c context.Context, req *types.QueryGetBlobRequest) (*types.QueryGetBlobResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var blob types.Blob
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasBlob(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlobKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBlobIDBytes(req.Id)), &blob)

	return &types.QueryGetBlobResponse{Blob: &blob}, nil
}
