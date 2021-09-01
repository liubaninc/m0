package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/liubaninc/m0/x/permission/types"
)

func createNAccount(keeper *Keeper, ctx sdk.Context, n int) []types.Account {
	items := make([]types.Account, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetAccount(ctx, items[i])
	}
	return items
}

func TestAccountGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccount(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestAccountRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccount(ctx, item.Index)
		_, found := keeper.GetAccount(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestAccountGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllAccount(ctx))
}
