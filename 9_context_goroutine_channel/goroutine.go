package main

import (
	"fmt"
	"runtime"
)

func goroutineMethod() {
	runtime.GOMAXPROCS(2)

	go printMultipleTimes(5, "Or no pudding") // it will finish the `for` loop first

	printMultipleTimes(5, "Eat yer meat")

	var input string
	fmt.Scanln(&input)
}

func printMultipleTimes(limit int, message string) {
	for i := 1; i <= limit; i++ {
		fmt.Printf("%d. %s\n", i, message)
	}
}
