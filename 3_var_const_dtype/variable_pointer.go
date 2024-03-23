package main

import "fmt"

func variablePointerMethod() {
	var pointer4 uint8 = 4
	var pointer8 *uint8 = &pointer4 // getting reference of `pointer4`
	*pointer8 = 1                   // assigning value to the referenced variable of `pointer8` (`pointer4`)

	fmt.Printf("Value of pointer4 now: %d\n", pointer4)
	fmt.Printf("Value of pointer8 now: %d\n", *pointer8)

	fmt.Printf("Value of pointer4 now: %d\n", pointer4)
}
