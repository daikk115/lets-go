package main

import (
	"fmt"
	"os"
)

// Refer: https://victoriametrics.com/blog/defer-in-go/

func main() {
	// Defer just delay run by putting it in stack, not make new go routine
	defer fmt.Println("This line is defer for later execution")
	fmt.Println("Hello")
	doSomething()
	deferIsFILO()
}

func doSomething() error {
	f, err := os.Open("phuong-secrets.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer f.Close()
	fmt.Println("Ok opening file")
	return nil
}

func deferIsFILO() {
	defer fmt.Println("First In Last Out", 1)
	defer fmt.Println("First In Last Out", 2)
	defer fmt.Println("First In Last Out", 3)
}
