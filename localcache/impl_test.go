package localcache

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {

	timeNow = func() time.Time {
		return time.Unix(1654041600, 0)
	}

	expect := &localCache{
		data: map[string]cacheItem{
			"test": {
				value:    "123",
				expiration: timeNow().Add(30 * time.Second),
			},
		},
	}

	cache := NewLocalCache()

	cache.Set("test", "123")
	assert.Equal(t, reflect.DeepEqual(cache, expect), true, "Should be equal")
}

func TestGet(t *testing.T) {

	cache := &localCache{
		data: map[string]cacheItem{
			"test": {
				value:    "123",
				expiration: timeNow().Add(30 * time.Second),
			},
		},
	}

	value, err := cache.Get("test")
	assert.Nil(t, err, "Should be nil")
	assert.Equal(t, value, "123", "Should be equal")
}

func TestGet_ErrKeyNotExist(t *testing.T) {
	c := NewLocalCache()

	val, err := c.Get("a")

	assert.Nil(t, val)
	assert.Equal(t, err, ErrKeyNotExist, "key should not exist.")
}

func TestGet_ErrKeyExpired(t *testing.T) {
	cache := &localCache{
		data: map[string]cacheItem{
			"test": {
				value:    "123",
				expiration: timeNow().Add(30 * time.Second),
			},
		},
	}

	// mock time.Time.Before()
	timeBefore = func(_ time.Time, _ time.Time) bool {
		return true
	}
	val, err := cache.Get("test")

	assert.Nil(t, val)
	assert.Equal(t, err, ErrKeyExpired, "key should expired.")
}
