package main

import "fmt"

func evenSum() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		sum := 0
		for j := 0; j < n; j++ {
			var temp int
			fmt.Scan(&temp)
			sum += temp
		}
		if sum%2 == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
	}
}
