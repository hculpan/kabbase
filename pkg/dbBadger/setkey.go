package dbbadger

import "github.com/dgraph-io/badger/v3"

func SetKey(db *badger.DB, key string, value []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), value)
		return err
	})
}
