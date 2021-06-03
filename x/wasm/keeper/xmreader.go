package keeper

import (
	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/liubaninc/m0/x/wasm/xmodel"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ xmodel.Iterator = &Iterator{}
	_ xmodel.XMReader = &XMReader{}
)

func NewIterator(ctx sdk.Context, k Keeper, bucket string, startKey []byte, endKey []byte) *Iterator {
	iter := &Iterator{
		k:        k,
		ctx:      ctx,
		bucket:   bucket,
		startKey: startKey,
		endKey:   endKey,
	}
	iter.First()
	return iter
}

func NewXMReader(ctx sdk.Context, k Keeper) *XMReader {
	return &XMReader{
		ctx: ctx,
		k:   k,
	}
}

// Iterator data structure for XModel Iterator
type Iterator struct {
	k   Keeper
	ctx sdk.Context

	bucket   string
	startKey []byte
	endKey   []byte
	iter     sdk.Iterator
	err      error
	first    bool
}

// Data get data pointer to VersionedData for XMIterator
func (di *Iterator) Data() *xmodel.VersionedData {
	version := di.iter.Value()
	verData, err := di.k.GetVersionedData(di.ctx, version)
	di.err = err
	return verData
}

func (di *Iterator) First() bool {
	di.iter = di.k.IterVersioned(di.ctx, di.bucket, di.startKey, di.endKey)
	di.first = true
	return true
}

// Next check if next element exist
func (di *Iterator) Next() bool {
	if !di.first {
		di.iter.Next()
	}
	di.first = false
	return di.iter.Valid()
}

// Key get key for XMIterator
func (di *Iterator) Key() []byte {
	key := di.iter.Key()
	return key[len(types.ExtUtxoTablePrefix):]
}

// Error return error info for XMIterator
func (di *Iterator) Error() error {
	if err := di.iter.Error(); err != nil {
		return err
	}
	return di.err
}

// Release release XMIterator
func (di *Iterator) Release() {
	di.iter.Close()
}

// XMReader xmodel interface for reader
type XMReader struct {
	ctx sdk.Context
	k   Keeper
}

// Get get value for specific key, return value with version
func (s *XMReader) Get(bucket string, key []byte) (*xmodel.VersionedData, error) {
	verData, err := s.k.GetLastVersionedData(s.ctx, bucket, key)
	if err != nil {
		return nil, err
	}
	if verData == nil {
		verData = types.MakeEmptyVersionedData(bucket, key)
	}
	return verData, nil
}

// Select select all kv from a bucket, can set key range, left closed, right opend
func (s *XMReader) Select(bucket string, startKey []byte, endKey []byte) (xmodel.Iterator, error) {
	return NewIterator(s.ctx, s.k, bucket, startKey, endKey), nil
}
