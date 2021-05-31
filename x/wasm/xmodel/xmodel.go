package xmodel

import (
	sdkpb "github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/go/pb"
)

// Iterator iterator interface
type Iterator interface {
	Data() *VersionedData
	First() bool
	Next() bool
	Error() error
	Key() []byte
	Release()
}

// XMReader xmodel interface for reader
type XMReader interface {
	//读取一个key的值，返回的value就是有版本的data
	Get(bucket string, key []byte) (*VersionedData, error)
	//扫描一个bucket中所有的kv, 调用者可以设置key区间[startKey, endKey)
	Select(bucket string, startKey []byte, endKey []byte) (Iterator, error)
}

// XModel xmodel interface for reader
type XModel interface {
	// Put put a pair of <key, value> into xModel
	Put(bucket string, key []byte, value []byte) error
	// Del delete one key from outPutCache, marked its value as `DelFlag`
	Del(bucket string, key []byte) error
	// Get get value for specific key, return value with version
	Get(bucket string, key []byte) (*VersionedData, error)
	// Select select all kv from a bucket, can set key range, left closed, right opend
	Select(bucket string, startKey []byte, endKey []byte) (Iterator, error)

	// AddEvent add contract event into xmodel
	Transfer(from, to string, amount string) error
	// Transfer transfer tokens using utxo into xmodel
	AddEvent(events ...*ContractEvent)

	// GetRWSets get read/write sets
	GetRWSets() ([]*VersionedData, []*PureData, error)
}

// ChainCore is the interface of chain service
type ChainCore interface {
	// QueryTransaction query confirmed tx
	QueryTransaction(txid []byte) (*sdkpb.Transaction, error)
	// QueryBlock query block
	QueryBlock(height int64) (*sdkpb.Block, error)
}
