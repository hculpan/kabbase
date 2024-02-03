package dbbadger

import "github.com/dgraph-io/badger/v3"

func GetKeyValuesWithPrefix(db *badger.DB, prefix string) (map[string][]byte, error) {
	keyValues := make(map[string][]byte)
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = []byte(prefix)
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.ValidForPrefix(opts.Prefix); it.Next() {
			item := it.Item()
			key := string(item.Key())

			err := item.Value(func(val []byte) error {
				keyValues[key] = val
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return keyValues, err
}
