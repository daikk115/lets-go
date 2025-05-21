package main

import (
	"fmt"
	"time"
)

// goroutine là một lightweight thread được quản lý bởi Go runtime, dùng để thực thi đồng thời (concurrent) các hàm.
// main là một goroutine đặc biệt, và nó phải "giữ sống" chương trình đủ lâu để các goroutine khác chạy xong (thường dùng sync.WaitGroup, time.Sleep, v.v.).
// Nếu main() thoát, toàn bộ chương trình kết thúc, bao gồm mọi goroutine khác (nếu chúng chưa kết thúc).

func main() {
	go fmt.Println("New goroutine is running")
	fmt.Println("Main go routine")
	time.Sleep(2 * time.Second)
	fmt.Println("End after 2 seconds")
}
