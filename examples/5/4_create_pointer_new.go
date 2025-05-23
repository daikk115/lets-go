package main

import "fmt"

func test(a *int) {
	*a = 2
}

func main() {
	a := 1
	test(&a)

	fmt.Printf("a = %d\n", a)
	ptr := new(int)
	*ptr = 100
	fmt.Printf("Ptr = %#x, Ptr value = %d\n", ptr, *ptr)
}
