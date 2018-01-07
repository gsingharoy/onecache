package onecache

import "time"

type Cache struct {
	recordMap map[string]record
	Timestamp time.Time
}

// Returns a new Cache instance
func New() *Cache {
	c := &Cache{
		recordMap: make(map[string]record),
		Timestamp: time.Now(),
	}
	return c
}

// Sets an entry to the cached records
//
// @params
// key:         [string] key of the record
// v:           [interface{}] any record which should be the value of the cache
// expiresIn:   [int64] time in seconds in which the cached value will be expired. If the cache is supposed to
//				be stored for an infinite time then keep this value -1
func (c *Cache) Set(key string, v interface{}, expiresIn int64) {
	// TODO: Introduce a layered key storage so that the same mutex is not used always.
	cMutex.Lock()
	defer cMutex.Unlock()
	c.recordMap[key] = record{
		value:     v,
		timestamp: time.Now(),
		expiresIn: expiresIn,
	}
}

// Finds the record in the cache.
//
// @params
// key: [string] key of the record
//
// @return
// interface{} value of the record
// bool : indicates if the record has been found or not
func (c *Cache) Find(key string) (interface{}, bool) {
	r, ok := c.recordMap[key]
	if !ok || r.hasExpired() {
		return nil, false
	}
	// To optimize this, this go routine can be a channel which will accept all clean() requests
	go c.clean()
	return r.value, true
}
