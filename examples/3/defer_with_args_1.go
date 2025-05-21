package main

import "fmt"

func pushAnalytic(a int) {
	fmt.Println("Address:", &a)
	fmt.Println("Value:", a)
}

func main() {
	a := 10
	defer pushAnalytic(a)
	a = 20
	b := 11
	pushAnalytic(a)
	pushAnalytic(b)
}

// Stack frame is temporary allocated in the stack while function running --> same Address here
