package main

import (
	"context"
	"fmt"
)

func contextPassingMethod() {
	// context transfers values that extends from the parent context
	ctxParent := context.Background()
	ctx1 := context.WithValue(ctxParent, "key1", "King")
	ctx2 := context.WithValue(ctx1, "key2", "Queen")
	ctx3 := context.WithValue(ctx2, "key3", "Prince")
	ctx4 := context.WithValue(ctxParent, "key4", "Peoples")
	ctx5 := context.WithValue(ctx1, "key5", "Council")

	fmt.Println(ctx5.Value("key5"))
	fmt.Println(ctx5.Value("key4"))
	fmt.Println(ctx5.Value("key3"))
	fmt.Println(ctx5.Value("key2"))
	fmt.Println(ctx5.Value("key1"))

	fmt.Println(ctx3.Value("key2"))
	fmt.Println(ctx4.Value("key4"))
}
