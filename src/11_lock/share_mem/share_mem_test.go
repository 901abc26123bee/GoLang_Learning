package share_mem_test

import (
	"sync"
	"testing"
	"time"
)
// RWLocl vs Mutex

func TestCounter(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 1) // avoid TestCounter run faster than for loop inside its
	t.Logf("counter: %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait() // precisely waits until till all groutine(for loop) finish
	t.Logf("counter: %d", counter)
}