package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := func() { // a is defined as an anonymous (nameless) function,
		fmt.Println("Hello")
	}
	fmt.Println(reflect.TypeOf(a))
	a()
}
