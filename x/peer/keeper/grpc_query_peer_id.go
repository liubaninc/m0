package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/liubaninc/m0/x/peer/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PeerIDAll(c context.Context, req *types.QueryAllPeerIDRequest) (*types.QueryAllPeerIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var peerIDs []*types.PeerID
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	peerIDStore := prefix.NewStore(store, types.KeyPrefix(types.PeerIDKey))

	pageRes, err := query.Paginate(peerIDStore, req.Pagination, func(key []byte, value []byte) error {
		var peerID types.PeerID
		if err := k.cdc.UnmarshalBinaryBare(value, &peerID); err != nil {
			return err
		}

		peerIDs = append(peerIDs, &peerID)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPeerIDResponse{PeerID: peerIDs, Pagination: pageRes}, nil
}

func (k Keeper) PeerID(c context.Context, req *types.QueryGetPeerIDRequest) (*types.QueryGetPeerIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPeerID(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetPeerIDResponse{PeerID: &val}, nil
}
