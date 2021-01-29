package main

import (
	"fmt"
)

func main() {
	var input float64
	prompt := "Enter a floating point number: "
	fmt.Printf(prompt)
	fmt.Scan(&input)
	fmt.Println("Truncated Value =", int64(input))
}