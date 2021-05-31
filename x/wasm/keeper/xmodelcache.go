package keeper

import (
	"encoding/json"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	"github.com/liubaninc/m0/x/wasm/types"
	"github.com/syndtr/goleveldb/leveldb/comparer"
	"github.com/syndtr/goleveldb/leveldb/memdb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/liubaninc/m0/x/wasm/xmodel"
)

var _ xmodel.XModel = &XMCache{}

const (
	// DefaultMemDBSize 默认内存db大小
	DefaultMemDBSize = 32
)

var (
	// ErrHasDel is returned when key was marked as del
	ErrHasDel = errors.New("Key has been mark as del")
	// ErrNotFound is returned when key is not found
	ErrNotFound = errors.New("Key not found")
)

// XMCache data structure for XModel Cache
type XMCache struct {
	inputsCache  *memdb.DB // Key: bucket_key; Value: VersionedData
	outputsCache *memdb.DB // Key: bucket_key; Value: PureData
	isPenetrate  bool      // 是否穿透到model层
	model        xmodel.XMReader
	utxoCache    *UtxoCache
	events       []*xmodel.ContractEvent
}

// NewXModelCache new an instance of XModel Cache
func NewXModelCache(model xmodel.XMReader, utxovm *UtxoReader) (*XMCache, error) {
	return &XMCache{
		isPenetrate:  true,
		model:        model,
		inputsCache:  memdb.New(comparer.DefaultComparer, DefaultMemDBSize),
		outputsCache: memdb.New(comparer.DefaultComparer, DefaultMemDBSize),
		utxoCache:    NewUtxoCache(utxovm),
	}, nil
}

// NewXModelCacheWithInputs make new XModelCache with Inputs
func NewXModelCacheWithInputs(vdatas []*xmodel.VersionedData, utxoInputs []*utxotypes.Input) *XMCache {
	xc := &XMCache{
		isPenetrate:  false,
		inputsCache:  memdb.New(comparer.DefaultComparer, DefaultMemDBSize),
		outputsCache: memdb.New(comparer.DefaultComparer, DefaultMemDBSize),
	}
	for _, vd := range vdatas {
		rawKey := types.MakeRawKey(vd.PureData.Bucket, vd.PureData.Key)
		valBuf, _ := json.Marshal(vd)
		xc.inputsCache.Put(rawKey, valBuf)
	}
	xc.utxoCache = NewUtxoCacheWithInputs(utxoInputs)
	return xc
}

// Get 读取一个key的值，返回的value就是有版本的data
func (xc *XMCache) Get(bucket string, key []byte) (*xmodel.VersionedData, error) {
	// Level1: get from outputsCache
	verData, err := xc.getFromOuputsCache(bucket, key)
	if err != nil && err != memdb.ErrNotFound {
		return nil, err
	}
	if verData != nil {
		if types.IsDelFlag(verData.PureData.Value) {
			return nil, ErrHasDel
		}
		return verData, nil
	}

	// Level2: get and set from inputsCache
	verData, err = xc.getAndSetFromInputsCache(bucket, key)
	if err != nil {
		return nil, err
	}
	if types.IsEmptyVersionedData(verData) {
		return nil, ErrNotFound
	}
	if types.IsDelFlag(verData.PureData.Value) {
		return nil, ErrHasDel
	}
	return verData, nil
}

// Level1 读取，从outputsCache中读取
func (xc *XMCache) getFromOuputsCache(bucket string, key []byte) (*xmodel.VersionedData, error) {
	buKey := types.MakeRawKey(bucket, key)
	val, err := xc.outputsCache.Get(buKey)
	if err != nil {
		return nil, err
	}

	data := &xmodel.VersionedData{}
	if err = json.Unmarshal(val, data); err != nil {
		return nil, err
	}
	return data, nil
}

// Level2 读取，从inputsCache中读取, 读取不到的情况下，如果isPenetrate为true，会更深一层次从model里读取，并且会将内容填充到readSets中
func (xc *XMCache) getAndSetFromInputsCache(bucket string, key []byte) (*xmodel.VersionedData, error) {
	buKey := types.MakeRawKey(bucket, key)
	valBuf, err := xc.inputsCache.Get(buKey)
	if err != nil && err != memdb.ErrNotFound {
		return nil, err
	}
	if err == memdb.ErrNotFound {
		if !xc.isPenetrate {
			return nil, err
		}
		err := xc.setInputCache(buKey)
		if err != nil {
			return nil, err
		}
	}
	valBuf, err = xc.inputsCache.Get(buKey)
	data := &xmodel.VersionedData{}
	if err = json.Unmarshal(valBuf, data); err != nil {
		return nil, err
	}
	return data, nil
}

// Put put a pair of <key, value> into XModel Cache
func (xc *XMCache) Put(bucket string, key []byte, value []byte) error {
	buKey := types.MakeRawKey(bucket, key)
	val := &xmodel.VersionedData{
		PureData: &xmodel.PureData{
			Key:    key,
			Value:  value,
			Bucket: bucket,
		},
	}
	valBuf, err := json.Marshal(val)
	if err != nil {
		return err
	}
	if bucket != types.TransientBucket {
		xc.Get(bucket, key)
	}
	return xc.outputsCache.Put(buKey, valBuf)
}

