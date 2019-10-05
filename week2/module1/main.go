package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numbers []int
	var input string
	index := 0

	fmt.Printf("input a number or press 'x' to finish.\n")

	for {
		promptInput(index, &input)

		if checkExit(input) {
			break
		}

		numbers = addNum(numbers, input)
		index++
	}

	fmt.Printf("before bubble sort the numbers are %v.\n", numbers)

	BubbleSort(numbers)

	fmt.Printf("after bubbles sort the number are %v:", numbers)
}

// BubbleSort sort slice number
func BubbleSort(numbers []int) {
	i := len(numbers)
	j := i
	rounds := 0
	var sortResults []bool

	for index := 0; index < i; index++ {
		if index == 0 || contains(sortResults, true) {
			sortResults = nil
			for index := 0; index < j-1; index++ {
				sli := numbers[index : index+2]
				sortResults = append(sortResults, Swap(sli))
			}
		} else {
			break
		}
		rounds++
	}

	fmt.Printf("total rounds is %d\n", rounds)
}

func contains(s []bool, e bool) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Swap sort slice number
func Swap(source []int) bool {
	isSort := false
	if (source)[0] > (source)[1] {
		temp := (source)[0]
		(source)[0] = (source)[1]
		(source)[1] = temp
		isSort = true
	}

	return isSort
}

func promptInput(index int, input *string) {
	fmt.Printf("This is the %d number: ", index+1)
	fmt.Scan(input)
}

func checkExit(input string) bool {
	quit := false
	if strings.ToUpper(input) == "X" {
		quit = true
	}

	return quit
}

func addNum(numbers []int, input string) []int {
	if v, e := strconv.Atoi(input); e == nil {
		numbers = append(numbers, v)
	} else {
		fmt.Println("Please insert the correct number.")
	}

	return numbers
}
