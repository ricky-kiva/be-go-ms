package main

import (
	"fmt"
	"runtime"
)

func channelSelectMethod() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 5, 6, 7, 4, 1, 10}
	fmt.Println("Numbers:", numbers)

	var channelAverage = make(chan float64)
	go getAverage(numbers, channelAverage)

	var channelMax = make(chan int)
	go getMax(numbers, channelMax)

	for i := 0; i < 2; i++ {
		select { // wait on multiple channel operations, select value from channel that first fulfills
		case avg := <-channelAverage:
			fmt.Printf("Average: %.2f\n", avg)
		case max := <-channelMax:
			fmt.Printf("Max: %d\n", max)
		}
	}
}

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, num := range numbers {
		sum += num
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, num := range numbers {
		if max < num {
			max = num
		}
	}

	ch <- max
}
