package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Name is the name object
type Name struct {
	fname string
	lname string
}

func main() {
	initFile()
	readFile()
}

func readFile() {
	fmt.Println("Hi Your read new read the file: data.txt")
	var sli = make([]Name, 0)
	file, err := os.Open("data.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := strings.Split(scanner.Text(), " ")
		a := Name{fname: name[0], lname: name[1]}
		sli = append(sli, a)
	}

	for _, value := range sli {
		fmt.Println(value.lname, value.fname)
	}
}

func initFile() {

	message := []byte("Oliver Jake \nJack Connor \nHarry Callum")
	ioutil.WriteFile("data.txt", message, 0644)

	// f, err := os.Create("data.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// l, err := f.WriteString("Oliver Jake \nJack Connor \nHarry Callum")
	// if err != nil {
	// 	fmt.Println(err)
	// 	f.Close()
	// 	return
	// }

	// fmt.Println(l, "bytes written successfully")
	// err = f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
