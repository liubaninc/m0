package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/validator/types"
	"github.com/stretchr/testify/assert"
)

func createNValidato(keeper *Keeper, ctx sdk.Context, n int) []types.Validato {
	items := make([]types.Validato, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendValidato(ctx, items[i])
	}
	return items
}

func TestValidatoGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidato(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetValidato(ctx, item.Id))
	}
}

func TestValidatoExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidato(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasValidato(ctx, item.Id))
	}
}

func TestValidatoRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidato(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidato(ctx, item.Id)
		assert.False(t, keeper.HasValidato(ctx, item.Id))
	}
}

func TestValidatoGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidato(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllValidato(ctx))
}

func TestValidatoCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNValidato(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetValidatoCount(ctx))
}
