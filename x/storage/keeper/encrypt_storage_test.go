package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/liubaninc/m0/x/storage/types"
)

func createNEncryptStorage(keeper *Keeper, ctx sdk.Context, n int) []types.EncryptStorage {
	items := make([]types.EncryptStorage, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetEncryptStorage(ctx, items[i])
	}
	return items
}

func TestEncryptStorageGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNEncryptStorage(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEncryptStorage(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestEncryptStorageRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNEncryptStorage(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEncryptStorage(ctx, item.Index)
		_, found := keeper.GetEncryptStorage(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestEncryptStorageGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNEncryptStorage(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllEncryptStorage(ctx))
}
