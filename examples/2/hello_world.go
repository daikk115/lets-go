package main

import "fmt" // Implements formatted I/O

func print() {
	var name string
	var name2 *string
	name2 = &name
	name = "Dai"
	fmt.Printf("Hello %s\n", name)
	*name2 = "Dong"
	fmt.Printf("Hello %s\n", name)
}

/* Say Hello-World */
func main() {
	print()
}
