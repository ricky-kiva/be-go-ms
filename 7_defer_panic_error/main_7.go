package main

import "fmt"

func main() {
	deferMethod()
	errorMethod()

	// custom error
	var name string
	fmt.Print("Your name: ")
	fmt.Scanln(&name)
	if valid, err := validateFilledString(name); valid {
		fmt.Println("Hello,", name)
	} else {
		fmt.Println(err.Error())
	}

	exitMethod()
}
