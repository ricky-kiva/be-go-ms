package main

import "fmt"

func operatorsMethod() {
	// arithmetics
	var value uint8 = (((2+6)%3)*4 - 2) / 3
	fmt.Printf("Result: %d\n", value)

	// relational
	var two uint8 = 2
	var three uint8 = 3
	var resultIsTwo bool = (value == two)
	var resultIsThree bool = (value == three)
	fmt.Printf("Result is equal to %d? %t\n", two, resultIsTwo)
	fmt.Printf("Result is equal to %d? %t\n", three, resultIsThree)
}
