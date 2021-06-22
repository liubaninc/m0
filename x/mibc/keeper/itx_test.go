package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/mibc/types"
	"github.com/stretchr/testify/assert"
)

func createNItx(keeper *Keeper, ctx sdk.Context, n int) []types.Itx {
	items := make([]types.Itx, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendItx(ctx, items[i])
	}
	return items
}

func TestItxGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNItx(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetItx(ctx, item.Id))
	}
}

func TestItxExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNItx(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasItx(ctx, item.Id))
	}
}

func TestItxRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNItx(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveItx(ctx, item.Id)
		assert.False(t, keeper.HasItx(ctx, item.Id))
	}
}

func TestItxGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNItx(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllItx(ctx))
}

func TestItxCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNItx(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetItxCount(ctx))
}
