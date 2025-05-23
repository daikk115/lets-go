package main

import (
	"flag"
	"fmt"
)

func main() {
	// Khai báo flag
	name := flag.String("name", "Go Developer", "Your name")
	age := flag.Int("age", 25, "Your age")
	debug := flag.Bool("debug", false, "Enable debug mode")

	// Parse các flag từ command line
	flag.Parse()

	// Sử dụng
	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
	fmt.Println("Debug mode:", *debug)
}

// Example command line:
// go run 4_8_useful_package/8_flag_args.go --name Daidv --age 30
