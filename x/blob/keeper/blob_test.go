package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/blob/types"
	"github.com/stretchr/testify/assert"
)

func createNBlob(keeper *Keeper, ctx sdk.Context, n int) []types.Blob {
	items := make([]types.Blob, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendBlob(ctx, items[i])
	}
	return items
}

func TestBlobGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBlob(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetBlob(ctx, item.Id))
	}
}

func TestBlobExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBlob(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasBlob(ctx, item.Id))
	}
}

func TestBlobRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBlob(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBlob(ctx, item.Id)
		assert.False(t, keeper.HasBlob(ctx, item.Id))
	}
}

func TestBlobGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBlob(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllBlob(ctx))
}

func TestBlobCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBlob(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetBlobCount(ctx))
}
