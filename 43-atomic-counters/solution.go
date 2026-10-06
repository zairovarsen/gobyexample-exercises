//go:build solution

package atomiccounters

import (
	"sync"
	"sync/atomic"
)

func CountOps(goroutines, perGoroutine int) uint64 {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for range goroutines {
		wg.Go(func() {
			for range perGoroutine {
				ops.Add(1)
			}
		})
	}

	wg.Wait()
	return ops.Load()
}

type Counter struct {
	n atomic.Int64
}

func (c *Counter) Inc() {
	c.n.Add(1)
}

func (c *Counter) Add(delta int64) {
	c.n.Add(delta)
}

func (c *Counter) Value() int64 {
	return c.n.Load()
}

type Flag struct {
	set atomic.Bool
}

func (f *Flag) TrySet() bool {
	return f.set.CompareAndSwap(false, true)
}

func (f *Flag) IsSet() bool {
	return f.set.Load()
}

type HighWater struct {
	max atomic.Int64
}

func (h *HighWater) Observe(v int64) {
	for {
		cur := h.max.Load()
		if v <= cur || h.max.CompareAndSwap(cur, v) {
			return
		}
	}
}

func (h *HighWater) Max() int64 {
	return h.max.Load()
}
