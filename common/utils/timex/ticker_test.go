package timex

import (
	"sync/atomic"
	"testing"
	"time"
)

// TestRealTickerDoTick 测试说明：模拟Ticker计数10下，每次计数间隔1s
func TestRealTickerDoTick(t *testing.T) {
	ticker := NewTicker(time.Second * 1)
	defer ticker.Stop()

	var count uint32
	for range ticker.Chan() {
		count++
		if count > 10 {
			break
		}
	}
}

func TestSimulateTicker(t *testing.T) {
	var total int32 = 15
	ticker := NewSimulateTicker()
	defer ticker.Stop()

	var count int32 = 0
	go func() {
		for range ticker.Chan() {
			if atomic.AddInt32(&count, 1) == total {
				ticker.Done()
			}
		}
	}()

	for i := 0; i < 8; i++ {
		ticker.DoTick()
	}
}
