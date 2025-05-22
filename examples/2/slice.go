package main

import (
	"fmt"
	"reflect"
)

// array is value type, so when you pass it to a function, it copies the entire array
func testing(arr []int) {
	arr[0] = 42
	slice := append(arr, 33)
	fmt.Println("Type", reflect.TypeOf(slice))
	fmt.Println(slice)
	fmt.Println("Len", len(slice))
	fmt.Println("Cap", cap(slice))
}

func main() {
	arr := []int{1, 2}
	fmt.Println("Type", reflect.TypeOf(arr))
	fmt.Println(arr)
	fmt.Println("Len", len(arr))
	fmt.Println("Cap", cap(arr))
	testing(arr)
}
