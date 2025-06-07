package main

import (
	"fmt"
	"os"
)

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	// if recover() is called out of defer function, it always nil.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f: ", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		f, err := os.Open("phuong-secrets.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			panic(fmt.Sprintf("Error opening file:", err))
		}
		defer f.Close()
		panic(fmt.Sprintf("Panic with value %v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
