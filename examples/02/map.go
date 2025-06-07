package main

import (
	"fmt"
)

func testing(my_map map[string]int) {
	fmt.Println(my_map)
	value, key := my_map["MyOwn"]
	fmt.Println("Value", value, "Key", key)
	my_map["MyOwn"] = 100
	value = my_map["MyOwn"]
	fmt.Println("Value Only", value)
	fmt.Println(my_map)
}

func main() {
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, // A trailing comma is required
	}

	testing(monthdays)
}
