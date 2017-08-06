# Onecache

In memory cache implementation in golang.

[![Build Status](https://travis-ci.org/gsingharoy/onecache.svg?branch=master)](https://travis-ci.org/gsingharoy/onecache)

## Description
This package contains a simple in memory caching functionality. It supports simply a lookup with a cache key, which contains an expiration contract. the cache storage is periodically cleaned from old expired cached values.

**Note** This package is useful for small instances (1-2) and not suitable for complex high traffic systems.

## Usage

```go
import "github.com/gsingharoy/onecache"

// generate a new cache instance
c := onecache.New()

// Any struct
m := &MyAwesomeStruct{
  ...
}

// set entry to the cache
c.Set("my-awesome-key", m, 3600) // will expire in one hour

// lookup for entry in cache

v, found := c.Find("my-awesome-key")
if !found {
  // logic when cache is missed
}
```

In case you want the cache to never expire then pass expiresIn as `-1`
```go
c.Set("my-awesome-key", m, -1)
```
