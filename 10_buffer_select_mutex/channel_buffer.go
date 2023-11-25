package main

import (
	"fmt"
	"runtime"
	"time"
)

func bufferChannelMethod() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3) // make channel with buffer size

	go func() { // launch goroutines
		for {
			receivedData := <-messages // continuously receives data from `messages` channel
			fmt.Println("Received data:", receivedData)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Send data:", i)
		messages <- i // sends data to `messages` channel
	}

	time.Sleep(1 * time.Second) // will wait for all the goroutines to be processed
}
