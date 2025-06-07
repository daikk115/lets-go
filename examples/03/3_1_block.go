package main

import "fmt"

func main() {
	fmt.Println("main: started")
	{
		x := 10
		fmt.Println("block: defer started")
		fmt.Println("x = ", x)
		fmt.Println("block: stopped")
	}
	fmt.Println("main: stopped")
	// fmt.Println("x = ", x) --> Compile error because out of block scope
}
