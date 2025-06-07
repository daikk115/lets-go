package main

import (
	"fmt"
	"sort"
)

func sortSlice() {
	fmt.Println("\n------------------ sortArray ------------------")
	arr := []int{5, 2, 8, 1, 4}
	fmt.Println("Before sorting:", arr, &arr[0])
	sort.Ints(arr)
	fmt.Println("After sorting:", arr, &arr[0])
}

func main() {
	sortSlice()
}
