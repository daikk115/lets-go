package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

// Embedding all read and write methods
type ReadWriter interface {
	Reader
	Writer
}

type TestStruct struct {
}

func (t TestStruct) Read() {
	fmt.Println("Read something!")
}

func (t TestStruct) Write() {
	fmt.Println("Write something!")
}

func main() {
	t := TestStruct{}
	t.Read()
	t.Write()
}
