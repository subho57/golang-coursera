package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// name struct has two fields, fname for the first name, and lname for the last name
type name struct {
	fname string
	lname string
}

func read() {
	/**
	 * Create a txt file with space separated first and last name only. Each line should contain one name only.
	 * The name of the file should not be more than 1 word, suffixed by .txt extension.
	 */
	var filename string
	fmt.Print("Enter the file name(e.g.: names.txt): ")
	fmt.Scan(&filename)

	// reading the file
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(errors.New("Something went wrong : FILE NOT FOUND"))
	} else if !strings.HasSuffix(filename, ".txt") {
		fmt.Println(errors.New("Something went wrong : ONLY .txt FILE SUPPORTED"))
	} else {
		// creating a slice of struct
		names := []name{}

		strdat := strings.Split(string(dat), "\n")
		for _, value := range strdat {
			var tempName name
			tempName.fname = strings.Split(value, " ")[0]
			tempName.lname = strings.Split(value, " ")[1]
			names = append(names, tempName)
		}
		fmt.Println("File reading complete...\nPrinting the names...")
		for i, v := range names {
			fmt.Printf("Name %d: %s %s\n", i+1, v.fname, v.lname)
		}
	}
}
