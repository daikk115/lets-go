package main

import (
	"fmt"
)

func printit(x int) {
	fmt.Println("Print", x)
}

func callback(y int, f func(int)) {
	fmt.Println("Callback", y, f)
	f(y)
}

func main() {
	fmt.Println("Starting main")
	callback(10, printit)
}
