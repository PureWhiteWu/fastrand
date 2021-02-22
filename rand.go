package rand

import (
	impl "math/rand"
	"runtime"
	"time"

	"github.com/PureWhiteWu/ppin"
)

const (
	cacheLineSize = 64
)

var (
	shardsLen int
)

type lockedSource struct {
	_ [cacheLineSize]byte
	*impl.Rand
}

func (ls *lockedSource) Intn(n int) (r int) {
	r = ls.Rand.Intn(n)
	return
}

type FastRand []*lockedSource

func init() {
	shardsLen = runtime.GOMAXPROCS(0)
	defaultRand = New()
}

func New() FastRand {
	s := make([]*lockedSource, shardsLen)
	for i := 0; i < shardsLen; i++ {
		s[i] = &lockedSource{
			Rand: impl.New(impl.NewSource(time.Now().UnixNano())),
		}
	}
	return s
}

func (r FastRand) Intn(n int) int {
	idx := ppin.Pin()
	ret := r[idx%shardsLen].Intn(n)
	ppin.Unpin()
	return ret
}

var defaultRand FastRand

func Intn(n int) int {
	return defaultRand.Intn(n)
}
