package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/liubaninc/m0/x/pki/types"
)

func TestCertificateQuerySingle(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCertificate(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCertificateRequest
		response *types.QueryGetCertificateResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCertificateRequest{Index: msgs[0].Index},
			response: &types.QueryGetCertificateResponse{Certificate: &msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCertificateRequest{Index: msgs[1].Index},
			response: &types.QueryGetCertificateResponse{Certificate: &msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetCertificateRequest{Index: "missing"},
			err:     status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Certificate(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestCertificateQueryPaginated(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCertificate(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCertificateRequest {
		return &types.QueryAllCertificateRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CertificateAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, &msgs[j], resp.Certificate[j-i])
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CertificateAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, &msgs[j], resp.Certificate[j-i])
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CertificateAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CertificateAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
