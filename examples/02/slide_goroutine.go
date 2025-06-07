package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{1, 2, 3}

	for i := 0; i < 3; i++ {
		go func() {
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println("Trong goroutine:", s)
		}()
	}

	time.Sleep(time.Duration(1) * time.Second)
	s = append(s, 4, 5, 6) // tạo array mới nếu vượt cap

	// Chờ các goroutine chạy xong (demo dùng sleep cho đơn giản)
	select {}
}
