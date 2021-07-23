package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/liubaninc/m0/x/pki/types"
)

func createNCertificate(keeper *Keeper, ctx sdk.Context, n int) []types.Certificate {
	items := make([]types.Certificate, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetCertificate(ctx, items[i])
	}
	return items
}

func TestCertificateGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCertificate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCertificate(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestCertificateRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCertificate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCertificate(ctx, item.Index)
		_, found := keeper.GetCertificate(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestCertificateGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNCertificate(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllCertificate(ctx))
}
