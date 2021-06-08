package sdk

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (c *Client) GetAccount(address string) (*authtypes.QueryAccountResponse, error) {
	queryClient := authtypes.NewQueryClient(c.ctx)
	res, err := queryClient.Account(context.Background(), &authtypes.QueryAccountRequest{
		Address: address,
	})
	return res, err
}

func (c *Client) GetAccounts(key []byte, offset uint64, limit uint64, countTotal bool) (*authtypes.QueryAccountsResponse, error) {
	queryClient := authtypes.NewQueryClient(c.ctx)
	res, err := queryClient.Accounts(context.Background(), &authtypes.QueryAccountsRequest{
		Pagination: &query.PageRequest{
			Key:        key,
			Offset:     offset,
			Limit:      limit,
			CountTotal: countTotal,
		},
	})
	return res, err
}

func (c *Client) GetAccountBalances(address string, key []byte, offset uint64, limit uint64, countTotal bool) (*banktypes.QueryAllBalancesResponse, error) {
	queryClient := banktypes.NewQueryClient(c.ctx)
	res, err := queryClient.AllBalances(context.Background(), &banktypes.QueryAllBalancesRequest{
		Address: address,
		Pagination: &query.PageRequest{
			Key:        key,
			Offset:     offset,
			Limit:      limit,
			CountTotal: countTotal,
		},
	})
	return res, err
}

func (c *Client) GetAccountBalance(address string, denom string) (*banktypes.QueryBalanceResponse, error) {
	queryClient := banktypes.NewQueryClient(c.ctx)
	res, err := queryClient.Balance(context.Background(), &banktypes.QueryBalanceRequest{
		Address: address,
		Denom:   denom,
	})
	return res, err
}

func (c *Client) GetTotalSupply() (*banktypes.QueryTotalSupplyResponse, error) {
	queryClient := banktypes.NewQueryClient(c.ctx)
	res, err := queryClient.TotalSupply(context.Background(), &banktypes.QueryTotalSupplyRequest{})
	return res, err
}

func (c *Client) GetSupply(denom string) (*banktypes.QuerySupplyOfResponse, error) {
	queryClient := banktypes.NewQueryClient(c.ctx)
	res, err := queryClient.SupplyOf(context.Background(), &banktypes.QuerySupplyOfRequest{
		Denom: denom,
	})
	return res, err
}
