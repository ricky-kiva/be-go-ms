package main

import "fmt"

func operatorsMethod() {
	// arithmetics operator
	var value uint8 = (((2+6)%3)*4 - 2) / 3
	fmt.Printf("Result: %d\n", value)

	// relational operator
	var two uint8 = 2
	var three uint8 = 3
	var resultIsTwo bool = (value == two)
	var resultIsThree bool = (value == three)
	fmt.Printf("Result is equal to %d? %t\n", two, resultIsTwo)
	fmt.Printf("Result is equal to %d? %t\n", three, resultIsThree)

	// logic operator
	var andTrue bool = true && true
	var andFalse bool = false && true
	var orTrue bool = true || false
	var orFalse bool = false || false
	var notFalse bool = !true

	fmt.Printf("Logic variable answers: %t, %t, %t, %t, %t\n", andTrue, andFalse, orTrue, orFalse, notFalse)
}
