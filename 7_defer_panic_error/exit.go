package main

import (
	"fmt"
	"os"
)

func exitMethod() {
	defer fmt.Println("This is will run on the end of this function")
	os.Exit(1) // exit from `main()` execution
	fmt.Println("This will run before exit")
}
