package main

import "fmt"

// Embedded struct
type Person struct {
	Name string
	Age  int
}

// Struct that embeds Person
type Employee struct {
	Person     // embedded
	EmployeeID string
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "Alice",
			Age:  30,
		},
		EmployeeID: "E12345",
	}

	// Access embedded fields directly
	emp.Name = "Bob"
	fmt.Println(emp.Name)       // Alice
	fmt.Println(emp.Age)        // 30
	fmt.Println(emp.EmployeeID) // E12345
}
