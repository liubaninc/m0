package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/liubaninc/m0/x/authority/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Account(c context.Context, req *types.QueryAccountRequest) (*types.QueryAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	_, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	// get account
	acc := k.GetAccount(ctx, req.Address)
	return &types.QueryAccountResponse{
		Account: &acc,
	}, nil
}

func (k Keeper) AllAccount(c context.Context, req *types.QueryAllAccountRequest) (*types.QueryAllAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accs []*types.Account
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var acc types.Account
		if err := k.cdc.UnmarshalBinaryBare(value, &acc); err != nil {
			return err
		}

		accs = append(accs, &acc)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAccountResponse{
		Account:    accs,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) PendingAccount(c context.Context, req *types.QueryPendingAccountRequest) (*types.QueryPendingAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	_, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// get pending account
	pendAcc := k.GetPendingAccount(ctx, req.Address)

	return &types.QueryPendingAccountResponse{
		PendingAccount: &pendAcc,
	}, nil
}

func (k Keeper) AllPendingAccount(c context.Context, req *types.QueryAllPendingAccountRequest) (*types.QueryAllPendingAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accs []*types.PendingAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var acc types.PendingAccount
		if err := k.cdc.UnmarshalBinaryBare(value, &acc); err != nil {
			return err
		}

		accs = append(accs, &acc)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPendingAccountResponse{
		PendingAccount: accs,
		Pagination:     pageRes,
	}, nil
}

func (k Keeper) PendingAccountRevocation(c context.Context, req *types.QueryPendingAccountRevocationRequest) (*types.QueryPendingAccountRevocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	_, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// get pending account revocation
	revoc := k.GetPendingAccountRevocation(ctx, req.Address)

	return &types.QueryPendingAccountRevocationResponse{
		PendingAccountRevocation: &revoc,
	}, nil
}

func (k Keeper) AllPendingAccountRevocation(c context.Context, req *types.QueryAllPendingAccountRevocationRequest) (*types.QueryAllPendingAccountRevocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	var accs []*types.PendingAccountRevocation
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingAccountRevocationPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var acc types.PendingAccountRevocation
		if err := k.cdc.UnmarshalBinaryBare(value, &acc); err != nil {
			return err
		}

		accs = append(accs, &acc)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPendingAccountRevocationResponse{
		PendingAccountRevocation: accs,
		Pagination:               pageRes,
	}, nil
}
