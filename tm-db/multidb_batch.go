package db

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type multiDBBatch struct {
	db    *MultiDB
	batch *leveldb.Batch
	size  int
	keys  map[string]bool
}

var _ Batch = (*multiDBBatch)(nil)

func newMultiDBBatch(db *MultiDB) *multiDBBatch {
	return &multiDBBatch{db: db, batch: new(leveldb.Batch), keys: map[string]bool{}}
}

// Set implements Batch.
func (b *multiDBBatch) Set(key, value []byte) error {
	if len(key) == 0 {
		return errKeyEmpty
	}
	if value == nil {
		return errValueNil
	}
	if b.batch == nil {
		return errBatchClosed
	}
	b.batch.Put(key, value)
	return nil
}

// Delete implements Batch.
func (b *multiDBBatch) Delete(key []byte) error {
	if len(key) == 0 {
		return errKeyEmpty
	}
	if b.batch == nil {
		return errBatchClosed
	}
	b.batch.Delete(key)
	return nil
}

// Write implements Batch.
func (b *multiDBBatch) Write() error {
	return b.write(false)
}

// WriteSync implements Batch.
func (b *multiDBBatch) WriteSync() error {
	return b.write(true)
}

func (b *multiDBBatch) write(sync bool) error {
	if b.batch == nil {
		return errBatchClosed
	}
	err := b.db.db.Write(b.batch, &opt.WriteOptions{Sync: sync})
	if err != nil {
		return err
	}
	// Make sure batch cannot be used afterwards. Callers should still call Close(), for errors.
	return b.Close()
}

// Close implements Batch.
func (b *multiDBBatch) Close() error {
	if b.batch != nil {
		b.batch.Reset()
		b.batch = nil
	}
	return nil
}
