package main

import "fmt"

func pairMe() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b, c uint
		fmt.Scan(&a, &b, &c)
		if a+b == c || b+c == a || a+c == b {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