// Del delete one key from outPutCache, marked its value as `DelFlag`
func (xc *XMCache) Del(bucket string, key []byte) error {
	return xc.Put(bucket, key, []byte(types.DelFlag))
}

// Select select all kv from a bucket, can set key range, left closed, right opend
// When xc.isPenetrate equals true, three-way merge, When xc.isPenetrate equals false, two-way merge
func (xc *XMCache) Select(bucket string, startKey []byte, endKey []byte) (xmodel.Iterator, error) {
	return xc.NewXModelCacheIterator(bucket, startKey, endKey, comparer.DefaultComparer)
}

// GetRWSets get read/write sets
func (xc *XMCache) GetRWSets() ([]*xmodel.VersionedData, []*xmodel.PureData, error) {
	readSets, err := xc.getReadSets()
	if err != nil {
		return nil, nil, err
	}
	writeSets, err := xc.getWriteSets()
	if err != nil {
		return nil, nil, err
	}
	return readSets, writeSets, nil
}

func (xc *XMCache) getReadSets() ([]*xmodel.VersionedData, error) {
	var readSets []*xmodel.VersionedData
	iter := xc.inputsCache.NewIterator(&util.Range{Start: nil, Limit: nil})
	defer iter.Release()
	for iter.Next() {
		val := iter.Value()
		vd := &xmodel.VersionedData{}
		err := json.Unmarshal(val, vd)
		if err != nil {
			return nil, err
		}
		readSets = append(readSets, vd)
	}
	return readSets, nil
}

func (xc *XMCache) getWriteSets() ([]*xmodel.PureData, error) {
	var writeSets []*xmodel.PureData
	iter := xc.outputsCache.NewIterator(&util.Range{Start: nil, Limit: nil})
	defer iter.Release()
	for iter.Next() {
		val := iter.Value()
		vd := &xmodel.VersionedData{}
		err := json.Unmarshal(val, vd)
		if err != nil {
			return nil, err
		}
		writeSets = append(writeSets, vd.PureData)
	}
	return writeSets, nil
}

// Transfer transfer tokens using utxo
func (xc *XMCache) Transfer(from, to string, amount string) error {
	fromAddr, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		return errors.New("parse from address error")
	}
	toAddr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		return errors.New("parse to address error")
	}
	coins, err := sdk.ParseCoinsNormalized(amount)
	if err != nil {
		return errors.New("parse amount error")
	}
	return xc.utxoCache.Transfer(fromAddr, toAddr, coins)
}

// AddEvent add contract event to xmodel cache
func (xc *XMCache) AddEvent(events ...*xmodel.ContractEvent) {
	xc.events = append(xc.events, events...)
}

// GetUtxoRWSets returns the inputs and outputs of utxo
func (xc *XMCache) GetUtxoRWSets() ([]*utxotypes.Input, []*utxotypes.Output) {
	return xc.utxoCache.GetRWSets()
}

// writeUtxoRWSet write utxos to TransientBucket
func (xc *XMCache) writeUtxoRWSet() error {
	inputs, outputs := xc.utxoCache.GetRWSets()
	var in, out []byte
	var err error
	if len(inputs) != 0 {
		in, err = json.Marshal(inputs)
		if err != nil {
			return err
		}
	}
	if len(outputs) != 0 {
		out, err = json.Marshal(outputs)
		if err != nil {
			return err
		}
	}
	if in != nil {
		err = xc.Put(types.TransientBucket, []byte(types.ContractUtxoInputKey), in)
		if err != nil {
			return err
		}
	}
	if out != nil {
		err = xc.Put(types.TransientBucket, []byte(types.ContractUtxoOutputKey), out)
		if err != nil {
			return err
		}
	}
	return nil
}

func (xc *XMCache) writeEventRWSet() error {
	if len(xc.events) == 0 {
		return nil
	}
	buf, err := json.Marshal(xc.events)
	if err != nil {
		return err
	}
	return xc.Put(types.TransientBucket, []byte(types.ContractEventKey), buf)
}

// WriteTransientBucket write transient bucket data.
// transient bucket is a special bucket used to store some data
// generated during the execution of the contract, but will not be referenced by other txs.
func (xc *XMCache) WriteTransientBucket() error {
	err := xc.writeUtxoRWSet()
	if err != nil {
		return err
	}

	err = xc.writeEventRWSet()
	if err != nil {
		return err
	}
	return nil
}

func (xc *XMCache) isDel(rawKey []byte) bool {
	val, err := xc.outputsCache.Get(rawKey)
	if err == memdb.ErrNotFound {
		return false
	}
	data := &xmodel.VersionedData{}
	err = json.Unmarshal(val, data)
	if err != nil {
		return false
	}
	return types.IsDelFlag(data.GetPureData().GetValue())
}

func (xc *XMCache) setInputCache(rawKey []byte) error {
	if val, _ := xc.inputsCache.Get(rawKey); val != nil {
		return nil
	}
	bucket, key, err := types.ParseRawKey(rawKey)
	if err != nil {
		return err
	}
	val, err := xc.model.Get(bucket, key)
	if err != nil {
		return err
	}
	valBuf, _ := json.Marshal(val)
	return xc.inputsCache.Put(rawKey, valBuf)
}
