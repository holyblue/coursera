package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var sli = make([]int, 3)
	var userInput string

	index := 0
	for {
		fmt.Printf("Please enter the %d number: ", index+1)
		fmt.Scan(&userInput)

		if strings.ToUpper(userInput) == "X" {
			break
		}

		if v, e := strconv.Atoi(userInput); e == nil {
			if index < 3 {
				sli[index] = v
			} else {
				sli = append(sli, v)
			}
			index++
		} else {
			fmt.Println("The input is not number!! Please enter the number again.")
		}
	}

	sort.Ints(sli)
	fmt.Println("the sorted value is bellow")
	fmt.Println(sli)

	main2()
}

func main2() {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2]
	z := x[1:5]
	fmt.Print(len(y), cap(y), len(z), cap(z))

	fmt.Println(z[3])
}
