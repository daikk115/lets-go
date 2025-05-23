/*
The input function is applied to each member in the list and a new list
containing these calculated values is returned.
*/
package main

import "fmt"

func applyFunctionToItems(inputFunction func(int) int, oldList []int) []int {
	newList := make([]int, len(oldList))
	for k, v := range oldList {
		newList[k] = inputFunction(v)
	}

	return newList
}

func main() {
	m := []int{2, 3, 4}
	f := func(i int) int {
		return i * i
	}

	fmt.Println("List:", (m))
	fmt.Println("Squared each items:", applyFunctionToItems(f, m))
}
