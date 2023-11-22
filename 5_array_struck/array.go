package main

import "fmt"

func arrayMethod() {
	var sharks = [4]string{"Blue", "Great White", "Whale", "Thresher"}

	fmt.Println("Listed sharks:\t\t", sharks)
	fmt.Println("Listed sharks total:\t", len(sharks))
	fmt.Println("First indexed shark:\t", sharks[0])

	var anotherSharks = sharks
	anotherSharks = [4]string{"Nurse", "Port Jackson", "Tiger", "Goblin"}

	fmt.Println("Another Listed sharks:\t", anotherSharks)
	fmt.Println("Compare the sharks:\t", sharks) // array won't change its value if assigned to another variable

	var unlimitedSharks = [...]string{"Blue", "Great White", "Whale", "Thresher", "Nurse", "Port Jackson", "Tiger", "Goblin"}
	fmt.Println("Joined sharks:\t\t", unlimitedSharks)
	fmt.Println("Total sharks:\t\t", len(unlimitedSharks))
}
