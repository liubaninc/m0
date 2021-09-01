package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/liubaninc/m0/x/storage/types"
)

func createNRecryptAccount(keeper *Keeper, ctx sdk.Context, n int) []types.RecryptAccount {
	items := make([]types.RecryptAccount, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetRecryptAccount(ctx, items[i])
	}
	return items
}

func TestRecryptAccountGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecryptAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRecryptAccount(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestRecryptAccountRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecryptAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRecryptAccount(ctx, item.Index)
		_, found := keeper.GetRecryptAccount(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestRecryptAccountGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRecryptAccount(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllRecryptAccount(ctx))
}
