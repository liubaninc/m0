package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/liubaninc/m0/x/pki/types"
)

func createNCertificates(keeper *Keeper, ctx sdk.Context, n int) []types.Certificates {
	items := make([]types.Certificates, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetCertificates(ctx, items[i])
	}
	return items
}

func TestCertificatesGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCertificates(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCertificates(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestCertificatesRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCertificates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCertificates(ctx, item.Index)
		_, found := keeper.GetCertificates(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestCertificatesGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCertificates(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllCertificates(ctx))
}
