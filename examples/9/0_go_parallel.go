package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOMAXPROCS is set to:", runtime.GOMAXPROCS(0)) // Get current value
	runtime.GOMAXPROCS(2)                                       // Use 2 CPU cores
	fmt.Println("GOMAXPROCS is set to:", runtime.GOMAXPROCS(0)) // Get current value
}

// ubuntu@vm ~/G/l/e/9 (master)> GOMAXPROCS=4 go run 0_go_parallel.go
// GOMAXPROCS is set to: 4
// GOMAXPROCS is set to: 2
