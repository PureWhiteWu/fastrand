package rand

import (
	"sync"
	"testing"
)

func BenchmarkLockFree(b *testing.B) {
	r := NewLockFree()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Intn(100)
	}
}

func BenchmarkLockFreePar(b *testing.B) {
	r := NewLockFree()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r.Intn(100)
		}
	})
}

func TestIntnLF(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = IntnLF(101)
		}()
	}

	wg.Wait()
}
