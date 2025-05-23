package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(2)  // like append()
	l.PushFront(1) // like insert to index 0
	l.PushBack(3)

	fmt.Println("Head Value:", l.Front().Value)
	fmt.Println("Tail Value:", l.Back().Value)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}
