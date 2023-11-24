package main

import (
	"fmt"
	"strconv"
)

func errorMethod() {
	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)
	var number int
	var err error
	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number, "is a number")
	} else {
		fmt.Println(input, "is not a number")
		fmt.Println("Error ->", err.Error())
	}
}
