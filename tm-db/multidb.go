package db

import (
	"fmt"
	"github.com/tendermint/tm-db/mstorage"
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func init() {
	dbCreator := func(name string, dir string, options map[string]interface{}) (DB, error) {
		return NewMultiDB(name, dir, options)
	}
	registerDBCreator(MultiDBBackend, dbCreator, false)
}

type MultiDB struct {
	db *leveldb.DB
}

var _ DB = (*MultiDB)(nil)

func NewMultiDB(name string, dir string, options map[string]interface{}) (*MultiDB, error) {
	return NewMultiDBWithOpts(name, dir, options, nil)
}

func NewMultiDBWithOpts(name string, dir string, options map[string]interface{}, o *opt.Options) (*MultiDB, error) {
	dbPath := filepath.Join(dir, name+".db")
	if options == nil || len(options) == 0 {
		db, err := leveldb.OpenFile(dbPath, o)
		if err != nil {
			return nil, err
		}
		return &MultiDB{
			db: db,
		}, nil
	}

	//多盘存储初始化
	dataPaths := options["dataPaths"].([]string)
	store, err := mstorage.OpenFile(dbPath, false, dataPaths)
	if err != nil {
		return nil, err
	}
	db, err := leveldb.Open(store, o)
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		//db, err = leveldb.Recover(store, nil)
		return nil, err
	}
	// (Re)check for errors and abort if opening of the db failed
	if err != nil {
		return nil, err
	}
	return &MultiDB{
		db: db,
	}, nil
}

// Get implements DB.
func (db *MultiDB) Get(key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, errKeyEmpty
	}
	res, err := db.db.Get(key, nil)
	if err != nil {
		if err == errors.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return res, nil
}

// Has implements DB.
func (db *MultiDB) Has(key []byte) (bool, error) {
	bytes, err := db.Get(key)
	if err != nil {
		return false, err
	}
	return bytes != nil, nil
}

// Set implements DB.
func (db *MultiDB) Set(key []byte, value []byte) error {
	if len(key) == 0 {
		return errKeyEmpty
	}
	if value == nil {
		return errValueNil
	}
	if err := db.db.Put(key, value, nil); err != nil {
		return err
	}
	return nil
}

// SetSync implements DB.
func (db *MultiDB) SetSync(key []byte, value []byte) error {
	if len(key) == 0 {
		return errKeyEmpty
	}
	if value == nil {
		return errValueNil
	}
	if err := db.db.Put(key, value, &opt.WriteOptions{Sync: true}); err != nil {
		return err
	}
	return nil
}

// Delete implements DB.
func (db *MultiDB) Delete(key []byte) error {
	if len(key) == 0 {
		return errKeyEmpty
	}
	if err := db.db.Delete(key, nil); err != nil {
		return err
	}
	return nil
}

// DeleteSync implements DB.
func (db *MultiDB) DeleteSync(key []byte) error {
	if len(key) == 0 {
		return errKeyEmpty
	}
	err := db.db.Delete(key, &opt.WriteOptions{Sync: true})
	if err != nil {
		return err
	}
	return nil
}

func (db *MultiDB) DB() *leveldb.DB {
	return db.db
}

// Close implements DB.
func (db *MultiDB) Close() error {
	if err := db.db.Close(); err != nil {
		return err
	}
	return nil
}

// Print implements DB.
func (db *MultiDB) Print() error {
	str, err := db.db.GetProperty("leveldb.stats")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", str)

	itr := db.db.NewIterator(nil, nil)
	for itr.Next() {
		key := itr.Key()
		value := itr.Value()
		fmt.Printf("[%X]:\t[%X]\n", key, value)
	}
	return nil
}

// Stats implements DB.
func (db *MultiDB) Stats() map[string]string {
	keys := []string{
		"leveldb.num-files-at-level{n}",
		"leveldb.stats",
		"leveldb.sstables",
		"leveldb.blockpool",
		"leveldb.cachedblock",
		"leveldb.openedtables",
		"leveldb.alivesnaps",
		"leveldb.aliveiters",
	}

	stats := make(map[string]string)
	for _, key := range keys {
		str, err := db.db.GetProperty(key)
		if err == nil {
			stats[key] = str
		}
	}
	return stats
}

// NewBatch implements DB.
func (db *MultiDB) NewBatch() Batch {
	return newMultiDBBatch(db)
}

// Iterator implements DB.
func (db *MultiDB) Iterator(start, end []byte) (Iterator, error) {
	if (start != nil && len(start) == 0) || (end != nil && len(end) == 0) {
		return nil, errKeyEmpty
	}
	itr := db.db.NewIterator(&util.Range{Start: start, Limit: end}, nil)
	return newMultiDBIterator(itr, start, end, false), nil
}

// ReverseIterator implements DB.
func (db *MultiDB) ReverseIterator(start, end []byte) (Iterator, error) {
	if (start != nil && len(start) == 0) || (end != nil && len(end) == 0) {
		return nil, errKeyEmpty
	}
	itr := db.db.NewIterator(&util.Range{Start: start, Limit: end}, nil)
	return newMultiDBIterator(itr, start, end, true), nil
}
