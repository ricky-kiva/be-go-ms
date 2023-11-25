package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	sync.Mutex // mutex will handle race condition on concurrency
	value      int
}

func (c *counter) Add() {
	c.Lock() // ensure only 1 goroutine can enter the `counter` struct
	c.value++
	c.Unlock() // releasing the lock
}

// if `mutex` doesn't applied, the `value` in `counter` won't reach 1.000.000
func mutexMethod() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter
	// var mtx sync.Mutex // can also mutex by this variable

	for i := 0; i < 1000; i++ {
		wg.Add(1) // make 1 goroutines on the `WaitGroup` for each `i` iteration
		go func() {
			for j := 0; j < 1000; j++ {
				// mtx.Lock()
				meter.Add() // modify `value` in `counter`
				// mtx.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(meter.value)
}
