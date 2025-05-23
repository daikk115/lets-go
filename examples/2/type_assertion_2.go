package main

import "fmt"

type e interface{}

func checkType(input e) {
	switch input.(type) {
	case int:
		fmt.Println("Integer")
		return
	case string:
		fmt.Println("String")
		return
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	a := 1
	b := "hello"
	c := 3.14
	checkType(a)
	checkType(b)
	checkType(c)
}
