package main

import "fmt"

type NameAge struct {
	name string // both non exported fiedls
	age  int
}

func main() {
	a := new(NameAge)
	a.name = "Peter"
	a.age = 42
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", *a)
}
