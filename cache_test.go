package onecache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()
	if c.Timestamp.After(time.Now()) {
		t.Error("Wrong timestamp generated")
	}
}

func TestCacheSet(t *testing.T) {
	c := New()
	c.Set("sample-key", "sample-value", 1000)
	r, ok := c.recordMap["sample-key"]
	assert.Equal(t, true, ok)
	assert.Equal(t, "sample-value", r.value)
	assert.Equal(t, int64(1000), r.expiresIn)
}

func TestCacheFind(t *testing.T) {
	c := New()
	t.Log("When the record does not exist")
	_, found := c.Find("invalid-cache")
	assert.Equal(t, false, found)

	t.Log("When the record exists but has expired")
	c.recordMap["key1"] = record{
		expiresIn: 10,
		timestamp: time.Now().Add(time.Second * time.Duration(11) * -1),
	}
	_, found = c.Find("key1")
	assert.Equal(t, false, found)

	t.Log("When the record exists but has expired")
	c.recordMap["key2"] = record{
		expiresIn: 1000,
		timestamp: time.Now().Add(time.Second * time.Duration(11) * -1),
		value:     "some-value",
	}
	v, found := c.Find("key2")
	assert.Equal(t, true, found)
	assert.Equal(t, "some-value", v)
}

func TestCache_clean(t *testing.T) {
	c := New()
	c.recordMap["key1"] = record{
		expiresIn: 10,
		timestamp: time.Now().Add(time.Second * time.Duration(11) * -1),
	}
	c.recordMap["key2"] = record{
		expiresIn: 1000,
		timestamp: time.Now().Add(time.Second * time.Duration(11) * -1),
		value:     "some-value",
	}
	c.clean()
	_, ok := c.recordMap["key1"]
	assert.Equal(t, false, ok)
	_, ok = c.recordMap["key2"]
	assert.Equal(t, true, ok)
}
