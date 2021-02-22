package rand

import (
	impl "math/rand"
	"sync"
	"testing"
)

func BenchmarkRand(b *testing.B) {
	r := NewRand()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Intn(100)
	}
}

func BenchmarkRandStd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.Intn(100)
	}
}

func BenchmarkRandPar(b *testing.B) {
	r := NewRand()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r.Intn(100)
		}
	})
}

func BenchmarkRandParStd(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			impl.Intn(100)
		}
	})
}

func TestRand(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = Intn(101)
		}()
	}

	wg.Wait()
}
