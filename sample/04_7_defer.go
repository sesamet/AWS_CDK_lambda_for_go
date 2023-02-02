package main

import "fmt"

func FILO(end int) {
	fmt.Println("counting")
	for i := 0; i < end; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	// Deferred function calls are pushed onto a stack.
	// When a function returns, its deferred calls are executed in last-in-first-out order.
	FILO(10)
}

// Defer
// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
