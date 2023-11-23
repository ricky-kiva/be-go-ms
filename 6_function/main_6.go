package main

import "fmt"

// run by `go run .`
func main() {
	var nameFirstAndLast = []string{"Ricky", "Skywalker"}
	printMessageToPersonFullName("Hello,", nameFirstAndLast)
	printMessageToPersonFullName("Good morning,", nameFirstAndLast)

	area, circumference := calculateCircle(5)
	fmt.Printf("area: %.2f, circumference: %.2f\n", area, circumference)

	var average float64 = calculateAverage(4, 8, 12, 16)
	fmt.Printf("Average: %.2f\n", average)

	// closure function
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	var numbersForMinMax = []int{2, 3, 4, 5, 6, 10, 100, 1, 85, 23, 56}
	var min, max = getMinMax(numbersForMinMax)
	fmt.Printf("min: %d, max: %d\n", min, max)
}
