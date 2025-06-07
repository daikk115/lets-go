package main

import "fmt"

func main() {
	var i interface{} = 42

	// Type assertion to int
	n := i.(int)
	fmt.Println("Asserted int:", n)

	// Safe assertion
	m, ok := i.(string)
	if ok {
		fmt.Println("Safe asserted int:", m)
	} else {
		fmt.Println("Assertion string failed")
	}
}
