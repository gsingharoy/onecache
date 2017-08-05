package onecache

import "time"

// Represents a cached record
type record struct {
	value     interface{}
	timestamp time.Time // timestamp at which the cache was created
	expiresIn int64     // Expiry time of the cache in seconds
}

func (r *record) hasExpired() bool {
	return r.expiresIn > -1 &&
		time.Now().After(r.timestamp.Add(time.Second*time.Duration(r.expiresIn)))
}
