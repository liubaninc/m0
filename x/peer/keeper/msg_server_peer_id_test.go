package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/liubaninc/m0/x/peer/types"
)

func TestPeerIDMsgServerCreate(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	srv := NewMsgServerImpl(*keeper)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		idx := fmt.Sprintf("%d", i)
		expected := &types.MsgCreatePeerID{Creator: creator, Index: idx}
		_, err := srv.CreatePeerID(wctx, expected)
		require.NoError(t, err)
		rst, found := keeper.GetPeerID(ctx, expected.Index)
		require.True(t, found)
		assert.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestPeerIDMsgServerUpdate(t *testing.T) {
	creator := "A"
	index := "any"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdatePeerID
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdatePeerID{Creator: creator, Index: index},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdatePeerID{Creator: "B", Index: index},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgUpdatePeerID{Creator: creator, Index: "missing"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			keeper, ctx := setupKeeper(t)
			srv := NewMsgServerImpl(*keeper)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreatePeerID{Creator: creator, Index: index}
			_, err := srv.CreatePeerID(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdatePeerID(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := keeper.GetPeerID(ctx, expected.Index)
				require.True(t, found)
				assert.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestPeerIDMsgServerDelete(t *testing.T) {
	creator := "A"
	index := "any"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeletePeerID
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeletePeerID{Creator: creator, Index: index},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeletePeerID{Creator: "B", Index: index},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeletePeerID{Creator: creator, Index: "missing"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			keeper, ctx := setupKeeper(t)
			srv := NewMsgServerImpl(*keeper)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreatePeerID(wctx, &types.MsgCreatePeerID{Creator: creator, Index: index})
			require.NoError(t, err)
			_, err = srv.DeletePeerID(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := keeper.GetPeerID(ctx, tc.request.Index)
				require.False(t, found)
			}
		})
	}
}
