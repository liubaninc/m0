package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/liubaninc/m0/x/validator/types"
)

func createNLastValidatorPower(keeper *Keeper, ctx sdk.Context, n int) []types.LastValidatorPower {
	items := make([]types.LastValidatorPower, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetLastValidatorPower(ctx, items[i])
	}
	return items
}

func TestLastValidatorPowerGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNLastValidatorPower(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLastValidatorPower(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestLastValidatorPowerRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNLastValidatorPower(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLastValidatorPower(ctx, item.Index)
		_, found := keeper.GetLastValidatorPower(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestLastValidatorPowerGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNLastValidatorPower(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllLastValidatorPower(ctx))
}
