package main

import (
	"fmt"
)

type I interface{}
type R struct{ i int }
type S struct{ s string }

func f(p I) {
	switch t := p.(type) {
	case S:
		fmt.Println("S", t.s)
	case R:
		fmt.Println("R", t.i)
	default:
	}
}

func main() {
	r := R{i: 1}
	s := S{s: "hello"}
	f(r)
	f(s)
}
