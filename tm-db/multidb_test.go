package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func TestMultiDBNewGoLevelDB(t *testing.T) {
	name := fmt.Sprintf("test_%x", randStr(12))
	defer cleanupDBDir("", name)

	// Test we can't open the db twice for writing
	wr1, err := NewMultiDB(name, "", map[string]interface{}{})
	require.Nil(t, err)
	_, err = NewMultiDB(name, "", map[string]interface{}{})
	require.NotNil(t, err)
	wr1.Close() // Close the db to release the lock

	// Test we can open the db twice for reading only
	ro1, err := NewMultiDBWithOpts(name, "", map[string]interface{}{}, &opt.Options{ReadOnly: true})
	require.Nil(t, err)
	defer ro1.Close()
	ro2, err := NewMultiDBWithOpts(name, "", map[string]interface{}{}, &opt.Options{ReadOnly: true})
	require.Nil(t, err)
	defer ro2.Close()
}

func BenchmarkMultiDBRandomReadsWrites(b *testing.B) {
	name := fmt.Sprintf("test_%x", randStr(12))
	db, err := NewMultiDB(name, "", map[string]interface{}{})
	if err != nil {
		b.Fatal(err)
	}
	defer func() {
		db.Close()
		cleanupDBDir("", name)
	}()

	benchmarkRandomReadsWrites(b, db)
}
