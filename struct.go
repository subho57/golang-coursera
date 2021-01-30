package main

import "fmt"

// A P represents a struct containing a string and an corresponding integer, basically a key-value pair.
type P struct {
	x string
	y int
}

func maxStruct() {
	b := P{"x", -1}
	a := [...]P{P{"a", 10},
		P{"b", 2},
		P{"c", 3}}
	for _, z := range a {
		if z.y > b.y {
			b = z
		}
	}
	fmt.Println(b.x)
}
