package main

import "fmt"

func main() {
	var value uint8 = (((2+6)%3)*4 - 2) / 3
	fmt.Printf("Result: %d\n", value)
}
