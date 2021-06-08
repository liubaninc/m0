package sdk

import (
	"context"

	wasmtypes "github.com/liubaninc/m0/x/wasm/types"

	"github.com/cosmos/cosmos-sdk/types/query"
)

func (c *Client) GetContract(name string) (*wasmtypes.QueryGetContractResponse, error) {
	queryClient := wasmtypes.NewQueryClient(c.ctx)
	res, err := queryClient.Contract(context.Background(), &wasmtypes.QueryGetContractRequest{
		Name: name,
	})
	return res, err
}

func (c *Client) GetContracts(key []byte, offset uint64, limit uint64, countTotal bool) (*wasmtypes.QueryAllContractResponse, error) {
	queryClient := wasmtypes.NewQueryClient(c.ctx)
	res, err := queryClient.ContractAll(context.Background(), &wasmtypes.QueryAllContractRequest{
		Pagination: &query.PageRequest{
			Key:        key,
			Offset:     offset,
			Limit:      limit,
			CountTotal: countTotal,
		},
	})
	return res, err
}

func (c *Client) GetAccountContracts(address string, key []byte, offset uint64, limit uint64, countTotal bool) (*wasmtypes.QueryGetAccountAllContractResponse, error) {
	queryClient := wasmtypes.NewQueryClient(c.ctx)
	res, err := queryClient.AccountContractAll(context.Background(), &wasmtypes.QueryGetAccountAllContractRequest{
		Name: address,
		Pagination: &query.PageRequest{
			Key:        key,
			Offset:     offset,
			Limit:      limit,
			CountTotal: countTotal,
		},
	})
	return res, err
}
