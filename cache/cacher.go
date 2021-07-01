package cache /* import "s32x.com/httpclient/cache" */

import (
	"errors"
	"time"
)

var (
	// ErrKeyNotFound is an error that a cacher returns when a Key doesn't exist
	ErrKeyNotFound = errors.New("key not found")
)

// Cacher is an interface that provides the methods needed to store and retrieve
// data in a a cache
type Cacher interface {
	// Set persists a new record to the cache with the passed value at the
	// passed key and with the passed optional duration
	Set(key string, val []byte, expiration ...time.Duration) error
	// Get retrieves a value from the cache and returns it
	Get(key string) ([]byte, error)
}
