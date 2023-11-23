package main

import (
	"fmt"
	"strings"
)

// run by `go run .`
func main() {
	// function
	var nameFirstAndLast = []string{"Ricky", "Skywalker"}
	printMessageToPersonFullName("Hello,", nameFirstAndLast)
	printMessageToPersonFullName("Good morning,", nameFirstAndLast)

	// multiple return function
	area, circumference := calculateCircle(5)
	fmt.Printf("area: %.2f, circumference: %.2f\n", area, circumference)

	// variadic function
	var average float64 = calculateAverage(4, 8, 12, 16)
	fmt.Printf("Average: %.2f\n", average)

	// closure function
	var getMinMax = getMinMaxClosure()

	var numbersForMinMax = []int{2, 3, 4, 5, 6, 10, 100, 1, 85, 23, 56}
	var min, max = getMinMax(numbersForMinMax)
	fmt.Printf("min: %d, max: %d\n", min, max)

	var filterOString = func(input string) bool {
		input = strings.ToLower(input)
		return strings.Contains(input, "o")
	}

	var stringsToFilter = []string{"Yogyakarta", "Okinawa", "Jakarta", "Tokyo", "San Francisco", "Vienna", "Munich", "London"}
	var oStringArray = filterCallback(stringsToFilter, filterOString)
	fmt.Println(oStringArray)
}
