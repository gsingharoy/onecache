package onecache

import (
	"testing"
	"time"
)

func TestRecordHasExpired(t *testing.T) {
	t.Log("When the record has not expired")
	r1 := &record{expiresIn: 1000, timestamp: time.Now()}
	if r1.hasExpired() {
		t.Error("Expected the record to not have expired")
	}

	t.Log("When the record has expired")
	r2 := &record{expiresIn: 10, timestamp: time.Now().Add(time.Second * time.Duration(11) * -1)}
	if !r2.hasExpired() {
		t.Error("Expected the record to have expired")
	}
}
