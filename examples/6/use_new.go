package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	x := new(int)    // *int
	y := new(Person) // *Person
	y.Name = "Dai Dang"
	y.Age = 29

	*x = 42
	fmt.Println(*x) // Output: 42
	fmt.Println(*y) // Output: {Dai Dang 29}
}
