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
}
