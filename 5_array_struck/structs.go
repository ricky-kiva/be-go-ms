package main

import "fmt"

type character struct {
	name string
	hp   int
	mana int
}

func structMethod() {
	var ricky character
	ricky.name = "Rickyslash"
	ricky.hp = 85
	ricky.mana = 98

	fmt.Println("Name:\t", ricky.name)
	fmt.Println("HP:\t", ricky.hp)
	fmt.Println("Mana:\t", ricky.mana)
}
