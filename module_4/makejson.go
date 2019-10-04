package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Person is the object of person
type Person struct {
	name    string
	address string
}

func main() {
	person1 := new(Person)
	person1.name = getInput("name")
	person1.address = getInput("address")
	fmt.Printf("Your name is '%s' and your address is '%s' \n", person1.name, person1.address)

	m := make(map[string]string)
	m["name"] = getInput("name")
	m["address"] = getInput("address")
	b, _ := json.Marshal(m)

	fmt.Printf("%s", b)
}

func getInput(fieldName string) string {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your %s: ", fieldName)

	if scanner.Scan() {
		input = scanner.Text()
	}

	return input
}
