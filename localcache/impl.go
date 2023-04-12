package localcache

import (
	"sync"
	"time"
)

var (
	timeBefore = time.Time.Before
	timeNow = time.Now
)

type localCache struct {
	data map[string]cacheItem
	sync.RWMutex
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

func NewLocalCache() *localCache {
	return &localCache{data: make(map[string]cacheItem)}
}

func (lc *localCache) Get(key string) (interface{}, error) {
	lc.RLock()
	defer lc.RUnlock()

	item, found := lc.data[key]
	if !found {
		return nil, ErrKeyNotExist
	}

	if timeBefore(item.expiration, timeNow()) {
		delete(lc.data, key)
		return nil, ErrKeyExpired
	}

	return item.value, nil
}

func (lc *localCache) Set(key string, value interface{}) {
	lc.Lock()
	defer lc.Unlock()

	lc.data[key] = cacheItem{
		value:      value,
		expiration: timeNow().Add(defaultExTime),
	}
}
