package main

import "fmt"

func mapsMethod() {
	var egg = map[string]int{
		"january":  50,
		"february": 40,
		"march":    60,
	}

	egg["april"] = 65
	fmt.Println("Egg production to April:", egg)
}
