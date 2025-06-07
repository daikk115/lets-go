/* Create a simple stack which can hold a fixed number of ints */

package main

import (
	"fmt"
)

func main() {
	s := my_struct{
		single_integer: 0,
		slice_integer:  []int{1, 2, 3, 4},
	}
	fmt.Println(s)
	s.plus_method_for_receiver(26)
	s.minus_method_for_receiver(9)
}

type my_struct struct {
	single_integer int
	slice_integer  []int
}

func (s *my_struct) plus_method_for_receiver(k int) {
	s.slice_integer[s.single_integer] += k
	s.single_integer++
	fmt.Println(s)
}

func (s *my_struct) minus_method_for_receiver(k int) {
	s.slice_integer[s.single_integer] = k
	s.single_integer--
	fmt.Print(s)
}
