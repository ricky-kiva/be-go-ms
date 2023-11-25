package main

/* In the context of concurrency, your computer is like a kitchen with several workers (cores)
who each can handle different to-do lists (threads). Each project (process) might have one
or more to-do lists, and these lists include specific jobs (tasks) that need to be done.
The computer, like a smart kitchen manager, juggles all these projects and tasks efficiently
to get everything done. */

// run by `go run .`
func main() {
	goroutineMethod()
	channelMethod()
}
