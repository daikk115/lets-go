package main

import "fmt"

func main() {
	n := []int{5, -1, 0, 12, 3, 5}
	// n := []int{1, 2, 5, 4, 8}
	n2 := []int{5, -1, 0, 12, 3, 5}
	// n2 := []int{1, 2, 5, 4, 8}
	bubblesort(n)
	bubblesort_optimize(n2)
}

func bubblesort(n []int) {
	fmt.Print("--------bubblesort---------\n")
	fmt.Print(n, "\n")
	number_if_swaps := 0
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			fmt.Print("Compare index: ", j, " vs ", i, "\n")
			number_if_swaps += 1
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
				fmt.Print(n, "\n")
			}
		}
	}
	fmt.Print("number_if_swaps=", number_if_swaps, "\n")
}

func bubblesort_optimize(unsorted_list []int) {
	fmt.Print("--------bubblesort_optimize---------\n")
	fmt.Print(unsorted_list, "\n")
	n := len(unsorted_list)
	number_if_swaps := 0
	swapped := true
	for i := n - 1; i >= 1; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			fmt.Print("Compare index: ", j, " vs ", j+1, "\n")
			number_if_swaps += 1
			if unsorted_list[j] > unsorted_list[j+1] {
				swapped = true
				unsorted_list[j], unsorted_list[j+1] = unsorted_list[j+1], unsorted_list[j]
				fmt.Print(unsorted_list, "\n")
			}
		}
		if swapped == false {
			break
		}
	}
	fmt.Print("number_if_swaps=", number_if_swaps, "\n")
}
