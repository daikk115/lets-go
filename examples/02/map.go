package main

import (
	"fmt"
)

func testing(my_map map[string]int) {
	fmt.Println(my_map)
	value, existing := my_map["MyOwn"]
	fmt.Println("Value", value, "existing", existing)
	my_map["MyOwn"] = 100
	fmt.Println(my_map)
	value, existing = my_map["MyOwn"]
	value = my_map["MyOwn"]
	fmt.Println("Value", value, "existing", existing)
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
