package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func print() {
	fmt.Println("\n------------------ Print ------------------")
	a := 3
	p := Person{
		Name: "John",
		Age:  30,
	}
	fmt.Printf("Value: %v, Go Syntax Value: %#v, Type: %T\n", a, a, a)
	fmt.Printf("Value: %v, Go Syntax Value: %#v, Type: %T\n", p, p, p)
}

func main() {
	print()
}
