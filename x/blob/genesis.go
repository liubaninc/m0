package blob

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/liubaninc/m0/x/blob/keeper"
	"github.com/liubaninc/m0/x/blob/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the comment
	for _, elem := range genState.CommentList {
		k.SetComment(ctx, *elem)
	}

	// Set comment count
	k.SetCommentCount(ctx, genState.CommentCount)

	// Set all the blob
	for _, elem := range genState.BlobList {
		k.SetBlob(ctx, *elem)
	}

	// Set blob count
	k.SetBlobCount(ctx, genState.BlobCount)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all comment
	commentList := k.GetAllComment(ctx)
	for _, elem := range commentList {
		elem := elem
		genesis.CommentList = append(genesis.CommentList, &elem)
	}

	// Set the current count
	genesis.CommentCount = k.GetCommentCount(ctx)

	// Get all blob
	blobList := k.GetAllBlob(ctx)
	for _, elem := range blobList {
		elem := elem
		genesis.BlobList = append(genesis.BlobList, &elem)
	}

	// Set the current count
	genesis.BlobCount = k.GetBlobCount(ctx)

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
