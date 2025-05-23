package main

import (
	"fmt"
	"os"
)

// flag package look better
func main() {
	fmt.Println("All arguments:", os.Args)
	fmt.Println("Program name:", os.Args[0])

	if len(os.Args) > 1 {
		fmt.Println("First argument:", os.Args[1])
	} else {
		fmt.Println("No extra arguments provided.")
	}
}

// ubuntu@vm ~/G/l/e/10 (master)> go run 10_2_os_args.go input1 input2
// All arguments: [/tmp/go-build3907458438/b001/exe/10_2_os_args input1 input2]
// Program name: /tmp/go-build3907458438/b001/exe/10_2_os_args
// First argument: input1
