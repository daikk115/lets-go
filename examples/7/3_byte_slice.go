package main

import "fmt"

func main() {
	str := "hello"
	bytes := []byte(str)    // string to []byte
	newStr := string(bytes) // []byte to string

	fmt.Println("Original:", str)
	fmt.Println("Bytes:", bytes)
	fmt.Println("Converted back:", newStr)
}
