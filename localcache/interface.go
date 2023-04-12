package localcache

import (
	"errors"
	"time"
)

var (
	defaultExTime = time.Second * 30
	// ErrKeyNotExist means that key doesn't exist.
	ErrKeyNotExist = errors.New("key does not exist.")
	// ErrKeyExpired means that key is expired.
	ErrKeyExpired = errors.New("key expired.")
)

// Cache is used to record data that user frequently fetch
type Cache interface {
	// Get is used to fetch value by key.
	Get(key string) (interface{}, error)
	// Set is used to set value by key
	Set(key string, val interface{})
}