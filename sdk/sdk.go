package sdk

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"os"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/liubaninc/m0/app"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/client"
)

type Client struct {
	client.Context
	locked int64 // utxo locked time if needed
}

func (c Client) WithLock(lock int64) Client {
	c.locked = lock
	return c
}

func MustNew(rpcURI string, kr keyring.Keyring) Client {
	client, err := New(rpcURI, kr)
	if err != nil {
		panic(err)
	}
	return client
}

func New(rpcURI string, kr keyring.Keyring) (Client, error) {
	c := Client{
		Context: client.Context{},
		locked:  60,
	}

	node, err := client.NewClientFromNode(rpcURI)
	if err != nil {
		return c, err
	}

	genesis, err := node.Genesis(context.Background())
	if err != nil {
		return c, err
	}

	encodingConfig := app.MakeEncodingConfig()
	c.Context = c.Context.
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithOutput(os.Stdout).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithNodeURI(rpcURI).
		WithClient(node).
		WithChainID(genesis.Genesis.ChainID).
		WithSignModeStr("direct").
		WithBroadcastMode("block").
		WithOutputFormat("json").
		WithKeyring(kr)
	return c, nil
}


func (c *Client) GenerateAndBroadcastTx(from string, fees string, memo string, timeoutHeight uint64, msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	var info keyring.Info
	if addr, err := sdk.AccAddressFromBech32(from); err == nil {
		info, err = c.Keyring.KeyByAddress(addr)
		if err != nil {
			return nil, err
		}
	} else {
		info, err = c.Keyring.Key(from)
		if err != nil {
			return nil, err
		}
	}

	clientCtx := c.Context.
		WithFromName(info.GetName()).
		WithFromAddress(info.GetAddress())

	txf, err := tx.PrepareFactory(clientCtx, c.factory())
	if err != nil {
		return nil, err
	}
	txf = txf.WithFees(fees).
		WithMemo(memo).
		WithTimeoutHeight(timeoutHeight)

	_, adjusted, err := tx.CalculateGas(clientCtx.QueryWithData, txf, msgs...)
	if err != nil {
		return nil, err
	}

	txf = txf.WithGas(adjusted)

	txBuilder, err := tx.BuildUnsignedTx(txf, msgs...)
	if err != nil {
		return nil, err
	}

	err = tx.Sign(txf, clientCtx.GetFromName(), txBuilder, true)
	if err != nil {
		return nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return clientCtx.BroadcastTx(txBytes)
}


func (c *Client) factory() tx.Factory {
	signMode := signing.SignMode_SIGN_MODE_UNSPECIFIED
	switch c.SignModeStr {
	case flags.SignModeDirect:
		signMode = signing.SignMode_SIGN_MODE_DIRECT
	case flags.SignModeLegacyAminoJSON:
		signMode = signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	}
	txf := tx.Factory{}.
		WithKeybase(c.Keyring).
		WithTxConfig(c.TxConfig).
		WithAccountRetriever(c.AccountRetriever).
		WithSignMode(signMode).
		WithChainID(c.ChainID).
		WithSimulateAndExecute(true).
		WithGasAdjustment(1.5)
	return txf
}


func init() {
	app.SetConfig()
}