package main

import "fmt"

func panicMethod() {
	defer fmt.Println("Defer will still be executed after Panic")
	panic("This is panic message") // unlike `exit`, `defer` will still be called on `panic`
	fmt.Println("This code is unreachable because of `panic`")
}
