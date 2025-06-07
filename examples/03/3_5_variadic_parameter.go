package main

import "fmt"

func func1(arg ...int) { // the variadic parameter is just a slice.
	for k, v := range arg {
		fmt.Println("Index", k, "Value", v)
	}
}

func func2(arg ...string) { // the variadic parameter is just a slice.
	for k, v := range arg {
		fmt.Println("Index", k, "Value", v)
	}
}

func main() {
	func1(1, 2)
	func1(3, 4, 5)
	func("a", "b")
	func2("c", "d", "e")
}
