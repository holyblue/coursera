package main

import "fmt"

func main() {
	var appleNum int

	fmt.Println("Number of apples?")

	num, err := fmt.Scan(&appleNum)

	fmt.Println(appleNum, num)

	if err != nil {
		println(err.Error)
	}
}