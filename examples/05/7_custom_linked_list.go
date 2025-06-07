package main

import (
	"fmt"
)

type Value int

type zNode struct {
	Value
	prev, next *zNode
}

type zList struct {
	head, tail *zNode
}

func (currentList *zList) InsertHead(value Value) {
	newNode := &zNode{Value: value}

	if currentList.head == nil {
		// If the list is empty, set both head and tail to the new node
		currentList.head = newNode
		currentList.tail = newNode
	} else {
		newNode.next = currentList.head
		currentList.head.prev = newNode
	}
	currentList.head = newNode
}

func (currentList *zList) InsertTail(value Value) {
	newNode := &zNode{Value: value}

	if currentList.tail == nil {
		// If the list is empty, set both head and tail to the new node
		currentList.head = newNode
		currentList.tail = newNode
	} else {
		newNode.prev = currentList.tail
		currentList.tail.next = newNode
	}
	currentList.tail = newNode
}

func main() {
	l := new(zList)
	l.InsertHead(2)
	l.InsertTail(3)
	l.InsertHead(1)

	fmt.Println("Head Value:", l.head.Value)
	fmt.Println("Tail Value:", l.tail.Value)

	if l.head != nil {
		for n := l.head; n != nil; n = n.next {
			fmt.Printf("%v\n", n.Value)
		}
	} else {
		fmt.Println("List is empty")
	}
}
