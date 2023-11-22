package main

import "fmt"

func typeInferenceMethod() {
	var name string = "Rickyslash"
	occupation := "Software Engineer" // type inference
	fmt.Printf("- %s, %s\n", name, occupation)
}
