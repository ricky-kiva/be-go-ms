package main

import "fmt"

func makeFunctionMethod() {
	// `make()` can initiate `channel`, `slice`, or `map`

	// making slice
	var cats = make([]string, 2)
	cats[0] = "Persian"
	cats[1] = "Anggora"

	var anotherCats = cats
	anotherCats[0] = "Bob"
	anotherCats[1] = "Mainecoon"

	fmt.Println("Cats:\t\t", cats)
	fmt.Println("Another cats:\t", anotherCats)
}
