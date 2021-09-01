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

func (k Keeper) CertificateAll(c context.Context, req *types.QueryAllCertificateRequest) (*types.QueryAllCertificateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var certificates []*types.Certificate
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	certificateStore := prefix.NewStore(store, types.KeyPrefix(types.CertificateKey))

	pageRes, err := query.Paginate(certificateStore, req.Pagination, func(key []byte, value []byte) error {
		var certificate types.Certificate
		if err := k.cdc.UnmarshalBinaryBare(value, &certificate); err != nil {
			return err
		}

		certificates = append(certificates, &certificate)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCertificateResponse{Certificate: certificates, Pagination: pageRes}, nil
}

func (k Keeper) Certificate(c context.Context, req *types.QueryGetCertificateRequest) (*types.QueryGetCertificateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCertificate(ctx, req.Issuer, req.SerialNumber)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetCertificateResponse{Certificate: &val}, nil
}
