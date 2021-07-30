package peer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/peer/keeper"
	"github.com/liubaninc/m0/x/peer/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	// this line is used by starport scaffolding # genesis/module/init
	// Set all the peerID
	for _, elem := range genState.PeerIDList {
		k.SetPeerID(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Params = k.GetParams(ctx)
	// this line is used by starport scaffolding # genesis/module/export
	// Get all peerID
	peerIDList := k.GetAllPeerID(ctx)
	for _, elem := range peerIDList {
		elem := elem
		genesis.PeerIDList = append(genesis.PeerIDList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
