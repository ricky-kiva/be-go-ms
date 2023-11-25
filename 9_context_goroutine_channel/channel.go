package main

import (
	"fmt"
	"runtime"
)

func channelMethod() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string) // make channel

	var sayHello = func(name string) {
		var data = fmt.Sprintf("Hello, %s!", name)
		messages <- data
	}

	go sayHello("Rickyslash") // run `sayHello` closure on goroutine
	go sayHello("Geralt of Rivia")
	go sayHello("Peter Quill")

	var helloFirst = <-messages // get data from channel
	fmt.Println(helloFirst)
	var helloSecond = <-messages
	fmt.Println(helloSecond)
	var helloThird = <-messages
	fmt.Println(helloThird)
}
