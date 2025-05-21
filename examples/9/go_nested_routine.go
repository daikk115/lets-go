package main

import (
	"fmt"
	"time"
)

// Chỉ cần main() chưa kết thúc, các nested routine vẫn sẽ tiếp tục

func main() {
	go func() {
		fmt.Println("Goroutine 1 started")

		// Goroutine bên trong Goroutine
		go func() {
			fmt.Println("Nested Goroutine started")
			time.Sleep(2 * time.Second)
			fmt.Println("Nested Goroutine stopped")
		}()

		fmt.Println("Goroutine 1 stopped")
	}()

	// Chờ tất cả goroutine hoàn thành
	time.Sleep(5 * time.Second)
	fmt.Println("Main Goroutine stopped")
}
