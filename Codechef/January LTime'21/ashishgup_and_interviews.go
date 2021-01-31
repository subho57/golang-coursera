package main

import (
	"fmt"
	"math"
)

func ashishgupAndInterviews() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, k, solved int
		fmt.Scan(&n, &k)
		tooSlow, noneGreaterthanOne := false, true
		for j := 0; j < n; j++ {
			var temp int
			fmt.Scan(&temp)
			if temp != -1 {
				solved++
			}
			if temp > k {
				tooSlow = true
			}
			if temp > 1 {
				noneGreaterthanOne = false
			}
		}
		if solved < int(math.Ceil(float64(n)/2.0)) {
			fmt.Println("Rejected")
		} else if tooSlow {
			fmt.Println("Too Slow")
		} else if noneGreaterthanOne && solved == n {
			fmt.Println("Bot")
		} else {
			fmt.Println("Accepted")
		}
	}
}
