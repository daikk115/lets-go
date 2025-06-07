package main

import (
	"fmt"
	"reflect"
)

// array is value type, so when you pass it to a function, it copies the entire array
func testing(arr [2]int) {
	arr[0] = 42
	fmt.Println("Type", reflect.TypeOf(arr))
	fmt.Println(arr)
	fmt.Println("Len", len(arr))
	fmt.Println("Cap", cap(arr))
}

func main() {
	arr := [2]int{1, 2} // The size is part of the type, fixed size
	fmt.Println("Type", reflect.TypeOf(arr))
	fmt.Println("Len", len(arr))
	fmt.Println("Cap", cap(arr))
	fmt.Println(arr)
	testing(arr)
}
