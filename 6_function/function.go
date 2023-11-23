package main

import (
	"fmt"
	"strings"
)

func printMessageToPersonFullName(message string, arr []string) {
	var joined string = strings.Join(arr, " ")
	fmt.Println(message, joined)
}
