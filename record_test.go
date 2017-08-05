package onecache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRecordHasExpired(t *testing.T) {
	t.Log("When the record has not expired")
	r1 := &record{expiresIn: 1000, timestamp: time.Now()}
	assert.Equal(t, false, r1.hasExpired())

	t.Log("When the record has expired")
	r2 := &record{expiresIn: 10, timestamp: time.Now().Add(time.Second * time.Duration(11) * -1)}
	assert.Equal(t, true, r2.hasExpired())
}
