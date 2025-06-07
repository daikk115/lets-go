package main

import (
	"fmt"
	"strconv"
)

func convertString() {
	fmt.Println("\n------------------ convertString ------------------")
	s := "123"
	i, err := strconv.Atoi(s) // string to int
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(i)

	i2 := 456
	s2 := strconv.Itoa(i2) // int to string
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(s2)
}

func main() {
	convertString()
}
