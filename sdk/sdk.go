package sdk

import (
	"github.com/liubaninc/m0/app"

	"github.com/cosmos/cosmos-sdk/client"
)

type Client struct {
	ctx client.Context
}

func New(rpcURI string, block bool) (*Client, error) {
	node, err := client.NewClientFromNode(rpcURI)
	if err != nil {
		return nil, err
	}

	mode := "sync"
	if block {
		mode = "block"
	}

	encodingConfig := app.MakeEncodingConfig()
	clientCtx := client.Context{}.WithNodeURI(rpcURI).
		WithClient(node).
		WithBroadcastMode(mode).
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithTxConfig(encodingConfig.TxConfig)
	return &Client{
		ctx: clientCtx,
	}, nil
}

func MustNew(rpcURI string, block bool) *Client {
	client, err := New(rpcURI, block)
	if err != nil {
		panic(err)
	}
	return client
}

func init() {
	app.SetConfig()
}
