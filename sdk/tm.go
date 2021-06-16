package sdk

import (
	"context"
	"encoding/hex"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

func (c Client) GetABCIInfo() (*coretypes.ResultABCIInfo, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.ABCIInfo(context.Background())
}

func (c Client) GetBlock(height int64) (*coretypes.ResultBlock, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.Block(context.Background(), &height)
}

func (c Client) GetBlockLatest() (*coretypes.ResultBlock, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.Block(context.Background(), nil)
}

func (c Client) GetBlockResult(height int64) (*coretypes.ResultBlockResults, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.BlockResults(context.Background(), &height)
}

func (c Client) GetGenesis() (*coretypes.ResultGenesis, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.Genesis(context.Background())
}

func (c Client) GetHealth() (*coretypes.ResultHealth, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.Health(context.Background())
}

func (c Client) GetNetInfo() (*coretypes.ResultNetInfo, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.NetInfo(context.Background())
}

func (c Client) GetStatus() (*coretypes.ResultStatus, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.Status(context.Background())
}

func (c Client) GetTx(hash string) (*coretypes.ResultTx, error) {
	rpcClient, _ := c.GetNode()
	txHash, err := hex.DecodeString(hash)
	if err != nil {
		return nil, err
	}
	return rpcClient.Tx(context.Background(), txHash, false)
}

func (c Client) GetTxSearch(query string, page int, per_page int) (*coretypes.ResultTxSearch, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.TxSearch(context.Background(), query, false, &page, &per_page, "")
}

func (c Client) GetUnconfirmedTxsNum() (*coretypes.ResultUnconfirmedTxs, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.NumUnconfirmedTxs(context.Background())
}

func (c Client) GetUnconfirmedTxs(limit int) (*coretypes.ResultUnconfirmedTxs, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.UnconfirmedTxs(context.Background(), &limit)
}

func (c Client) GetValidators(height int64, page int, per_page int) (*coretypes.ResultValidators, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.Validators(context.Background(), &height, &page, &per_page)
}

func (c Client) GetValidatorsLatest(page int, per_page int) (*coretypes.ResultValidators, error) {
	rpcClient, _ := c.GetNode()
	return rpcClient.Validators(context.Background(), nil, &page, &per_page)
}
