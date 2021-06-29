package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/utxo/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Tampering(c context.Context, req *types.TamperingRequest) (*types.TamperingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	amount, err := sdk.ParseCoinNormalized(req.Amount)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid amount")
	}

	ctx := sdk.UnwrapSDKContext(c)
	cctx, w := ctx.CacheContext()

	ctx = cctx

	bytes := make([]byte, 16)
	rand.Read(bytes)
	hash := fmt.Sprintf("%X", tmhash.Sum(bytes))

	outs := make([]*types.Input, 1)
	input := &types.Input{
		RefTx:        hash,
		RefMsg:       0,
		RefOffset:    int32(0),
		FromAddr:     req.Address,
		Amount:       amount,
		FrozenHeight: 0,
	}
	outs[0] = input
	k.SetInput(ctx, types.Input{
		RefTx:        hash,
		RefMsg:       0,
		RefOffset:    int32(0),
		FromAddr:     req.Address,
		Amount:       amount,
		FrozenHeight: 0,
	})
	// bank
	if err := k.bankKeeper.AddCoins(ctx, addr, sdk.NewCoins(amount)); err != nil {
		return nil, status.Error(codes.Internal, "add amount")
	}

	acc := k.accountKeeper.GetAccount(ctx, addr)
	if acc == nil {
		k.accountKeeper.SetAccount(ctx, k.accountKeeper.NewAccountWithAddress(ctx, addr))
	}

	w()

	fmt.Println(k.GetInput(ctx, input.Index()))

	return &types.TamperingResponse{
		Outs: outs,
	}, nil
}
