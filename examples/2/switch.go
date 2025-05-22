package main

import "fmt"

func testing() {
	switch {
	case 1 == 2:
		fmt.Println("True")
	case 2 == 3:
		fmt.Println("False")
	default:
		fmt.Println("Default")
	}
}

func main() {
	testing()
}
