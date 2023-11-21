package main

import "fmt"

func switchConditionalMethod() {
	var point = 9
	switch point {
	case 9:
		fmt.Println("Perfect. Wow, you're so talented in archery!")
	case 8:
		fmt.Println("Great. You're the top performers on the match!")
	default:
		fmt.Println("I see a potential in this sport within you.")
	}
}
