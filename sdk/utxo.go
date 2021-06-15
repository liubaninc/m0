package sdk

import (
	"context"
	"fmt"

	utxotypes "github.com/liubaninc/m0/x/utxo/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (c *Client) GetToken(denom string) (*utxotypes.QueryGetTokenResponse, error) {
	if err := sdk.ValidateDenom(denom); err != nil {
		return nil, fmt.Errorf("invalid token %s (%s)", denom, err)
	}

	queryClient := utxotypes.NewQueryClient(c)
	res, err := queryClient.Token(context.Background(), &utxotypes.QueryGetTokenRequest{
		Name: denom,
	})
	return res, err
}

func (c *Client) GetTokens(key []byte, offset uint64, limit uint64, countTotal bool) (*utxotypes.QueryAllTokenResponse, error) {
	queryClient := utxotypes.NewQueryClient(c)
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

func (c *Client) GetInput(address, amounts string, lock int64) (*utxotypes.QueryInputResponse, error) {
	if _, err := sdk.AccAddressFromBech32(address); err != nil {
		return nil, fmt.Errorf("invalid address %s (%s)", address, err)
	}
	if _, err := sdk.ParseCoinsNormalized(amounts); err != nil {
		return nil, fmt.Errorf("invalid amount %s (%s)", amounts, err)
	}

	queryClient := utxotypes.NewQueryClient(c)
	res, err := queryClient.Input(context.Background(), &utxotypes.QueryInputRequest{
		Address: address,
		Amounts: amounts,
		Lock:    lock,
	})
	return res, err
}

func (c *Client) GetInputs(address, denom string, key []byte, offset uint64, limit uint64, countTotal bool) (*utxotypes.QueryAllInputResponse, error) {
	if _, err := sdk.AccAddressFromBech32(address); err != nil {
		return nil, fmt.Errorf("invalid address %s (%s)", address, err)
	}
	if err := sdk.ValidateDenom(denom); err != nil {
		return nil, fmt.Errorf("invalid denom %s (%s)", denom, err)
	}

	queryClient := utxotypes.NewQueryClient(c)
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

func (c *Client) IssueMsg(from string, tos []string, amounts []string, desc string, fees string) (sdk.Msg, error) {
	if _, err := sdk.AccAddressFromBech32(from); err != nil {
		return nil, fmt.Errorf("invalid from %s (%s)", from, err)
	}
	if len(tos) != len(amounts) {
		return nil, fmt.Errorf("mismatch len tos %d (amounts %d)", len(tos), len(amounts))
	}
	feeCoins, err := sdk.ParseCoinsNormalized(fees)
	if err != nil {
		return nil, fmt.Errorf("invalid fees %s (%s)", fees, err)
	}

	neededTotal := sdk.NewCoins()
	var outputs []*utxotypes.Output
	for index, to := range tos {
		if _, err := sdk.AccAddressFromBech32(to); err != nil {
			return nil, fmt.Errorf("invalid to %s (%s)", to, err)
		}
		amountCoins, err := sdk.ParseCoinsNormalized(amounts[index])
		if err != nil {
			return nil, fmt.Errorf("invalid amounts %s (%s)", amounts[index], err)
		}
		for _, amount := range amountCoins {
			outputs = append(outputs, &utxotypes.Output{
				ToAddr: to,
				Amount: amount,
			})
		}
	}

	for _, fee := range feeCoins {
		outputs = append(outputs, &utxotypes.Output{
			ToAddr: authtypes.NewModuleAddress(authtypes.FeeCollectorName).String(),
			Amount: fee,
		})
		neededTotal.Add(fee)
	}

	var inputs []*utxotypes.Input
	res, err := c.GetInput(from, neededTotal.String(), c.locked)
	if err != nil {
		return nil, fmt.Errorf("get input %s (%s)", from, err)
	}
	inputs = append(inputs, res.Inputs...)

	changeCoins := res.Amount.Sub(neededTotal)
	for _, coin := range changeCoins {
		outputs = append(outputs, &utxotypes.Output{
			ToAddr: from,
			Amount: coin,
		})
	}

	return utxotypes.NewMsgIssue(from, inputs, outputs, desc), nil
}

func (c *Client) DestroyMsg(from string, amounts string, desc string, fees string) (sdk.Msg, error) {
	if _, err := sdk.AccAddressFromBech32(from); err != nil {
		return nil, fmt.Errorf("invalid from %s (%s)", from, err)
	}
	amountCoins, err := sdk.ParseCoinsNormalized(amounts)
	if err != nil {
		return nil, fmt.Errorf("invalid amounts %s (%s)", amounts, err)
	}

	feeCoins, err := sdk.ParseCoinsNormalized(fees)
	if err != nil {
		return nil, fmt.Errorf("invalid fees %s (%s)", fees, err)
	}

	neededTotal := sdk.NewCoins()
	var outputs []*utxotypes.Output
	neededTotal = neededTotal.Add(amountCoins...)
	for _, fee := range feeCoins {
		outputs = append(outputs, &utxotypes.Output{
			ToAddr: authtypes.NewModuleAddress(authtypes.FeeCollectorName).String(),
			Amount: fee,
		})
		neededTotal.Add(fee)
	}

	var inputs []*utxotypes.Input
	res, err := c.GetInput(from, neededTotal.String(), c.locked)
	if err != nil {
		return nil, fmt.Errorf("get input %s (%s)", from, err)
	}
	inputs = append(inputs, res.Inputs...)

	changeCoins := res.Amount.Sub(neededTotal)
	for _, coin := range changeCoins {
		outputs = append(outputs, &utxotypes.Output{
			ToAddr: from,
			Amount: coin,
		})
	}

	return utxotypes.NewMsgDestroy(from, inputs, outputs, desc), nil
}

func (c *Client) SendMsg(from string, tos []string, amounts []string, desc string, fees string) (sdk.Msg, error) {
	if _, err := sdk.AccAddressFromBech32(from); err != nil {
		return nil, fmt.Errorf("invalid from %s (%s)", from, err)
	}
	if len(tos) != len(amounts) {
		return nil, fmt.Errorf("mismatch len tos %d (amounts %d)", len(tos), len(amounts))
	}
	feeCoins, err := sdk.ParseCoinsNormalized(fees)
	if err != nil {
		return nil, fmt.Errorf("invalid fees %s (%s)", fees, err)
	}

	neededTotal := sdk.NewCoins()
	var outputs []*utxotypes.Output
	for index, to := range tos {
		if _, err := sdk.AccAddressFromBech32(to); err != nil {
			return nil, fmt.Errorf("invalid to %s (%s)", to, err)
		}
		amountCoins, err := sdk.ParseCoinsNormalized(amounts[index])
		if err != nil {
			return nil, fmt.Errorf("invalid amounts %s (%s)", amounts[index], err)
		}
		for _, amount := range amountCoins {
			outputs = append(outputs, &utxotypes.Output{
				ToAddr: to,
				Amount: amount,
			})
			neededTotal.Add(amount)
		}
	}

	for _, fee := range feeCoins {
		outputs = append(outputs, &utxotypes.Output{
			ToAddr: authtypes.NewModuleAddress(authtypes.FeeCollectorName).String(),
			Amount: fee,
		})
		neededTotal.Add(fee)
	}

	var inputs []*utxotypes.Input
	res, err := c.GetInput(from, neededTotal.String(), c.locked)
	if err != nil {
		return nil, fmt.Errorf("get input %s (%s)", from, err)
	}
	inputs = append(inputs, res.Inputs...)

	changeCoins := res.Amount.Sub(neededTotal)
	for _, coin := range changeCoins {
		outputs = append(outputs, &utxotypes.Output{
			ToAddr: from,
			Amount: coin,
		})
	}
	return utxotypes.NewMsgIssue(from, inputs, outputs, desc), nil
}

func (c *Client) BroadcastIssueTx(from string, tos []string, amounts []string, desc string, fees string, memo string) (*sdk.TxResponse, error) {
	msg, err := c.IssueMsg(from, tos, amounts, desc, fees)
	if err != nil {
		return nil, err
	}
	return c.GenerateAndBroadcastTx(from, fees, memo, 0, msg)
}
func (c *Client) BroadcastDestroyTx(from string, amounts string, desc string, fees string, memo string) (*sdk.TxResponse, error) {
	msg, err := c.DestroyMsg(from, amounts, desc, fees)
	if err != nil {
		return nil, err
	}
	return c.GenerateAndBroadcastTx(from, fees, memo, 0, msg)
}
func (c *Client) BroadcastSendTx(from string, tos []string, amounts []string, desc string, fees string, memo string) (*sdk.TxResponse, error) {
	msg, err := c.SendMsg(from, tos, amounts, desc, fees)
	if err != nil {
		return nil, err
	}
	return c.GenerateAndBroadcastTx(from, fees, memo, 0, msg)
}
