package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/liubaninc/m0/x/wasm/types"
)

func createNApproveDeploy(keeper *Keeper, ctx sdk.Context, n int) []types.ApproveDeploy {
	items := make([]types.ApproveDeploy, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetApproveDeploy(ctx, items[i])
	}
	return items
}

func TestApproveDeployGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNApproveDeploy(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetApproveDeploy(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestApproveDeployRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNApproveDeploy(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveApproveDeploy(ctx, item.Index)
		_, found := keeper.GetApproveDeploy(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestApproveDeployGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNApproveDeploy(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllApproveDeploy(ctx))
}
