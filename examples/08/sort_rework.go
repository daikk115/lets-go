package main

import "fmt"

// A Sorter should have methods Len(), Less(i, j int) bool, and Swap(i, j int)
type Sorter interface {
	Len() int           // len() as a method
	Less(i, j int) bool // p[i] > p[j] as a method
	Swap(i, j int)      // p[i], p[j] = p[j], p[i] as a method
}

type Xi []int
type Xs []string

// Len() for type Xi
func (p Xi) Len() int {
	return len(p)
}

// Less() for type Xi
func (p Xi) Less(i int, j int) bool {
	return p[j] < p[i]
}

// Swap() for type Xi
func (p Xi) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// And so on for type Xs
func (p Xs) Len() int {
	return len(p)
}

func (p Xs) Less(i int, j int) bool {
	return p[j] < p[i]
}

func (p Xs) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

func main() {
	ints := Xi{26, 9, 1994}
	strings := Xs{"Not", "My", "Name"}
	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)
}
