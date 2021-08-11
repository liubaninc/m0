package mchain

import (
	"errors"

	sdk "github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/go/pb"
)

var (
	errUnimplemented = errors.New("unimplemented")
)

type chainCore struct {
}

// GetAccountAddress get addresses associated with account name
func (c *chainCore) GetAccountAddresses(accountName string) ([]string, error) {
	return []string{}, nil
}

// VerifyContractPermission verify permission of calling contract
func (c *chainCore) VerifyContractPermission(initiator string, authRequire []string, contractName, methodName string) (bool, error) {
	return true, nil
}

// VerifyContractOwnerPermission verify contract ownership permisson
func (c *chainCore) VerifyContractOwnerPermission(contractName string, authRequire []string) error {
	return nil
}

// QueryTransaction query confirmed tx
func (c *chainCore) QueryTransaction(txid []byte) (*sdk.Transaction, error) {
	return new(sdk.Transaction), nil
}

// QueryBlock query block
func (c *chainCore) QueryBlock(height int64) (*sdk.Block, error) {
	return new(sdk.Block), nil
}
