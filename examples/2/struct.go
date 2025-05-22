package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// struct is value types
func testing(persion Person) {
	persion.Name = "DongHM"
	persion.Age = 29
	fmt.Println(persion.Name, persion.Age)
}

func main() {
	daidv := Person{
		Name: "Daidv",
		Age:  30,
	}
	fmt.Println(daidv.Name, daidv.Age)
	testing(daidv)
}
