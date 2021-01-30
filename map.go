package main

import "fmt"

func mapExample1() {
	x := map[string]int{
		"ian": 1, "harris": 2}
	for i, j := range x {
		if i == "harris" {
			fmt.Print(i, j)
		}
	}
}
