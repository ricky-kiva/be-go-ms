package main

import "fmt"

// run by `go run .`
func main() {
	var nameFirstAndLast = []string{"Ricky", "Skywalker"}
	printMessageToPersonFullName("Hello,", nameFirstAndLast)
	printMessageToPersonFullName("Good morning,", nameFirstAndLast)

	area, circumference := calculateCircle(5)
	fmt.Printf("area: %.2f, circumference: %.2f\n", area, circumference)
}
