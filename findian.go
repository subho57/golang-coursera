package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a string: ")
	str, _ := reader.ReadString('\n')
	str = strings.Replace(str, "\n", "", -1)

	i := str[0] == 'i' || str[0] == 'I'
	a := false
	n := str[len(str) - 1] == 'n' || str[len(str) - 1] == 'N'

	for j := 1; j < len(str) - 1; j++ {
		if str[j] == 'a' || str[j] == 'A' {
			a = true
			break
		}
	}
	
	if i && a && n {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}