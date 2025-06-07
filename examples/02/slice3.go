package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2}
	t := &s
	fmt.Printf("Trước append: len=%d, cap=%d, addr=%p\n", len(s), cap(s), s)
	fmt.Println(t)
	s[0] = 42
	fmt.Printf("Sau gán: len=%d, cap=%d, addr=%p\n", len(s), cap(s), s)
	fmt.Println(t)
	s = append(s, 3, 4, 5)
	fmt.Printf("Sau append lan 1: len=%d, cap=%d, addr=%p\n", len(s), cap(s), s)
	s = append(s, 6)
	fmt.Printf("Sau append lan 2: len=%d, cap=%d, addr=%p\n", len(s), cap(s), s)
	fmt.Println(t)
}
