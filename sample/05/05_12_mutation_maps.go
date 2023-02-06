package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("1:", m)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	fmt.Println("2:", m)

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	fmt.Println("3:", m)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	fmt.Println("4:", m)

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	fmt.Println("5:", m)
}

// Mutating Maps
// Insert or update an element in map m:

// m[key] = elem
// Retrieve an element:

// elem = m[key]
// Delete an element:

// delete(m, key)
// Test that a key is present with a two-value assignment:

// elem, ok = m[key]
// If key is in m, ok is true. If not, ok is false.

// If key is not in the map, then elem is the zero value for the map's element type.

// Note: If elem or ok have not yet been declared you could use a short declaration form:

// elem, ok := m[key]
