package rand

import (
	"runtime"
)

const (
	cacheLineSize = 64
)

var (
	shardsLen       int
	defaultLocked   Locked
	defaultLockFree LockFree
)

func init() {
	shardsLen = runtime.GOMAXPROCS(0)
	defaultLocked = NewLocked()
	defaultLockFree = NewLockFree()
}
