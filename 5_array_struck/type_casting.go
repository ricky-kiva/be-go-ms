package main

import "fmt"

func typeCastMethod() {
	var number int = 69
	var floatNumber float32 = float32(number)
	fmt.Printf("Number is now a float %.1f\n", floatNumber)
}
