package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/liubaninc/m0/x/pki/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CertificatesAll(c context.Context, req *types.QueryAllCertificatesRequest) (*types.QueryAllCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var certificatess []*types.Certificates
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	certificatesStore := prefix.NewStore(store, types.KeyPrefix(types.CertificatesKey))

	pageRes, err := query.Paginate(certificatesStore, req.Pagination, func(key []byte, value []byte) error {
		var certificates types.Certificates
		if err := k.cdc.UnmarshalBinaryBare(value, &certificates); err != nil {
			return err
		}

		certificatess = append(certificatess, &certificates)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCertificatesResponse{Certificates: certificatess, Pagination: pageRes}, nil
}

func (k Keeper) Certificates(c context.Context, req *types.QueryGetCertificatesRequest) (*types.QueryGetCertificatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCertificates(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetCertificatesResponse{Certificates: &val}, nil
}
