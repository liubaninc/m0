package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/liubaninc/m0/x/wasm/types"
)

func createNProposeDeploy(keeper *Keeper, ctx sdk.Context, n int) []types.ProposeDeploy {
	items := make([]types.ProposeDeploy, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetProposeDeploy(ctx, items[i])
	}
	return items
}

func TestProposeDeployGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProposeDeploy(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProposeDeploy(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestProposeDeployRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProposeDeploy(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProposeDeploy(ctx, item.Index)
		_, found := keeper.GetProposeDeploy(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestProposeDeployGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNProposeDeploy(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllProposeDeploy(ctx))
}
