package main

import "fmt"

// run by `go run .`
func main() {
	// recover
	defer recoverMethod()

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

	// comment these to simulate non-`recover`-y mode
	panicMethod()
	exitMethod()
}
