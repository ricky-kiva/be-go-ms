package main

import "fmt"

func main() {
	var name string = "Ricky"
	const like string = "fishes"
	var numOfFishes uint8 = 0
	var isLikeFish bool = true
	var fishSliced float32 = 0.5

	// No need to read this

	name = "Rickyslash"

	fmt.Print("Hello, my name is ", name, "!\n")
	fmt.Println("I really like", like)
	fmt.Printf("Currently, I have %d %s\n", numOfFishes, like)
	fmt.Printf("But does it mean I still like fishes? The answer is %t\n", isLikeFish)
	fmt.Printf("If a fish sliced, we will get two %.1f fish\n", fishSliced)

	/* Zero Values is the default value when no explicit value is provided.
	These are Zero Values in Go:

	string = ""
	bool = false
	non-decimal = 0
	decimal = 0.0
	composite (arrays) = nil
	*/
}
