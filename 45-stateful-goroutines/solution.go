//go:build solution

package statefulgoroutines

import (
	"sync"
	"sync/atomic"
)

type ReadOp struct {
	Key  int
	Resp chan int
}

type WriteOp struct {
	Key  int
	Val  int
	Resp chan bool
}

func StartStore() (reads chan<- ReadOp, writes chan<- WriteOp) {
	readCh := make(chan ReadOp)
	writeCh := make(chan WriteOp)

	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-readCh:
				read.Resp <- state[read.Key]
			case write := <-writeCh:
				state[write.Key] = write.Val
				write.Resp <- true
			}
		}
	}()

	return readCh, writeCh
}

func Read(reads chan<- ReadOp, key int) int {
	read := ReadOp{Key: key, Resp: make(chan int)}
	reads <- read
	return <-read.Resp
}

func Write(writes chan<- WriteOp, key, val int) {
	write := WriteOp{Key: key, Val: val, Resp: make(chan bool)}
	writes <- write
	<-write.Resp
}

func CountOps(readers, writers, opsEach int) (readOps, writeOps uint64) {
	reads, writes := StartStore()
	var nReads, nWrites atomic.Uint64
	var wg sync.WaitGroup

	for r := range readers {
		wg.Go(func() {
			for i := range opsEach {
				Read(reads, (r+i)%5)
				nReads.Add(1)
			}
		})
	}
	for w := range writers {
		wg.Go(func() {
			for i := range opsEach {
				Write(writes, (w+i)%5, i)
				nWrites.Add(1)
			}
		})
	}

	wg.Wait()
	return nReads.Load(), nWrites.Load()
}

type Counter struct {
	incs  chan struct{}
	reads chan chan int
	quit  chan struct{}
}

func NewCounter() *Counter {
	c := &Counter{
		incs:  make(chan struct{}),
		reads: make(chan chan int),
		quit:  make(chan struct{}),
	}
	go func() {
		count := 0
		for {
			select {
			case <-c.incs:
				count++
			case resp := <-c.reads:
				resp <- count
			case <-c.quit:
				return
			}
		}
	}()
	return c
}

func (c *Counter) Inc() {
	// Unbuffered: the send completes when the goroutine has taken it.
	c.incs <- struct{}{}
}

func (c *Counter) Value() int {
	resp := make(chan int)
	c.reads <- resp
	return <-resp
}

func (c *Counter) Close() {
	close(c.quit)
}
