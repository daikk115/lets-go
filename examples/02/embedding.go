package main

import "fmt"

// Embedded struct
type Person struct {
	Name string
	Age  int
}

type Monkey struct {
	Name   string
	Banana int
}

// Struct that embeds Person
type Employee struct {
	Person // embedded
	Monkey
	// Name       string
	EmployeeID string
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "Alice",
			Age:  30,
		},
		Monkey: Monkey{
			Name:   "Monkey1",
			Banana: 30,
		},
		// Name:       "Alice",
		EmployeeID: "E12345",
	}

	// Access embedded fields directly
	// emp.Name = "Bob"
	fmt.Println(emp.Monkey.Name) // Alice
	fmt.Println(emp.Age)         // 30
	fmt.Println(emp.EmployeeID)  // E12345
}
