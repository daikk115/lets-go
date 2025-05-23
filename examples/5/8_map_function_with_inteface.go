package main

import "fmt"

// Define the empty interface as a type
type e interface{}

func checkType(input e) e {
	switch v := input.(type) {
	case int:
		return v * v
	case string:
		return v + v
	}
	return input
}

func mapF(n []e, f func(e) e) []e {
	m := make([]e, len(n))
	for k, v := range n {
		m[k] = f(v)
	}
	return m
}

func main() {
	m := []e{1, 2, 3, 4}
	s := []e{"a", "b", "c", "d"}
	fmt.Printf("Before\n")
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", s)
	mf := mapF(m, checkType)
	sf := mapF(s, checkType)
	fmt.Printf("\nAfter\n")
	fmt.Printf("%v\n", mf)
	fmt.Printf("%v\n", sf)
}
