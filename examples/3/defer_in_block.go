package main

import "fmt"

func main() {
	fmt.Println("main: started")
	{
		x := 10
		defer func() {
			fmt.Println("block: defer started")
		}()
		fmt.Println("x = ", x)
		fmt.Println("block: stopped")
	}
	fmt.Println("main: stopped")
}
