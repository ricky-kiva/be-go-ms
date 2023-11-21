package main

import "fmt"

func forLoopMethod() {
	fmt.Print("Numbers to four: ")
	for i := 0; i < 5; i++ {
		if i == 4 {
			fmt.Printf("%d\n", i)
		} else {
			fmt.Printf("%d, ", i)
		}
	}

	fmt.Print("Numbers five to ten: ")
	var i = 5
	for i <= 10 {
		if i == 10 {
			fmt.Printf("%d\n", i)
		} else {
			fmt.Printf("%d, ", i)
		}
		i++
	}

	fmt.Print("Numbers six to nine: ")
	var j = 6
	for {
		if j == 9 {
			fmt.Printf("%d\n", j)
			break
		} else {
			fmt.Printf("%d, ", j)
		}
		j++
	}
}
