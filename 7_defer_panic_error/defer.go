package main

import "fmt"

func deferMethod() {
	defer fmt.Println("End of the program.") // this will be called on the end of this function
	fmt.Println("Welcome!")
}
