package main

import "fmt"

func testing() {
	for {
		// This is an infinite loop
		fmt.Println("Hello World")
		break
	}

	for i := 0; i < 10; i++ {
		// This is normal for loop
		fmt.Println("For", i)
	}

	x := 0
	for x < 10 {
		// This is a while loop
		fmt.Println("While", x)
		x++
	}
}

func main() {
	testing()
}
