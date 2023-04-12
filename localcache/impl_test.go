package localcache

import (
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	cache := NewLocalCache()

	timeNow = func() time.Time {
		return time.Unix(1654041600, 0)
	}

	type testCase struct {
		desc      string
		key       string
		value     interface{}
		expectVal interface{}
		err  error
	}

	testCases := []testCase{
		{desc: "first set", key: "key1", value: "value1", expectVal: "value1", err: nil},
		{desc: "second set and cover value ", key: "key1", value: "value2", expectVal: "value2", err: nil},
	}

	for _, tc := range testCases {
		cache.Set(tc.key, tc.value)

		if val := cache.data["key1"]; val.value != tc.expectVal || tc.err != nil {
			t.Errorf("Set failed")
		}
	}
}

func TestGet(t *testing.T) {

	type testCase struct {
		desc      string
		key       string
		expectVal interface{}
		err  error
		expired bool
	}

	cache := &localCache{
		data: map[string]cacheItem{
			"key": {
				value:    "value",
				expiration: timeNow().Add(30 * time.Second),
			},
		},
	}

	testCases := []testCase{
		{desc: "get success", key: "key", expectVal: "value", err: nil, expired: false},
		{desc: "get error ErrKeyNotExist", key: "key1", expectVal: nil, err: ErrKeyNotExist, expired: false},
		{desc: "get error ErrKeyExpired", key: "key", expectVal: nil, err: ErrKeyExpired, expired: true},
	}

	for _, tc := range testCases {
		if tc.expired {
		  // mock time.Time.Before()
	      timeBefore = func(_ time.Time, _ time.Time) bool {
		    return true
	      }
		}

		if val, err := cache.Get(tc.key); val != tc.expectVal || tc.err != err {
			t.Errorf("Get failed")
		}
	}
}
