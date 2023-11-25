package main

import "fmt"

func recoverMethod() {
	if r := recover(); r != nil {
		fmt.Println("Recovery executed. Error: ", r)
	} else {
		fmt.Println("Program runs perfectly")
	}
}
