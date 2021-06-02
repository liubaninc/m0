package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/liubaninc/m0/x/wasm/types"
)

func (k Keeper) PreExec(c context.Context, req *types.InvokeRPCRequest) (*types.InvokeRPCResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	res, err := PreExec(ctx, k, req)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrPreExec, "%s", err)
	}
	return res, nil
}
