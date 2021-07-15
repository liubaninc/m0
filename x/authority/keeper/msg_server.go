package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/authority/types"
	"math"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func AccountApprovalsCount(ctx sdk.Context, keeper Keeper) int {
	count := keeper.CountAccountsWithRole(ctx, types.AUTHORITY)
	fmt.Printf("count account with role : %n/n", count)
	return int(math.Round(types.AccountApprovalPercent * float64(count)))
}

func HasApprovalFrom(address string, approvals []string) bool {
	for _, approval := range approvals {
		if approval == address {
			return true
		}
	}
	return false
}
