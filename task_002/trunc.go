package main

import (
	"fmt"
)

func main() {
	var floatNum float64

	fmt.Println("Please enter a float number")

	num, err := fmt.Scan(&floatNum)

	if err == nil && num >= 1 {
		fmt.Printf("the original value: %f \n", floatNum)
		fmt.Printf("the trunc value: %d \n", int(floatNum))
	}
}
