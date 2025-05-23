package main

import "fmt"

func testing1() {
	var x interface{} = "hello"
	str := x.(string)
	fmt.Println(str)
	fmt.Print("Testing1 is OK\n----------------\n\n")
}

func testing2() {
	var x interface{} = "hello"
	str := x.(int)
	fmt.Println(str)
	fmt.Print("Testing2 is OK\n\n")
}

func main() {
	testing1()
	testing2()
}
