package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/liubaninc/m0/x/validator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValidatorAll(c context.Context, req *types.QueryAllValidatorRequest) (*types.QueryAllValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validatos []*types.Validator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	validatoStore := prefix.NewStore(store, types.KeyPrefix(types.ValidatorKey))

	pageRes, err := query.Paginate(validatoStore, req.Pagination, func(key []byte, value []byte) error {
		var validato types.Validator
		if err := k.cdc.UnmarshalBinaryBare(value, &validato); err != nil {
			return err
		}

		validatos = append(validatos, &validato)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValidatorResponse{Validator: validatos, Pagination: pageRes}, nil
}

func (k Keeper) Validator(c context.Context, req *types.QueryGetValidatorRequest) (*types.QueryGetValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validato types.Validator
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasValidator(ctx, req.Address) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get([]byte(req.Address)), &validato)

	return &types.QueryGetValidatorResponse{Validator: &validato}, nil
}
