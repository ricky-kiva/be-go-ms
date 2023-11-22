package main

import "fmt"

func slicesMethod() {
	var fruits = []string{"Banana", "Grape", "Apple"}

	var anotherFruits = fruits

	// slices has reference on the first variable. it will also modify the source variable
	anotherFruits[0] = "Orange"

	fmt.Printf("%s == %s\n", fruits, anotherFruits)
}
