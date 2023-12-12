package pool

import (
	"sync"
	"sync/atomic"
	"testing"
)

func DoCopyStack(a, b int) int {
	if b < 100 {
		return DoCopyStack(0, b+1)
	}
	return 0
}

func testPanicFunc() {
	panic("test")
}

func TestPool(t *testing.T) {
	p := NewPool("test", 100, 0)
	var n int32
	var wg sync.WaitGroup
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		p.Go(func() {
			defer wg.Done()
			atomic.AddInt32(&n, 1)
		})
	}
	wg.Wait()
	if n != 2000 {
		t.Error(n)
	}
}

func TestPoolPanic(t *testing.T) {
	p := NewPool("test", 100, 0)
	p.Go(testPanicFunc)
}
