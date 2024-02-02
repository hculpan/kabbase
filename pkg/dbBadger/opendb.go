package dbbadger

import (
	"github.com/dgraph-io/badger/v3"
)

func OpenDB(dbPath string) (*badger.DB, error) {
	return badger.Open(badger.DefaultOptions(dbPath))
}
