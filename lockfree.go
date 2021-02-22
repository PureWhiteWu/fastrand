package rand

import (
	impl "math/rand"
	"time"

	"github.com/PureWhiteWu/ppin"
)

type lockFreeSource struct {
	_ [cacheLineSize]byte
	*impl.Rand
}

// LockFree can only be used if GOMAXPROCS isn't adjusted at runtime.
type LockFree []*lockFreeSource

func (l LockFree) ExpFloat64() (r float64) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].ExpFloat64()
	ppin.Unpin()
	return
}

func (l LockFree) NormFloat64() (r float64) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].NormFloat64()
	ppin.Unpin()
	return
}

func (l LockFree) Seed(seed int64) {
	idx := ppin.Pin()
	l.check(idx)
	l[idx].Seed(seed)
	ppin.Unpin()
}

func (l LockFree) Int63() (r int64) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Int63()
	ppin.Unpin()
	return
}

func (l LockFree) Uint32() (r uint32) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Uint32()
	ppin.Unpin()
	return
}

func (l LockFree) Uint64() (r uint64) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Uint64()
	ppin.Unpin()
	return
}

func (l LockFree) Int31() (r int32) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Int31()
	ppin.Unpin()
	return
}

func (l LockFree) Int() (r int) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Int()
	ppin.Unpin()
	return
}

func (l LockFree) Int63n(n int64) (r int64) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Int63n(n)
	ppin.Unpin()
	return
}

func (l LockFree) Int31n(n int32) (r int32) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Int31n(n)
	ppin.Unpin()
	return
}

func (l LockFree) Intn(n int) (r int) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Intn(n)
	ppin.Unpin()
	return
}

func (l LockFree) Float64() (r float64) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Float64()
	ppin.Unpin()
	return
}

func (l LockFree) Float32() (r float32) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Float32()
	ppin.Unpin()
	return
}

func (l LockFree) Perm(n int) (r []int) {
	idx := ppin.Pin()
	l.check(idx)
	r = l[idx].Perm(n)
	ppin.Unpin()
	return
}

func (l LockFree) Shuffle(n int, swap func(i, j int)) {
	idx := ppin.Pin()
	l.check(idx)
	l[idx].Shuffle(n, swap)
	ppin.Unpin()
}

func (l LockFree) Read(p []byte) (n int, err error) {
	idx := ppin.Pin()
	l.check(idx)
	n, err = l[idx].Read(p)
	ppin.Unpin()
	return
}

func (l LockFree) check(idx int) {
	if idx > shardsLen {
		// should not happen
		panic("pid > shard length, maybe GOMAXPROCS is adjusted at runtime")
	}
}

func NewLockFree() LockFree {
	s := make([]*lockFreeSource, shardsLen)
	for i := 0; i < shardsLen; i++ {
		s[i] = &lockFreeSource{
			Rand: impl.New(impl.NewSource(time.Now().UnixNano())),
		}
	}
	return s
}

func ExpFloat64LF() float64 {
	return defaultLockFree.ExpFloat64()
}

func NormFloat64LF() float64 {
	return defaultLockFree.NormFloat64()
}

func SeedLF(seed int64) {
	defaultLockFree.Seed(seed)
}

func Int63LF() int64 {
	return defaultLockFree.Int63()
}

func Uint32LF() uint32 {
	return defaultLockFree.Uint32()
}

func Uint64LF() uint64 {
	return defaultLockFree.Uint64()
}

func Int31LF() int32 {
	return defaultLockFree.Int31()
}

func IntLF() int {
	return defaultLockFree.Int()
}

func Int63nLF(n int64) int64 {
	return defaultLockFree.Int63n(n)
}

func Int31nLF(n int32) int32 {
	return defaultLockFree.Int31n(n)
}

func IntnLF(n int) int {
	return defaultLockFree.Intn(n)
}

func Float64LF() float64 {
	return defaultLockFree.Float64()
}

func Float32LF() float32 {
	return defaultLockFree.Float32()
}

func PermLF(n int) []int {
	return defaultLockFree.Perm(n)
}

func ShuffleLF(n int, swap func(i, j int)) {
	defaultLockFree.Shuffle(n, swap)
}

func ReadLF(p []byte) (int, error) {
	return defaultLockFree.Read(p)
}
