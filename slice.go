package main

import (
	"fmt"
	"sort"
	"strconv"
)

func sliceExample1() {
	x := [...]int{4, 8, 5}
	y := x[0:2]
	z := x[1:3]
	y[0] = 1
	z[1] = 3
	fmt.Print(x)
}

func sliceExample2() {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2]
	z := x[1:4]
	fmt.Print(len(y), cap(y), len(z), cap(z))
}

func sliceExample3() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
  }
  

func slice() {
	sli := make([]int, 0, 3)
	for {
		var ch string
		fmt.Print("Enter an integer: ")
		fmt.Scan(&ch)
		if ch[0] == 'X' {
			break
		}
		num, err := strconv.Atoi(ch)
		if err == nil {
			sli = append(sli, num)
			sort.Ints(sli)
			fmt.Println(sli)
		}

	}
}
