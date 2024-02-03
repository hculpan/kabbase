package dbbadger

import "github.com/dgraph-io/badger/v3"

func DeleteKey(db *badger.DB, key string) error {
	return db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
}
