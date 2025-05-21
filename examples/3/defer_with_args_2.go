package main

import "fmt"

func pushAnalytic(a *int) {
	fmt.Println("Address:", &a)
	fmt.Println("Value: ", *a)
}

func main() {
	a := 10
	defer pushAnalytic(&a)
	a = 20
	pushAnalytic(&a)
}

// 2 pointers have different addresses but point to the same value
