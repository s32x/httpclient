package cachers /* import "s32x.com/httpclient/cache/cachers" */

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgraph-io/badger/v3"
	"s32x.com/httpclient/cache"
)

// Compile time checking of cache.Cacher implementation
var _ cache.Cacher = (*Badger)(nil)

// Badger is a Cacher implementation that utilizes badger DB
type Badger struct {
	db                *badger.DB
	defaultExpiration time.Duration
}

// NewBadger creates and returns a new Badger struct that contains an
// underlying badger database stored at the passed name path in the systems
// temporary directory
func NewBadger(name string, defaultExpiration time.Duration) (*Badger, error) {
	opts := badger.DefaultOptions(os.TempDir() + name)
	opts.Logger = nil
	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("badger db open: %w", err)
	}
	return &Badger{
		db:                db,
		defaultExpiration: defaultExpiration,
	}, nil
}

// Set persists a new record to the cache with the passed value at the passed
// key and with the passed optional duration
func (b *Badger) Set(key string, val []byte, expiration ...time.Duration) error {
	// Set the expiration if one is not passed
	if len(expiration) == 0 {
		expiration = append(expiration, b.defaultExpiration)
	}

	// Update/Set the record in the database with the expiration
	if err := b.db.Update(func(txn *badger.Txn) error {
		return txn.SetEntry(badger.NewEntry([]byte(key), val).WithTTL(expiration[0]))
	}); err != nil {
		return fmt.Errorf("badger db update: %w", err)
	}
	return nil
}

// Get retrieves a value from the cache and returns it
func (b *Badger) Get(key string) ([]byte, error) {
	var out []byte
	if err := b.db.View(func(txn *badger.Txn) error {
		// Retrieve the item
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		// Read the value from the item
		if err := item.Value(func(val []byte) error {
			out = val
			return nil
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			err = cache.ErrKeyNotFound
		}
		return nil, fmt.Errorf("badger db view: %w", err)
	}
	return out, nil
}
