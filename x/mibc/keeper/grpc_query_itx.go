package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/liubaninc/m0/x/mibc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ItxAll(c context.Context, req *types.QueryAllItxRequest) (*types.QueryAllItxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var itxs []*types.Itx
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	itxStore := prefix.NewStore(store, types.KeyPrefix(types.ItxKey))

	pageRes, err := query.Paginate(itxStore, req.Pagination, func(key []byte, value []byte) error {
		var itx types.Itx
		if err := k.cdc.UnmarshalBinaryBare(value, &itx); err != nil {
			return err
		}

		itxs = append(itxs, &itx)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllItxResponse{Itx: itxs, Pagination: pageRes}, nil
}

func (k Keeper) Itx(c context.Context, req *types.QueryGetItxRequest) (*types.QueryGetItxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var itx types.Itx
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasItx(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItxKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetItxIDBytes(req.Id)), &itx)

	return &types.QueryGetItxResponse{Itx: &itx}, nil
}
