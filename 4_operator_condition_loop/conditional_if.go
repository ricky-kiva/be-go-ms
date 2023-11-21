package main

import "fmt"

func ifConditionalMethod() {
	var grade float32 = 3.6
	if grade == 10 {
		fmt.Println("Perfection, teachers must be very proud.")
	} else if grade > 5 {
		fmt.Println("Graduated. congrats!")
	} else if grade == 4 {
		fmt.Println("Almost bro")
	} else {
		fmt.Printf("Your grade is %.1f. It's okay, be great on the things you love\n", grade)
	}
}
