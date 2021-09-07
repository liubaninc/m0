package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/liubaninc/m0/x/storage/types"
)

func createNShareEncryptStorage(keeper *Keeper, ctx sdk.Context, n int) []types.ShareEncryptStorage {
	items := make([]types.ShareEncryptStorage, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetShareEncryptStorage(ctx, items[i])
	}
	return items
}

func TestShareEncryptStorageGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNShareEncryptStorage(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetShareEncryptStorage(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestShareEncryptStorageRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNShareEncryptStorage(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveShareEncryptStorage(ctx, item.Index)
		_, found := keeper.GetShareEncryptStorage(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestShareEncryptStorageGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNShareEncryptStorage(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllShareEncryptStorage(ctx))
}
