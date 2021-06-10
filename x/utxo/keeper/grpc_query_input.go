package keeper

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/liubaninc/m0/x/utxo/types"
)

func (k Keeper) InputAll(c context.Context, req *types.QueryAllInputRequest) (*types.QueryAllInputResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}
	if err := sdk.ValidateDenom(req.Denom); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid denom")
	}

	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	inputStore := prefix.NewStore(store, types.KeyPrefix(types.InputKey))
	addrStore := prefix.NewStore(inputStore, types.KeyPrefix(fmt.Sprintf("%X_%s", addr.Bytes(), req.Denom)))

	var inputs []*types.Input
	amount := sdk.NewCoin(req.Denom, sdk.ZeroInt())
	pageRes, err := query.Paginate(addrStore, req.Pagination, func(key []byte, value []byte) error {
		var input types.Input
		if err := k.cdc.UnmarshalBinaryBare(value, &input); err != nil {
			return err
		}

		inputs = append(inputs, &input)
		amount = amount.Add(input.Amount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryAllInputResponse{
		Amount:     amount,
		Inputs:     inputs,
		Pagination: pageRes,
	}, nil
}
func (k Keeper) Input(c context.Context, req *types.QueryInputRequest) (*types.QueryInputResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	amount, err := sdk.ParseCoinsNormalized(req.Amounts)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid amount")
	}

	ctx := sdk.UnwrapSDKContext(c)

	inputs, total, err := k.SelectUtxos(ctx, addr, amount, req.Lock)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryInputResponse{
		Amount: total,
		Inputs: inputs,
	}, nil
}
