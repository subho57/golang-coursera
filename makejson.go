package main

import (
	"encoding/json"
	"strings"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// initialising bufio
	reader := bufio.NewReader(os.Stdin)

	// reading the first param
	fmt.Print("Enter a Name: ")
	name, _ := reader.ReadString('\n')

	// reading the second param
	fmt.Print("Enter an Address: ")
	addr, _ := reader.ReadString('\n')

	// creating the map object
	mapObj := map[string]string {"name" : name, "address" : addr}

	// Marshalling the json from map
	jsonObj, _ := json.Marshal(mapObj)

	// printing the json by converting it into string, then removing \r and \n characters
	fmt.Println("Printing the json:\n", strings.Replace(string(jsonObj), "\\r\\n", "", -1))
}