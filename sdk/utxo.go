package sdk

import (
	"context"

	utxotypes "github.com/liubaninc/m0/x/utxo/types"

	"github.com/cosmos/cosmos-sdk/types/query"
)

func (c *Client) GetToken(denom string) (*utxotypes.QueryGetTokenResponse, error) {
	queryClient := utxotypes.NewQueryClient(c.ctx)
	res, err := queryClient.Token(context.Background(), &utxotypes.QueryGetTokenRequest{
		Name: denom,
	})
	return res, err
}

func (c *Client) GetTokens(key []byte, offset uint64, limit uint64, countTotal bool) (*utxotypes.QueryAllTokenResponse, error) {
	queryClient := utxotypes.NewQueryClient(c.ctx)
	res, err := queryClient.TokenAll(context.Background(), &utxotypes.QueryAllTokenRequest{
		Pagination: &query.PageRequest{
			Key:        key,
			Offset:     offset,
			Limit:      limit,
			CountTotal: countTotal,
		},
	})
	return res, err
}

func (c *Client) GetInput(address, amount string, lock int64) (*utxotypes.QueryInputResponse, error) {
	queryClient := utxotypes.NewQueryClient(c.ctx)
	res, err := queryClient.Input(context.Background(), &utxotypes.QueryInputRequest{
		Address: address,
		Amount:  amount,
		Lock:    lock,
	})
	return res, err
}

func (c *Client) GetInputs(address, denom string, key []byte, offset uint64, limit uint64, countTotal bool) (*utxotypes.QueryAllInputResponse, error) {
	queryClient := utxotypes.NewQueryClient(c.ctx)
	res, err := queryClient.InputAll(context.Background(), &utxotypes.QueryAllInputRequest{
		Address: address,
		Denom:   denom,
		Pagination: &query.PageRequest{
			Key:        key,
			Offset:     offset,
			Limit:      limit,
			CountTotal: countTotal,
		},
	})
	return res, err
}
