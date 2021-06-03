package keeper

import (
	"context"
	"encoding/hex"

	sdkpb "github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/go/pb"
	"github.com/spf13/viper"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

type contractChainCore struct {
	rpcclient.Client
}

func newContractChainCore() *contractChainCore {
	nodeURI := viper.GetString("rpc.laddr")
	rpc, err := rpchttp.New(nodeURI, "/websocket")
	if err != nil {
		panic(err)
	}
	return &contractChainCore{
		Client: rpc,
	}
}

// QueryTransaction query confirmed tx
func (cc *contractChainCore) QueryTransaction(txid []byte) (*sdkpb.Transaction, error) {
	tx, err := cc.Client.Tx(context.Background(), txid, false)
	if err != nil {
		return nil, err
	}
	blk, err := cc.Client.Block(context.Background(), &tx.Height)
	if err != nil {
		return nil, err
	}
	return &sdkpb.Transaction{
		Txid:    tx.Hash.String(),
		Blockid: blk.BlockID.Hash.String(),
	}, nil
}

// QueryBlock query block
func (cc *contractChainCore) QueryBlock(height int64) (*sdkpb.Block, error) {
	blk, err := cc.Client.Block(context.Background(), &height)
	if err != nil {
		return nil, err
	}
	count := len(blk.Block.Txs)
	hashes := make([]string, count)
	for index, tx := range blk.Block.Txs {
		hashes[index] = hex.EncodeToString(tx.Hash())
	}
	return &sdkpb.Block{
		Blockid:  blk.BlockID.Hash.String(),
		PreHash:  blk.Block.LastBlockID.Hash.String(),
		Proposer: blk.Block.ProposerAddress,
		Height:   blk.Block.Height,
		Txids:    hashes,
		TxCount:  int32(count),
	}, nil
}
