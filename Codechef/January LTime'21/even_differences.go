package main

import "fmt"

func evenDifferences() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		odd := 0
		even := 0
		for j := 0; j < n; j++ {
			var temp int
			fmt.Scan(&temp)
			if temp%2 == 0 {
				even++
			} else {
				odd++
			}
		}
		if odd > even {
			fmt.Println(even)
		} else {
			fmt.Println(odd)
		}
	}
}
