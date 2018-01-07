package onecache

import (
	"math/rand"
	"sync"
	"time"
)

var cMutex sync.Mutex

var seed = rand.NewSource(time.Now().UnixNano())
var randGen = rand.New(seed)
