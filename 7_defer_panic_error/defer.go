package main

import "fmt"

func deferMethod() {
	defer fmt.Println("End of the program.")
	fmt.Println("Welcome!")
}
