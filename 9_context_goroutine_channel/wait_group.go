package main

import (
	"fmt"
	"runtime"
	"sync"
)

func waitGroupMethod() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data no: %d", i+1)
		wg.Add(1) // tells goroutine there is 1 more goroutine he should wait for
		go printAndEndWaitGroup(&wg, data)
	}
	wg.Wait() // waits all goroutines to complete (WaitGroup = 0). if not called, the program could ends prematurely (missing some goroutines to be processed)
}

func printAndEndWaitGroup(wg *sync.WaitGroup, message string) {
	defer wg.Done() // decrement the WaitGroup. called on the end of the function
	fmt.Println(message)
}
