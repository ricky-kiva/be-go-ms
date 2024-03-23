package main

import (
	"fmt"
	"os"
)

func exitMethod() {
	defer fmt.Println("Due to `exit`, defer won't execute as usual (in the end of function)")
	fmt.Println("This will run before exit")
	os.Exit(1) // exit from `main()` execution
}
