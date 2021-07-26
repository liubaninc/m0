package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/liubaninc/m0/x/peer/types"
)

func createNPeerID(keeper *Keeper, ctx sdk.Context, n int) []types.PeerID {
	items := make([]types.PeerID, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetPeerID(ctx, items[i])
	}
	return items
}

func TestPeerIDGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPeerID(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPeerID(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestPeerIDRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPeerID(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePeerID(ctx, item.Index)
		_, found := keeper.GetPeerID(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestPeerIDGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPeerID(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllPeerID(ctx))
}
