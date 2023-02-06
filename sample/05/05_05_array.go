package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	s := primes[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	nameA := names[0:2]
	nameB := names[1:3]
	fmt.Println(nameA, nameB)

	nameB[0] = "XXX"
	fmt.Println(nameA, nameB)
	fmt.Println(names)

}

// Arrays
// The type [n]T is an array of n values of type T.

// The expression

// var a [10]int
// declares a variable a as an array of ten integers.

// An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
