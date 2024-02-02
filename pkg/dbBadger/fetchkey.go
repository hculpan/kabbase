package dbbadger

import (
	"github.com/dgraph-io/badger/v3"
)

func FetchKey(db *badger.DB, key string) ([]byte, error) {
	result := []byte{}
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		// Retrieve the value
		err = item.Value(func(val []byte) error {
			result = val
			return nil
		})

		return err
	})

	return result, err
}
