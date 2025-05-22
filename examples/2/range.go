package main

import "fmt"

func testing() {
	// range can loop over slices, arrays, strings, maps & channels.
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		fmt.Println("Key:", k, "Value:", v)
	}
}

func main() {
	testing()
}
