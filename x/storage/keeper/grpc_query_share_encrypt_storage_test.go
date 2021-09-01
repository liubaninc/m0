package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/liubaninc/m0/x/storage/types"
)

func TestShareEncryptStorageQuerySingle(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNShareEncryptStorage(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetShareEncryptStorageRequest
		response *types.QueryGetShareEncryptStorageResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetShareEncryptStorageRequest{Index: msgs[0].Index},
			response: &types.QueryGetShareEncryptStorageResponse{ShareEncryptStorage: &msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetShareEncryptStorageRequest{Index: msgs[1].Index},
			response: &types.QueryGetShareEncryptStorageResponse{ShareEncryptStorage: &msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetShareEncryptStorageRequest{Index: "missing"},
			err:     status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ShareEncryptStorage(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestShareEncryptStorageQueryPaginated(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNShareEncryptStorage(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllShareEncryptStorageRequest {
		return &types.QueryAllShareEncryptStorageRequest{
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
			resp, err := keeper.ShareEncryptStorageAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, &msgs[j], resp.ShareEncryptStorage[j-i])
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ShareEncryptStorageAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, &msgs[j], resp.ShareEncryptStorage[j-i])
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ShareEncryptStorageAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ShareEncryptStorageAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
