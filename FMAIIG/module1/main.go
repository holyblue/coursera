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

	for {
		promptInput(index, &input)

		if checkExit(input) {
			break
		}

		addNum(&numbers, input)
		index++
	}

	fmt.Printf("the number is %v. \n", numbers)

	BubbleSort(numbers)

	fmt.Println(len(numbers))
	fmt.Printf("the number is %v:", numbers)
}

// BubbleSort sort slice number
func BubbleSort(numbers []int) {
	for index := 0; index < len(numbers) - 1; index++ {
		sli := numbers[index:index+2]

		fmt.Printf("slice before swap %v. \r\n", sli)
		Swap(&sli)
	}
}

// Swap sort slice number
func Swap(source *[]int) {
	if((*source)[0] > (*source)[1]) {
		temp := (*source)[0]
		(*source)[0] = (*source)[1]
		(*source)[1] = temp 
	}
}
 
func promptInput(index int, input *string) {
	fmt.Printf("請輸入第%d數字或按「x」結束輸入: ", index+1)
	fmt.Scan(input)
}

func checkExit(input string) bool {
	quit := false
	if strings.ToUpper(input) == "X" {
		quit = true
	}

	return quit
}

func addNum(numbers *[]int, input string) {
	if v, e := strconv.Atoi(input); e == nil {
		*numbers = append(*numbers, v)
	} else {
		fmt.Println("Please insert the correct number.")
	}
}
