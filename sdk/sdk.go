package sdk

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/flags"
	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/liubaninc/m0/app"

	"github.com/cosmos/cosmos-sdk/client/tx"

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
	tc := c
	tc.locked = lock
	return tc
}

func (c Client) WithHeight(height int64) Client {
	tc := c
	tc.Context = tc.Context.WithHeight(height)
	return tc
}

func (c Client) WithKeyring(kr keyring.Keyring) Client {
	tc := c
	tc.Context = tc.Context.WithKeyring(kr)
	return tc
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

func (c Client) GenerateTx(from string, fees string, memo string, timeoutHeight uint64, msgs ...sdk.Msg) (client.TxBuilder, error) {
	var fromAddress sdk.AccAddress
	if addr, err := sdk.AccAddressFromBech32(from); err == nil {
		//info, err := c.Keyring.KeyByAddress(addr)
		//if err != nil {
		//	return nil, err
		//}
		fromAddress = addr
	} else {
		info, err := c.Keyring.Key(from)
		if err != nil {
			return nil, err
		}
		fromAddress = info.GetAddress()
	}

	clientCtx := c.Context.
		// WithFromName(info.GetName()).
		WithFromAddress(fromAddress)

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
	return txBuilder, nil
}

func (c Client) SignTx(from string, multiSigAddrStr string, txBuilder client.TxBuilder, overwriteSig bool) error {
	var info keyring.Info
	if addr, err := sdk.AccAddressFromBech32(from); err == nil {
		info, err = c.Keyring.KeyByAddress(addr)
		if err != nil {
			return err
		}
	} else {
		info, err = c.Keyring.Key(from)
		if err != nil {
			return err
		}
	}
	clientCtx := c.Context.
		WithFromName(info.GetName()).
		WithFromAddress(info.GetAddress())

	if len(multiSigAddrStr) > 0 {
		addr, err := sdk.AccAddressFromBech32(multiSigAddrStr)
		if err != nil {
			return fmt.Errorf("invalid multiSigAddrStr %v (%v)", multiSigAddrStr, err)
		}
		clientCtx = clientCtx.WithFromAddress(addr)
	}

	found := false
	signers := txBuilder.GetTx().GetSigners()
	for _, s := range signers {
		if bytes.Equal(clientCtx.FromAddress, s.Bytes()) {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("invalid signer: %s(%s)", clientCtx.FromName, clientCtx.FromAddress.String())
	}

	txf, err := tx.PrepareFactory(clientCtx, c.factory())
	if err != nil {
		return err
	}

	if len(multiSigAddrStr) > 0 {
		// Multisigs only support LEGACY_AMINO_JSON signing
		txf = txf.WithSignMode(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON)
	}

	err = tx.Sign(txf, clientCtx.GetFromName(), txBuilder, overwriteSig)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) MultiSignTx(txBuilder client.TxBuilder, multisigPubKey cryptotypes.PubKey, signatures ...signingtypes.SignatureV2) error {
	multisigPub := multisigPubKey.(*kmultisig.LegacyAminoPubKey)
	multisigSig := multisig.NewMultisig(len(multisigPub.PubKeys))

	sequence := uint64(0)
	for _, sig := range signatures {
		sequence = sig.Sequence
		if err := multisig.AddSignatureV2(multisigSig, sig, multisigPub.GetPubKeys()); err != nil {
			return err
		}
	}

	sigV2 := signingtypes.SignatureV2{
		PubKey:   multisigPub,
		Data:     multisigSig,
		Sequence: sequence,
	}

	err := txBuilder.SetSignatures(sigV2)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) BroadcastTx(tx sdk.Tx) (*sdk.TxResponse, error) {
	txBytes, err := c.TxConfig.TxEncoder()(tx)
	if err != nil {
		return nil, err
	}
	return c.Context.BroadcastTx(txBytes)
}

func (c Client) GenerateAndBroadcastTx(from string, fees string, memo string, timeoutHeight uint64, msgs ...sdk.Msg) (*sdk.TxResponse, error) {
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

	found := false
	signers := txBuilder.GetTx().GetSigners()
	for _, s := range signers {
		if bytes.Equal(clientCtx.FromAddress, s.Bytes()) {
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("invalid signer: %s(%s)", clientCtx.FromName, clientCtx.FromAddress.String())
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

func (c Client) factory() tx.Factory {
	signMode := signingtypes.SignMode_SIGN_MODE_UNSPECIFIED
	switch c.SignModeStr {
	case flags.SignModeDirect:
		signMode = signingtypes.SignMode_SIGN_MODE_DIRECT
	case flags.SignModeLegacyAminoJSON:
		signMode = signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
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
