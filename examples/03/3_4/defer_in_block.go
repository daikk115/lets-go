package main

import "fmt"

func main() {
	fmt.Println("main: started")
	{
		x := 10
		defer func() {
			fmt.Println("block: defer started, x =", x)
		}()
		defer fmt.Println("block: defer started, x =", x)
		x += 5
		fmt.Println("x = ", x)
		fmt.Println("block: stopped")
	}
	fmt.Println("main: stopped")
}
